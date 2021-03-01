package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	"github.com/tendermint/cns/x/cns/types"
	"strings"
)

func (k Keeper) FetchUpdatedBalances(ctx sdk.Context, addr sdk.AccAddress) (sdk.Coins, error) {
	balances := k.bankKeeper.GetAllBalances(ctx, addr)
	var newBalances []sdk.Coin
	for _, bal := range balances {
		if strings.HasPrefix(bal.Denom, "ibc/") {
			denomPath, err := k.ibctransfer.DenomPathFromHash(ctx, bal.Denom)
			if err != nil {
				return nil, err
			}

			denomTrace := transfertypes.ParseDenomTrace(denomPath)
			identifiers := strings.Split(denomTrace.Path, "/")

			channel, ok := k.channelKeeper.GetChannel(ctx, identifiers[0], identifiers[1])
			if !ok {
				return nil, fmt.Errorf("unable to fetch channel info")
			}

			if len(channel.ConnectionHops) != 1 {
				return nil, fmt.Errorf("denom has multihop connections")
			}

			conn, ok := k.connKeeper.GetConnection(ctx, channel.ConnectionHops[0])
			if !ok {
				return nil, fmt.Errorf("unable to fetch conn info")
			}

			cs, ok := k.clientKeeper.GetClientState(ctx, conn.ClientId)
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
					cs0 := cs.ZeroCustomFields()
					cs1 := cs2.ZeroCustomFields()
					if cs0.ClientType() == cs1.ClientType() {
						matchedInfos = append(matchedInfos, info)
						return false
					}
				}
				return false
			})

			if len(matchedInfos) < 1 {
				return nil, fmt.Errorf("no infos matched")
			}

			connCS, ok := k.clientKeeper.GetClientConsensusState(ctx, conn.ClientId, cs.GetLatestHeight())
			if !ok {
				return nil, fmt.Errorf("unable to fetch client consensus state")
			}

			storeCS, ok := k.clientKeeper.GetClientConsensusState(ctx, matchedInfos[0].CanonicalIbcClient, cs2.GetLatestHeight())
			if !ok {
				return nil, fmt.Errorf("unable to fetch client consensus state")
			}

			if connCS.String() != storeCS.String() {
				return nil, fmt.Errorf("consensus state mismatch")
			}
			ibcDenom := "ibc/" + denomTrace.BaseDenom + "/" + matchedInfos[0].ChainName
			newBalances = append(newBalances, sdk.NewCoin(ibcDenom, bal.Amount))
		}
		newBalances = append(newBalances, bal)
	}
	return newBalances, nil
}
