package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mande-labs/mande/v1/x/voting/types"
)

// SetAggregateVoteCount set a specific aggregateVoteCount in the store from its index
func (k Keeper) SetAggregateVoteCount(ctx sdk.Context, aggregateVoteCount types.AggregateVoteCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))
	b := k.cdc.MustMarshal(&aggregateVoteCount)
	store.Set(types.AggregateVoteCountKey(
		aggregateVoteCount.Index,
	), b)
}

// GetAggregateVoteCount returns a aggregateVoteCount from its index
func (k Keeper) GetAggregateVoteCount(
	ctx sdk.Context,
	index string,
) (val types.AggregateVoteCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))

	b := store.Get(types.AggregateVoteCountKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllAggregateVoteCount(ctx sdk.Context) (list []types.AggregateVoteCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AggregateVoteCount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ReconcileAggregatedVotes(msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount, aggregateVoteReceiverCount *types.AggregateVoteCount) {
	voteCount := intAbs(msg.Count)
	switch msg.Mode {
	case 0:
		aggregateVoteCreatorCount.AggregateVotesCasted -= voteCount
		aggregateVoteReceiverCount.AggregateVotesReceived -= voteCount
	case 1:
		aggregateVoteCreatorCount.AggregateVotesCasted += voteCount
		aggregateVoteReceiverCount.AggregateVotesReceived += voteCount
	}
}
