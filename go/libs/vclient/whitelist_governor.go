package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/constants"
	"strings"
)

type WhitelistGovernorInput struct {
	Address string
}

type WhitelistGovernorOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
}

type WhitelistGovernorAbiInput struct {
	Address common.Address
}

func (i *WhitelistGovernorInput) Validate() error {
	if len(i.Address) == 0 {
		return errors.New("address must not be blank")
	}

	if !common.IsHexAddress(i.Address) {
		return errors.New("invalid address format")
	}

	return nil
}

func (i *WhitelistGovernorInput) ToAbiInput() WhitelistGovernorAbiInput {
	return WhitelistGovernorAbiInput{
		Address: common.HexToAddress(i.Address),
	}
}

func WhitelistGovernorReplaceError(prefix string, err error) error {
	if strings.Contains(err.Error(), "the message sender is not found or does not have sufficient permission") {
		return errors.Wrap(errors.New("the message sender is not found or does not have sufficient permission to perform whitelist user"), prefix)
	}
	return errors.Wrap(err, prefix)
}

func (c *Client) WhitelistGovernor(ctx context.Context, input *WhitelistGovernorInput) (*WhitelistGovernorOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	senderIsGovernor, err := c.contract.heart.IsGovernor(nil, c.publicKey)
	if err != nil {
		return nil, err
	}
	if !senderIsGovernor {
		return nil, errors.New("the message sender is not found or does not have sufficient permission to perform whitelist user")
	}

	isGovernor, err := c.contract.heart.IsGovernor(nil, input.ToAbiInput().Address)
	if err != nil {
		return nil, err
	}
	if isGovernor {
		return nil, errors.Errorf("the address %s has already been whitelisted as governor", input.Address)
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.heart.SetGovernor(opt, input.ToAbiInput().Address)
	if err != nil {
		return nil, WhitelistGovernorReplaceError("smart contract call error", err)
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
	if err != nil {
		return nil, WhitelistGovernorReplaceError("confirm transaction error", err)
	}

	return &WhitelistGovernorOutput{
		Tx:      tx,
		Receipt: receipt,
	}, nil
}
