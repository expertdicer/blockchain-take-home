package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func GetGranterBytes(granter string) []byte {
	key := sdk.AppendLengthPrefixedBytes([]byte(granter))

	return key
}

func (k Keeper) SetGrant(ctx sdk.Context, grant types.Grant) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GrantKey))
	b := k.cdc.MustMarshal(&grant)
	store.Set(GetGranterBytes(grant.Granter), b)
}

func (k Keeper) GetGrant(ctx sdk.Context, granter string) (val types.Grant, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GrantKey))
	b := store.Get(GetGranterBytes(granter))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) IsGranteeGrantedForPost(ctx sdk.Context, grantee string, postID uint64) (bool, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GrantKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)

		for _, g := range grant.Grantee {
			if g == grantee && grant.PostId == postID {
				return true, nil
			}
		}
	}

	return false, nil
}

func (k Keeper) GetAllGrants(ctx sdk.Context) ([]types.Grant, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GrantKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var grants []types.Grant
	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)
		grants = append(grants, grant)
	}

	return grants, nil
}
