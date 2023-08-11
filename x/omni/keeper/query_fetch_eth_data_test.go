package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "Omni/testutil/keeper"
	"Omni/testutil/nullify"
	"Omni/x/omni/types"
)

func TestFetchEthDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFetchEthData(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetFetchEthDataRequest
		response *types.QueryGetFetchEthDataResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetFetchEthDataRequest{Id: msgs[0].Id},
			response: &types.QueryGetFetchEthDataResponse{FetchEthData: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetFetchEthDataRequest{Id: msgs[1].Id},
			response: &types.QueryGetFetchEthDataResponse{FetchEthData: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetFetchEthDataRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.FetchEthData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestFetchEthDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFetchEthData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFetchEthDataRequest {
		return &types.QueryAllFetchEthDataRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FetchEthDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FetchEthData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FetchEthData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FetchEthDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FetchEthData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FetchEthData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FetchEthDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.FetchEthData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FetchEthDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
