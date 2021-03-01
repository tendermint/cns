package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cns module sentinel errors
var (
	ErrChainAlreadyExists = sdkerrors.Register(ModuleName, 2, "chain name is already registered")
)
