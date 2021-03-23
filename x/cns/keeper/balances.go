package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	types4 "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	types2 "github.com/cosmos/cosmos-sdk/x/ibc/core/03-connection/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	types3 "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	"github.com/tendermint/cns/x/cns/types"
	"reflect"
	"strings"
)

func (k Keeper) FetchUpdatedBalances(ctx sdk.Context, addr sdk.AccAddress) (sdk.Coins, error) {
	balances := k.bankKeeper.GetAllBalances(ctx, addr)
	if balances.Empty() {
		return nil, fmt.Errorf("address not found")
	}
	var newBalances []sdk.Coin
	for _, bal := range balances {
		if strings.HasPrefix(bal.Denom, "ibc/") {
			denomPath, err := k.ibctransfer.DenomPathFromHash(ctx, bal.Denom)
			if err != nil {
				return nil, err
			}

			denomTrace := transfertypes.ParseDenomTrace(denomPath)
			identifiers := strings.Split(denomTrace.Path, "/")

			conn, err := k.GetConnFromIdentifiers(ctx, identifiers)
			cs1, ok := k.clientKeeper.GetClientState(ctx, conn.ClientId)
			if !ok {
				return nil, fmt.Errorf("unable to fetch client state")
			}
			var cs2 exported.ClientState
			var matchedInfos []types.ChainInfo
			k.IterateAllInfos(ctx, func(info types.ChainInfo) bool {
				if info.CanonicalIbcClient != "" {
					cs2, ok = k.clientKeeper.GetClientState(ctx, info.CanonicalIbcClient)
					if !ok {
						return false
					}

					equal := k.CompareClientStates(ctx, cs1, cs2)
					if equal {
						matchedInfos = append(matchedInfos, info)
						return false
					}
				}
				return false
			})

			if len(matchedInfos) != 1 {
				return nil, fmt.Errorf("invalid client used for transfer")
			}

			connCS, ok := k.clientKeeper.GetLatestClientConsensusState(ctx, conn.ClientId)
			if !ok {
				return nil, fmt.Errorf("unable to fetch client consensus state")
			}

			storeCS, ok := k.clientKeeper.GetLatestClientConsensusState(ctx, matchedInfos[0].CanonicalIbcClient)
			if !ok {
				return nil, fmt.Errorf("unable to fetch client consensus state")
			}

			storeTmCs := storeCS.(*types3.ConsensusState)
			connTmCs := connCS.(*types3.ConsensusState)
			if storeTmCs.String() == connTmCs.String() {
				return nil, fmt.Errorf("consensus state mismatch")
			}

			ibcDenom := "ibc/" + denomTrace.BaseDenom + "/" + matchedInfos[0].ChainName
			bal = sdk.NewCoin(ibcDenom, bal.Amount)
		}
		newBalances = append(newBalances, bal)
	}
	return newBalances, nil
}

func (k Keeper) GetConnFromIdentifiers(ctx sdk.Context, identifiers []string) (types2.ConnectionEnd, error) {
	channel, ok := k.channelKeeper.GetChannel(ctx, identifiers[0], identifiers[1])
	if !ok {
		return types2.ConnectionEnd{}, fmt.Errorf("unable to fetch channel info")
	}

	if len(channel.ConnectionHops) != 1 {
		return types2.ConnectionEnd{}, fmt.Errorf("denom has multihop connections")
	}

	conn, ok := k.connKeeper.GetConnection(ctx, channel.ConnectionHops[0])
	if !ok {
		return types2.ConnectionEnd{}, fmt.Errorf("unable to fetch conn info")
	}

	return conn, nil
}

func (k Keeper) CompareClientStates(ctx sdk.Context, cs1, cs2 exported.ClientState) bool {
	cs1 = cs1.ZeroCustomFields()
	cs2 = cs2.ZeroCustomFields()
	x := cs1.ClientType() == cs2.ClientType()
	if x {
		_, ok := cs1.(*types3.ClientState)
		if !ok {
			return false
		}

		return k.VerifyTmClients(cs1, cs2)
	}

	return false
}

func (k Keeper) VerifyTmClients(cs1, cs2 exported.ClientState) bool {
	cs1.(*types3.ClientState).LatestHeight = types4.ZeroHeight()
	cs2.(*types3.ClientState).LatestHeight = types4.ZeroHeight()

	return reflect.DeepEqual(cs1, cs2)
}
