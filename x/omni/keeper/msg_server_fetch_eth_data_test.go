package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"Omni/x/omni/types"
)

func TestFetchEthDataMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateFetchEthData(ctx, &types.MsgCreateFetchEthData{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestFetchEthDataMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateFetchEthData
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateFetchEthData{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateFetchEthData{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateFetchEthData{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateFetchEthData(ctx, &types.MsgCreateFetchEthData{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateFetchEthData(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestFetchEthDataMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteFetchEthData
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteFetchEthData{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteFetchEthData{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteFetchEthData{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateFetchEthData(ctx, &types.MsgCreateFetchEthData{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteFetchEthData(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
