package keeper

import (
	"context"
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

func (k msgServer) RegisterChainName(ctx context.Context, request *types.MsgRegisterChainNameRequest) (*types.MsgRegisterChainNameResponse, error) {

	panic("implement me")
}

func (k msgServer) UpdateChainInfo(ctx context.Context, request *types.MsgUpdateChainInfoRequest) (*types.MsgUpdateChainInfoResponse, error) {
	panic("implement me")
}
