package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	now := ctx.BlockTime()
	post := types.Post{
		Creator:       msg.Creator,
		Title:         msg.Title,
		Body:          msg.Body,
		CreatedAt:     &now,
		LastUpdatedAt: &now,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
