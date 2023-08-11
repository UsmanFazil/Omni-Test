package keeper

import (
	"context"

	"Omni/x/omni/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FetchEthDataAll(goCtx context.Context, req *types.QueryAllFetchEthDataRequest) (*types.QueryAllFetchEthDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fetchEthDatas []types.FetchEthData
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	fetchEthDataStore := prefix.NewStore(store, types.KeyPrefix(types.FetchEthDataKey))

	pageRes, err := query.Paginate(fetchEthDataStore, req.Pagination, func(key []byte, value []byte) error {
		var fetchEthData types.FetchEthData
		if err := k.cdc.Unmarshal(value, &fetchEthData); err != nil {
			return err
		}

		fetchEthDatas = append(fetchEthDatas, fetchEthData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFetchEthDataResponse{FetchEthData: fetchEthDatas, Pagination: pageRes}, nil
}

func (k Keeper) FetchEthData(goCtx context.Context, req *types.QueryGetFetchEthDataRequest) (*types.QueryGetFetchEthDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	fetchEthData, found := k.GetFetchEthData(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFetchEthDataResponse{FetchEthData: fetchEthData}, nil
}
