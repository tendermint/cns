package keeper

import (
	"github.com/tendermint/cns/x/cns/types"
)

var _ types.QueryServer = Keeper{}
