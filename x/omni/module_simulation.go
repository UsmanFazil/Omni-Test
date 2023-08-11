package omni

import (
	"math/rand"

	"Omni/testutil/sample"
	omnisimulation "Omni/x/omni/simulation"
	"Omni/x/omni/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = omnisimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateFetchEthData = "op_weight_msg_fetch_eth_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFetchEthData int = 100

	opWeightMsgUpdateFetchEthData = "op_weight_msg_fetch_eth_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFetchEthData int = 100

	opWeightMsgDeleteFetchEthData = "op_weight_msg_fetch_eth_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFetchEthData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	omniGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		FetchEthDataList: []types.FetchEthData{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		FetchEthDataCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&omniGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateFetchEthData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateFetchEthData, &weightMsgCreateFetchEthData, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFetchEthData = defaultWeightMsgCreateFetchEthData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFetchEthData,
		omnisimulation.SimulateMsgCreateFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFetchEthData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateFetchEthData, &weightMsgUpdateFetchEthData, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFetchEthData = defaultWeightMsgUpdateFetchEthData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFetchEthData,
		omnisimulation.SimulateMsgUpdateFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFetchEthData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteFetchEthData, &weightMsgDeleteFetchEthData, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFetchEthData = defaultWeightMsgDeleteFetchEthData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFetchEthData,
		omnisimulation.SimulateMsgDeleteFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFetchEthData,
			defaultWeightMsgCreateFetchEthData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				omnisimulation.SimulateMsgCreateFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFetchEthData,
			defaultWeightMsgUpdateFetchEthData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				omnisimulation.SimulateMsgUpdateFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFetchEthData,
			defaultWeightMsgDeleteFetchEthData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				omnisimulation.SimulateMsgDeleteFetchEthData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
