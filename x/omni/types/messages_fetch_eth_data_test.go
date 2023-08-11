package types

import (
	"testing"

	"Omni/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFetchEthData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFetchEthData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFetchEthData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateFetchEthData{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateFetchEthData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFetchEthData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateFetchEthData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateFetchEthData{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteFetchEthData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteFetchEthData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteFetchEthData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteFetchEthData{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
