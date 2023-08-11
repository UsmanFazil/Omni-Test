package keeper

import (
	"math/big"
	"strings"
	"strconv"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
	"encoding/binary"

	"Omni/x/omni/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetFetchEthDataCount get the total number of fetchEthData
func (k Keeper) GetFetchEthDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FetchEthDataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFetchEthDataCount set the total number of fetchEthData
func (k Keeper) SetFetchEthDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FetchEthDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFetchEthData appends a fetchEthData in the store with a new id and update the count
func (k Keeper) AppendFetchEthData(
	ctx sdk.Context,
	fetchEthData types.FetchEthData,
) uint64 {

	// Connect to a local Ethereum node
	rpcEndpoint := os.Getenv("ETHEREUM_RPC_ENDPOINT")
	client, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		return 0;
	}

	// Contract address and storage position we want to query
	contractAddress := "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984" //uni token address
	storagePosition := "0x0" // zero index

	var result string
	err = client.Call(&result, "eth_getStorageAt", contractAddress, storagePosition, "latest")
	if err != nil {
		return 0;
	}

	// Convert the result to a big.Int for easier handling
	value := new(big.Int)
	value.SetString(strings.TrimPrefix(result, "0x"), 16)

	fetchEthData.Id , _ = strconv.ParseUint(fetchEthData.DataVal, 10, 0);
	fetchEthData.DataVal = value.String();

	// Create the fetchEthData
	count := k.GetFetchEthDataCount(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FetchEthDataKey))
	appendedValue := k.cdc.MustMarshal(&fetchEthData)
	store.Set(GetFetchEthDataIDBytes(fetchEthData.Id), appendedValue)

	// Update fetchEthData count
	k.SetFetchEthDataCount(ctx, count+1)

	return count
}

// SetFetchEthData set a specific fetchEthData in the store
func (k Keeper) SetFetchEthData(ctx sdk.Context, fetchEthData types.FetchEthData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FetchEthDataKey))
	b := k.cdc.MustMarshal(&fetchEthData)
	store.Set(GetFetchEthDataIDBytes(fetchEthData.Id), b)
}

// GetFetchEthData returns a fetchEthData from its id
func (k Keeper) GetFetchEthData(ctx sdk.Context, id uint64) (val types.FetchEthData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FetchEthDataKey))
	b := store.Get(GetFetchEthDataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFetchEthData removes a fetchEthData from the store
func (k Keeper) RemoveFetchEthData(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FetchEthDataKey))
	store.Delete(GetFetchEthDataIDBytes(id))
}

// GetAllFetchEthData returns all fetchEthData
func (k Keeper) GetAllFetchEthData(ctx sdk.Context) (list []types.FetchEthData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FetchEthDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FetchEthData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFetchEthDataIDBytes returns the byte representation of the ID
func GetFetchEthDataIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFetchEthDataIDFromBytes returns ID in uint64 format from a byte array
func GetFetchEthDataIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
