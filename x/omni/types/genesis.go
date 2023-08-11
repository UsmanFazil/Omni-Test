package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FetchEthDataList: []FetchEthData{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in fetchEthData
	fetchEthDataIdMap := make(map[uint64]bool)
	fetchEthDataCount := gs.GetFetchEthDataCount()
	for _, elem := range gs.FetchEthDataList {
		if _, ok := fetchEthDataIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for fetchEthData")
		}
		if elem.Id >= fetchEthDataCount {
			return fmt.Errorf("fetchEthData id should be lower or equal than the last id")
		}
		fetchEthDataIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
