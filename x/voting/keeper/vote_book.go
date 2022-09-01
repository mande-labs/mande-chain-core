package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mande-labs/mande/v1/x/voting/types"
)

// SetVoteBook set a specific voteBook in the store from its index
func (k Keeper) SetVoteBook(ctx sdk.Context, voteBook types.VoteBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	b := k.cdc.MustMarshal(&voteBook)
	store.Set(types.VoteBookKey(
		voteBook.Index,
	), b)
}

// GetVoteBook returns a voteBook from its index
func (k Keeper) GetVoteBook(
	ctx sdk.Context,
	index string,
) (val types.VoteBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))

	b := store.Get(types.VoteBookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllVoteBook(ctx sdk.Context) (list []types.VoteBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VoteBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
