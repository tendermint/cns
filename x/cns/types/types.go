package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var DefaultFee = func(denom string) sdk.Coin { return sdk.NewCoin(denom, sdk.NewInt(100000)) }

const DefaultDenom = sdk.DefaultBondDenom
