package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) RevokePermission(goCtx context.Context, msg *types.MsgRevokePermission) (*types.MsgRevokePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	if post.Creator != msg.Granter {
		return nil, sdkerrors.ErrUnauthorized
	}

	grant, found := k.GetGrant(ctx, msg.Granter)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Granter))
	}

	exist := grantExists(grant.Grantee, msg.Grantee)
	if !exist {
		return nil, fmt.Errorf("Grantee %s has been granted", msg.Grantee)
	}
	grant.Grantee = removeGrant(grant.Grantee, msg.Grantee)
	k.SetGrant(ctx, grant)
	return &types.MsgRevokePermissionResponse{}, nil
}

func removeGrant(slice []string, str string) []string {
	newSlice := []string{}
	for _, s := range slice {
		if s != str {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}
