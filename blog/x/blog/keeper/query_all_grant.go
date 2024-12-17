package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blog/x/blog/types"
)

func (k Keeper) AllGrant(goCtx context.Context, req *types.QueryAllGrantRequest) (*types.QueryAllGrantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	grants, err := k.GetAllGrants(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all grants: %v", err)
	}

	return &types.QueryAllGrantResponse{Grant: grants}, nil
}
