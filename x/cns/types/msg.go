package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var (
	_ sdk.Msg = &MsgRegisterChainName{}
	_ sdk.Msg = &MsgUpdateChainInfo{}
)

const (
	TypeMsgRegisterChainName = "register_name"
	TypeMsgUpdateChainInfo   = "update_info"
)

func NewMsgRegisterChainName(name, owner, seeds, sourceCodeUrl, ibcClient, commitHash, genesisHash string, version int64) *MsgRegisterChainName {
	return &MsgRegisterChainName{
		ChainName:          name,
		Owner:              owner,
		Seed:               strings.Split(seeds, ","),
		SourceCodeUrl:      sourceCodeUrl,
		CanonicalIbcClient: ibcClient,
		Version: &VersionInfo{
			Version:          version,
			SourceCommitHash: commitHash,
			GenesisHash:      genesisHash,
		},
	}
}

func (m *MsgRegisterChainName) Route() string {
	return RouterKey
}

func (m *MsgRegisterChainName) Type() string {
	return TypeMsgRegisterChainName
}

func (m *MsgRegisterChainName) ValidateBasic() error {
	return nil
}

func (m *MsgRegisterChainName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgRegisterChainName) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func NewMsgUpdateChainInfo(name, owner, seeds, sourceCodeUrl, ibcClient, commitHash, genesisHash string, version int64) *MsgUpdateChainInfo {
	return &MsgUpdateChainInfo{
		ChainName:          name,
		Owner:              owner,
		Seed:               strings.Split(seeds, ","),
		SourceCodeUrl:      sourceCodeUrl,
		CanonicalIbcClient: ibcClient,
		Version: &VersionInfo{
			Version:          version,
			SourceCommitHash: commitHash,
			GenesisHash:      genesisHash,
		},
	}
}
func (m *MsgUpdateChainInfo) Route() string {
	return RouterKey
}

func (m *MsgUpdateChainInfo) Type() string {
	return TypeMsgUpdateChainInfo
}

func (m *MsgUpdateChainInfo) ValidateBasic() error {
	return nil
}

func (m *MsgUpdateChainInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUpdateChainInfo) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}
