package omni

import (
	"Omni/x/omni/keeper"
	"Omni/x/omni/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the fetchEthData
	for _, elem := range genState.FetchEthDataList {
		k.SetFetchEthData(ctx, elem)
	}

	// Set fetchEthData count
	k.SetFetchEthDataCount(ctx, genState.FetchEthDataCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FetchEthDataList = k.GetAllFetchEthData(ctx)
	genesis.FetchEthDataCount = k.GetFetchEthDataCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
