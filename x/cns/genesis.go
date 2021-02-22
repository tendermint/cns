package cns

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cns/x/cns/keeper"
	"github.com/tendermint/cns/x/cns/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	//TODO: Use params to set fee and accepted denoms
	infos := genState.Info

	for _, info := range infos {
		k.SetChainInfo(ctx, info)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var infos []types.ChainInfo
	k.IterateAllInfos(ctx, func(info types.ChainInfo) bool {
		infos = append(infos, info)
		return false
	})

	return &types.GenesisState{
		Fee:  sdk.NewCoins(types.DefaultFee(types.DefaultDenom)),
		Info: infos,
	}
}
