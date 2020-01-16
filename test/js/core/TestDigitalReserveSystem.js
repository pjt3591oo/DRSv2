const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const PriceFeeders = artifacts.require('PriceFeeders');
const StableCredit = artifacts.require("StableCredit");
const Token = artifacts.require('Token');
const ReserveManager = artifacts.require('ReserveManager');
const MockContract = artifacts.require("MockContract");

const Web3 = require('web3');
const truffleAssert = require('truffle-assertions');
const helper = require('../testhelper');

let drs, heart, priceFeeder, reserveManager, veloCollateralAsset, otherCollateralAsset, stableCreditVUSD, mocks;

contract("DigitalReserveSystem test", async accounts => {

  before(async () => {
    mocks = {
      heart: await MockContract.new(),
      priceFeeder: await MockContract.new(),
      reserveManager: await MockContract.new(),
      veloCollateralAsset: await MockContract.new(),
      otherCollateralAsset: await MockContract.new(),
      stableCreditvUSD: await MockContract.new(),
    };

    heart = await Heart.at(mocks.heart.address);
    priceFeeder = await PriceFeeders.at(mocks.priceFeeder.address);
    reserveManager = await ReserveManager.at(mocks.reserveManager.address);
    veloCollateralAsset = await Token.at(mocks.veloCollateralAsset.address);
    otherCollateralAsset = await Token.at(mocks.otherCollateralAsset.address);
    stableCreditVUSD = await StableCredit.at(mocks.stableCreditvUSD.address);
  });

  beforeEach(async () => {
    drs = await DigitalReserveSystem.new(heart.address);
  });

  afterEach(async () => {
    const promises = [];
    for (const key in mocks) {
      promises.push(mocks[key].reset());
    }
    await Promise.all(promises);
  });

  describe("SetupCredit", async () => {
    it("should setup credit successfully", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        helper.address(0)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        helper.address(1)
      );
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getPriceFeeders"),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        helper.methodABI(priceFeeder, "getMedianPrice", [helper.address(0)]),
        1
      );

      const result = await drs.setup(
        Web3.utils.fromAscii("VELO"),
        Web3.utils.fromAscii("USD"),
        "vUSD",
        1
      );

      truffleAssert.eventEmitted(result, 'Setup', async event => {
        assert.equal(event.assetCode, "vUSD");
        helper.assertEqualByteString(event.peggedCurrency, "USD");
        assert.equal(event.peggedValue, 1);
        helper.assertEqualByteString(event.collateralAssetCode, "VELO");
        assert.ok(web3.utils.isAddress(event.assetAddress));

        const stableCredit = await StableCredit.at(event.assetAddress);
        const name = await stableCredit.name();
        assert.equal(name, "vUSD");

        return true;
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        false
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner")
      }
    });

    it("should fail, invalid assetCode format", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSDvUSDvUSDvUSDvUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: invalid assetCode format")
      }
    });

    it("should fail, assetCode has already been used", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: assetCode has already been used")
      }
    });

    it("should fail, collateralAssetCode does not exist", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: collateralAssetCode does not exist")
      }
    });

    it("should fail, collateralAssetCode - peggedCurrency pair does not exist", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        false
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: collateralAssetCode - peggedCurrency pair does not exist")
      }
    });

    it("should fail, price of link must have value more than 0", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getPriceFeeders"),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        helper.methodABI(priceFeeder, "getMedianPrice", [helper.address(0)]),
        0
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: price of link must have value more than 0")
      }
    });
  });

  describe("MintFromCollateralAmount", async () => {
    it("should mint from collateral correctly", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("THB"),
        accounts[1],
        Web3.utils.fromAscii('VELO'),
        veloCollateralAsset.address,
        'vTHB',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("stableCreditId")).encodeABI(),
        stableCredit.address
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );

      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        10000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCreditIssuanceFee().encodeABI(),
        100000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
        10000000
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("VELO")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
        true
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], veloCollateralAsset.address, 1).encodeABI(),
        true
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], drs.address, 8).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getDrsAddress().encodeABI(),
        drs.address
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.approve(accounts[0], 1).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      await mocks.reserveManager.givenMethodReturnBool(
        reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
        true
      );

      const result = await drs.mintFromCollateralAmount(100000000, 'vTHB');

      truffleAssert.eventEmitted(result, 'Mint', event => {
        return event.assetCode === "vTHB"
          && new web3.utils.BN(event.mintAmount).toNumber() === 990000000
          && new web3.utils.BN(event.collateralAmount).toNumber() === 100000000
          && event.assetAddress === stableCredit.address
          && event.collateralAssetCode === web3.utils.padRight(web3.utils.utf8ToHex("VELO"), 64);
      }, 'contract should emit the event correctly');

    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        false
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: stableCredit not exist", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        0
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: price of link must have value more than 0", err.reason)
      }
    });
  });

  describe("MintFromStableCreditAmount", async () => {
    it("should mint from stable credit correctly", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii('USD'),
        accounts[1],
        Web3.utils.fromAscii('VELO'),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii('stableCreditId')).encodeABI(),
        stableCredit.address
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );

      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii('')).encodeABI(),
        10000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCreditIssuanceFee().encodeABI(),
        100000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
        10000000
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii('VELO')).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getDrsAddress().encodeABI(),
        drs.address
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.approve(accounts[0], 10).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      await mocks.reserveManager.givenMethodReturnBool(
        reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
        true
      );

      const result = await drs.mintFromStableCreditAmount(100000000, Web3.utils.fromAscii("vUSD"));
      truffleAssert.eventEmitted(result, 'Mint', event => {
        const BN = web3.utils.BN;
        const eventMintAmount = new BN(event.mintAmount).toNumber();
        const eventCollateralAmount = new BN(event.collateralAmount).toNumber();
        return eventMintAmount === 100000000
          && web3.utils.isAddress(event.assetAddress) === true
          && web3.utils.hexToUtf8(event.assetCode) === 'vUSD'
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && eventCollateralAmount === 10101010
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        false
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: stableCredit not exist", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {

        assert.equal("DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        0
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: price of link must have value more than 0", err.reason)
      }
    });
  });

  describe("GetExchangeRate", async () => {
    it("should get exchange rate correctly", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditvUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        10000000
      );

      const result = await drs.getExchange("vUSD");
      const BN = web3.utils.BN;
      const pricePerAssetUnit = new BN(result[2]).toNumber();

      assert.equal("vUSD", result[0]);
      assert.equal(web3.utils.padRight(0x0, 64), result[1]);
      assert.equal(0, pricePerAssetUnit);
    });

    it("should fail, invalid assetCode format", async () => {
      const expectedErr = "Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem.getExchange: invalid assetCode format";
      try {
        await drs.getExchange("");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.getExchange("vUSDTHBKHRSGDJPYEURBDHUSSkbph");
      } catch (err) {
        assert.equal(expectedErr, err)
      }
    });

    it("should fail, stableCredit not exist", async () => {
      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: stableCredit not exist", err)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err)
      }
    });

    it("should fail, price of link must have value more than 0", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditvUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        0
      );

      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: price of link must have value more than 0", err)
      }
    });
  });

  describe("Get Collateral Health Check", async () => {
    it("should collateral health check correctly", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditvUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        10000000
      );

      const result = await drs.collateralHealthCheck("vUSD", "VELO");
      const BN = web3.utils.BN;
      const requiredAmount = new BN(result[1]).toNumber();
      const presentAmount = new BN(result[2]).toNumber();
      assert.equal(web3.utils.padRight(0x0, 64), result[0]);
      assert.equal(0, requiredAmount);
      assert.equal(0, presentAmount);
    });

    it("should fail, invalid assetCode format", async () => {
      const expectedErr = "Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem.collateralHealthCheck: invalid assetCode format";
      try {
        await drs.collateralHealthCheck("", "");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("VELO", "");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("", "VELO");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("", "VELOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("VELOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO", "");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("VELOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO", "VELO");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.collateralHealthCheck("VELO", "VELOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO");
      } catch (err) {
        assert.equal(expectedErr, err)
      }
    });

    it("should fail, stableCredit not exist", async () => {
      try {
        await drs.collateralHealthCheck("vUSD", "VELO");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: stableCredit not exist", err)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.collateralHealthCheck("vTHB", "VELO");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err)
      }
    });

    it("should fail, collateralAssetCode does not exist", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditvUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mocks.priceFeeder.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        10000000
      );

      try {
        await drs.collateralHealthCheck("vUSD", "VELO");
      } catch (err) {
        assert.equal("DigitalReserveSystem.collateralHealthCheck: collateralAssetCode does not exist1", err.reason)
      }
    });
  });
});

