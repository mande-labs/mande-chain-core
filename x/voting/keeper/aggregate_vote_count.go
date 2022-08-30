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
	index string
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

func (k Keeper) ReconcileAggregatedVotes(msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount, aggregateVoteReceiverCount *types.AggregateVoteCount) {
	if msg.Mode == 0 {
		switch operation {
		case 0: // uncast negative votes => ultimately positive votes are considered
			aggregateVoteCreatorCount.AggregateVotesCasted += msg.Count
			aggregateVoteCountOfReceiver.AggregateVotesReceived += msg.Count
		case 1:
			aggregateVoteCreatorCount.AggregateVotesCasted -= msg.Count
			aggregateVoteCountOfReceiver.AggregateVotesReceived -= msg.Count
		}
	} else if msg.Mode == 1 {
		switch operation {
		case 0: // cast negative votes => ultimately negative votes are considered
			aggregateVoteCreatorCount.AggregateVotesCasted -= msg.Count
			aggregateVoteCountOfReceiver.AggregateVotesReceived -= msg.Count
		case 1:
			aggregateVoteCreatorCount.AggregateVotesCasted += msg.Count
			aggregateVoteCountOfReceiver.AggregateVotesReceived += msg.Count
		}
	}
}
