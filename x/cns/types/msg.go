package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = &MsgRegisterChainName{}
	_ sdk.Msg = &MsgUpdateChainInfo{}
)

func (m *MsgRegisterChainName) Route() string {
	panic("implement me")
}

func (m *MsgRegisterChainName) Type() string {
	panic("implement me")
}

func (m *MsgRegisterChainName) ValidateBasic() error {
	panic("implement me")
}

func (m *MsgRegisterChainName) GetSignBytes() []byte {
	panic("implement me")
}

func (m *MsgRegisterChainName) GetSigners() []sdk.AccAddress {
	panic("implement me")
}

func (m *MsgUpdateChainInfo) Route() string {
	panic("implement me")
}

func (m *MsgUpdateChainInfo) Type() string {
	panic("implement me")
}

func (m *MsgUpdateChainInfo) ValidateBasic() error {
	panic("implement me")
}

func (m *MsgUpdateChainInfo) GetSignBytes() []byte {
	panic("implement me")
}

func (m *MsgUpdateChainInfo) GetSigners() []sdk.AccAddress {
	panic("implement me")
}
