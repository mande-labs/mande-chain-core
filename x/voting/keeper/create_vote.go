package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mande-labs/mande/v1/x/voting/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreateVote(ctx sdk.Context, msg *types.MsgCreateVote) (error){
	aggregateVoteCreatorCount, found := k.GetAggregateVoteCount(ctx, msg.Creator)
	if !found {
		aggregateVoteCreatorCount.Index = msg.Creator
		aggregateVoteCreatorCount.AggregateVotesCasted = 0
		aggregateVoteCreatorCount.AggregateVotesReceived = 0
	}

	switch msg.Mode {
	case 0:
		if err := k.uncastVote(ctx, msg, &aggregateVoteCreatorCount); err !=nil {
			return err
		}
		break
	case 1:
		if err := k.castVote(ctx, msg, &aggregateVoteCreatorCount); err !=nil {
			return err
		}
		break
	default:
		return sdkerrors.Wrapf(types.ErrInvalidVotingMode, "mode - (%s)", msg.Mode)
	}

	return nil
}

func (k Keeper) uncastVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) (error){
	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVoteRecord, msg.Receiver)
	}

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVotesCasted, msg.Receiver)
	}

	if msg.Count < 0 {
		if voteBookEntry.Negative >= uint64(msg.Count) {
			voteBookEntry.Negative -= uint64(msg.Count)
			k.SetVoteBook(ctx, voteBookEntry)
		} else {
			return sdkerrors.Wrapf(types.ErrNegativeVotesCastedLimit, "count - (%s)", msg.Count)
		}
	} else {
		if voteBookEntry.Positive >= uint64(msg.Count) {
			voteBookEntry.Positive -= uint64(msg.Count)
			k.SetVoteBook(ctx, voteBookEntry)
		} else {
			return sdkerrors.Wrapf(types.ErrPositiveVotesCastedLimit, "count = (%s)", msg.Count)
		}
	}
	
	k.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, &aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)

	return nil
}

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) (error){
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	voteBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()
	voteCastCount := aggregateVoteCreatorCount.AggregateVotesCasted + msg.Count
	if int64(voteBalance) < voteCastCount {
		return sdkerrors.Wrapf(types.ErrNotEnoughMand, "count - (%s)", msg.Count)
	}

	aggregateVoteCreatorCount.AggregateVotesCasted += msg.Count

	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		voteBookEntry.Index = voteBookIndex
		voteBookEntry.Positive = 0
		voteBookEntry.Negative = 0
	}

	if msg.Count < 0 {
		voteBookEntry.Negative += uint64(msg.Count)
	} else {
		voteBookEntry.Positive += uint64(msg.Count)
	}

	k.SetVoteBook(ctx, voteBookEntry)

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		aggregateVoteReceiverCount.Index = msg.Receiver
		aggregateVoteReceiverCount.AggregateVotesCasted = 0
		aggregateVoteReceiverCount.AggregateVotesReceived = 0
	}

	k.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, &aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	
	return nil
}
