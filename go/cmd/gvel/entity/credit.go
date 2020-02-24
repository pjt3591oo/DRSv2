package entity

import "github.com/velo-protocol/DRSv2/go/libs/vclient"

// SetupCredit
type SetupCreditInput struct {
	Passphrase          string
	AssetCode           string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}

type SetupCreditOutput struct {
	TxHash              string
	CreditOwnerAddress  string
	AssetCode           string
	AssetAddress        string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}

// MintCreditByCollateral
type MintCreditByCollateralInput struct {
	AssetCode        string
	CollateralAmount string
	Passphrase       string
}
type MintCreditByCollateralOutput struct {
	AssetCode          string
	StableCreditAmount string
	TxHash             string
}

// MintCreditByCredit
type MintCreditByCreditInput struct {
	Passphrase   string
	AssetCode    string
	CreditAmount string
}
type MintCreditByCreditOutput struct {
	TxHash              string
	AssetCode           string
	StableCreditAmount  string
	AssetAddress        string
	CollateralAssetCode string
	CollateralAmount    string
}

// GetCreditExchange
type GetCreditExchangeInput struct {
	AssetCode  string
	Passphrase string
}
type GetCreditExchangeOutput struct {
	AssetCode                     string
	CollateralAssetCode           string
	PriceInCollateralPerAssetUnit string
}

type RedeemStableCreditInput struct {
	RedeemAmount     string
	AssetCode        string
	Passphrase       string
}
type RedeemStableCreditOutput struct {
	CollateralAmount     string
	AssetCode            string
	TxHash               string
}

// RebalanceCredit
type RebalanceCollateralInput struct {
	Passphrase string
}
type RebalanceCollateralOutput struct {
	TxHash string
	vclient.RebalanceEvent
}
