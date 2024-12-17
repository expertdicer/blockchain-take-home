package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) GrantPermission(goCtx context.Context, msg *types.MsgGrantPermission) (*types.MsgGrantPermissionResponse, error) {
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
		k.SetGrant(ctx, types.Grant{
			Granter: msg.Granter,
			Grantee: []string{msg.Grantee},
			PostId:  msg.Id,
		})
		return &types.MsgGrantPermissionResponse{}, nil
	}

	fmt.Println("grant.Grantee :", grant.Grantee)
	fmt.Println("msg.Grantee :", msg.Grantee)

	exist := grantExists(grant.Grantee, msg.Grantee)
	if exist {
		return nil, fmt.Errorf("Grantee %s has been granted", msg.Grantee)
	}
	grant.Grantee = append(grant.Grantee, msg.Grantee)
	k.SetGrant(ctx, grant)
	return &types.MsgGrantPermissionResponse{}, nil
}

func grantExists(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
