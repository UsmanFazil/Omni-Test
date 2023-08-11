package keeper

import (
	"context"
	"fmt"

	"Omni/x/omni/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFetchEthData(goCtx context.Context, msg *types.MsgCreateFetchEthData) (*types.MsgCreateFetchEthDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var fetchEthData = types.FetchEthData{
		Creator: msg.Creator,
		DataVal: msg.DataVal,
	}

	id := k.AppendFetchEthData(
		ctx,
		fetchEthData,
	)

	return &types.MsgCreateFetchEthDataResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateFetchEthData(goCtx context.Context, msg *types.MsgUpdateFetchEthData) (*types.MsgUpdateFetchEthDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var fetchEthData = types.FetchEthData{
		Creator: msg.Creator,
		Id:      msg.Id,
		DataVal: msg.DataVal,
	}

	// Checks that the element exists
	val, found := k.GetFetchEthData(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetFetchEthData(ctx, fetchEthData)

	return &types.MsgUpdateFetchEthDataResponse{}, nil
}

func (k msgServer) DeleteFetchEthData(goCtx context.Context, msg *types.MsgDeleteFetchEthData) (*types.MsgDeleteFetchEthDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetFetchEthData(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFetchEthData(ctx, msg.Id)

	return &types.MsgDeleteFetchEthDataResponse{}, nil
}
