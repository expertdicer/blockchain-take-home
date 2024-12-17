package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	now := ctx.BlockTime()
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	isGranted, err := k.IsGranteeGrantedForPost(ctx, msg.Creator, msg.Id)
	if err != nil {
		return nil, fmt.Errorf("Error determining grant for %v: %v", msg.Creator, err)
	}
	if msg.Creator != val.Creator && !isGranted {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	post := types.Post{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Title:         msg.Title,
		Body:          msg.Body,
		CreatedAt:     val.CreatedAt,
		LastUpdatedAt: &now,
	}

	k.SetPost(ctx, post)
	return &types.MsgUpdatePostResponse{}, nil
}
