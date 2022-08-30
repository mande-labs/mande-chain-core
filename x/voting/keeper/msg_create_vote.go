package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mande-labs/mande/v1/x/voting/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreateVote(ctx sdk.Context, msgCreateVote *types.MsgCreateVote) (error){
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	aggregateVoteCreatorCount, found := k.Keeper.GetAggregateVoteCount(ctx, creator)
	if !found {
		aggregateVoteCreatorCount.Index = creator
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
		return sdkerrors.Wrap(types.ErrInvalidVotingMode, msg.Mode)
	}

	return nil
}

func (k Keeper) uncastVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) (error){
	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.Keeper.GetVoteBook(ctx, voteBookIndex)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVoteRecord, msg.Receiver)
	}

	aggregateVoteReceiverCount, found := k.Keeper.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVotesCasted, msg.Receiver)
	}

	switch msg.Operation {
	case 0:
		if voteBookEntry.Negative >= msg.Count {
			voteBookEntry.Negative -= msg.Count
			k.Keeper.SetVoteBook(ctx, voteBookEntry)
		} else {
			return sdkerrors.Wrap(types.ErrNegativeVotesCastedLimit, msg.Count)
		}
	case 1:
		if voteBookEntry.Positive >= msg.Count {
			voteBookEntry.Positive -= msg.Count
			k..Keeper.SetVoteBook(ctx, voteBookEntry)
		} else {
			return sdkerrors.Wrap(types.ErrPositiveVotesCastedLimit, msg.Count)
		}
	default:
		return sdkerrors.Wrap(types.ErrInvalidVotingOperation, msg.Operation)
	}
	
	k.Keeper.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, *aggregateVoteReceiverCount)
	k.Keeper.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)

	return nil
}

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) (error){
	voteBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()
	voteCastCount := aggregateVoteCreatorCount.AggregateVotesCasted + msg.Count
	if voteBalance < voteCastCount {
		return sdkerrors.Wrap(types.ErrNotEnoughMand, msg.Count)
	}

	aggregateVoteCreatorCount.AggregateVotesCasted += msg.Count

	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		voteBookEntry.Index = voteBookIndex
		voteBookEntry.Positive = 0
		voteBookEntry.Negative = 0
	}

	switch msg.Operation {
	case 0:
		voteBookEntry.Negative += msg.Count
	case 1:
		voteBookEntry.Positive += msg.Count
	}

	k.SetVoteBook(ctx, voteBookEntry)

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		aggregateVoteReceiverCount.Index = msg.Receiver
		aggregateVoteReceiverCount.AggregateVotesCasted = 0
		aggregateVoteReceiverCount.AggregateVotesReceived = 0
	}

	k.Keeper.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, *aggregateVoteReceiverCount)
	k.Keeper.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	
	return nil
}
