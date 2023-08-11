package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateFetchEthData = "create_fetch_eth_data"
	TypeMsgUpdateFetchEthData = "update_fetch_eth_data"
	TypeMsgDeleteFetchEthData = "delete_fetch_eth_data"
)

var _ sdk.Msg = &MsgCreateFetchEthData{}

func NewMsgCreateFetchEthData(creator string, dataVal string) *MsgCreateFetchEthData {
	return &MsgCreateFetchEthData{
		Creator: creator,
		DataVal: dataVal,
	}
}

func (msg *MsgCreateFetchEthData) Route() string {
	return RouterKey
}

func (msg *MsgCreateFetchEthData) Type() string {
	return TypeMsgCreateFetchEthData
}

func (msg *MsgCreateFetchEthData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateFetchEthData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateFetchEthData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFetchEthData{}

func NewMsgUpdateFetchEthData(creator string, id uint64, dataVal string) *MsgUpdateFetchEthData {
	return &MsgUpdateFetchEthData{
		Id:      id,
		Creator: creator,
		DataVal: dataVal,
	}
}

func (msg *MsgUpdateFetchEthData) Route() string {
	return RouterKey
}

func (msg *MsgUpdateFetchEthData) Type() string {
	return TypeMsgUpdateFetchEthData
}

func (msg *MsgUpdateFetchEthData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateFetchEthData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateFetchEthData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFetchEthData{}

func NewMsgDeleteFetchEthData(creator string, id uint64) *MsgDeleteFetchEthData {
	return &MsgDeleteFetchEthData{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteFetchEthData) Route() string {
	return RouterKey
}

func (msg *MsgDeleteFetchEthData) Type() string {
	return TypeMsgDeleteFetchEthData
}

func (msg *MsgDeleteFetchEthData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteFetchEthData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteFetchEthData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
