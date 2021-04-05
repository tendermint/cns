package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/23-commitment/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

func (k Keeper) ProveClientState(ctx sdk.Context, clientId string, proofHeight exported.Height, clientState exported.ClientState, proof []byte, header exported.Header) error {
	prefix := k.connKeeper.GetCommitmentPrefix()
	prefixBz := commitmenttypes.NewMerklePrefix(prefix.Bytes())

	return clientState.VerifyClientState(k.clientKeeper.ClientStore(ctx, clientId), k.cdc, proofHeight,
		&prefixBz, clientId, proof, clientState)

}
