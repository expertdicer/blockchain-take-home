package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blog/x/blog/types"
)

func (k Keeper) IsGranted(goCtx context.Context, req *types.QueryGrantRequest) (*types.QueryGrantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	isGranted, err := k.IsGranteeGrantedForPost(ctx, req.Grantee, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Error determining grant for %v: %v", req.Grantee, err)
	}

	return &types.QueryGrantResponse{IsGranted: isGranted}, nil
}
