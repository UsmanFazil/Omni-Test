package omni_test

import (
	"testing"

	keepertest "Omni/testutil/keeper"
	"Omni/testutil/nullify"
	"Omni/x/omni"
	"Omni/x/omni/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OmniKeeper(t)
	omni.InitGenesis(ctx, *k, genesisState)
	got := omni.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
