package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cns/x/cns/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) RegisterChainName(goCtx context.Context, msg *types.MsgRegisterChainNameRequest) (*types.MsgRegisterChainNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	cInfo := types.ChainInfo{
		ChainName:          msg.ChainName,
		Expiration:         0,
		Owner:              msg.Owner,
		CanonicalIbcClient: msg.CanonicalIbcClient,
		Seed:               msg.Seed,
		SourceCodeUrl:      msg.SourceCodeUrl,
		Version:            msg.Version,
		Fee:                sdk.NewCoins(types.DefaultFee(types.DefaultDenom)),
	}

	err = k.Register(ctx, cInfo)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterChainNameResponse{}, err
}

func (k msgServer) UpdateChainInfo(goCtx context.Context, msg *types.MsgUpdateChainInfoRequest) (*types.MsgUpdateChainInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	cInfo := types.ChainInfo{
		ChainName:          msg.ChainName,
		Expiration:         0,
		Owner:              msg.Owner,
		CanonicalIbcClient: msg.CanonicalIbcClient,
		Seed:               msg.Seed,
		SourceCodeUrl:      msg.SourceCodeUrl,
		Version:            msg.Version,
	}

	err = k.Update(ctx, cInfo)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateChainInfoResponse{}, nil

}
