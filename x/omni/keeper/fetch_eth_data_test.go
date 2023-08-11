package keeper_test

import (
	"testing"

	keepertest "Omni/testutil/keeper"
	"Omni/testutil/nullify"
	"Omni/x/omni/keeper"
	"Omni/x/omni/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNFetchEthData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FetchEthData {
	items := make([]types.FetchEthData, n)
	for i := range items {
		items[i].Id = keeper.AppendFetchEthData(ctx, items[i])
	}
	return items
}

func TestFetchEthDataGet(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNFetchEthData(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFetchEthData(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFetchEthDataRemove(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNFetchEthData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFetchEthData(ctx, item.Id)
		_, found := keeper.GetFetchEthData(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFetchEthDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNFetchEthData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFetchEthData(ctx)),
	)
}

func TestFetchEthDataCount(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNFetchEthData(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFetchEthDataCount(ctx))
}
