package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	mandeapp "github.com/mande-labs/mande/v1/app"
	"github.com/mande-labs/mande/v1/x/voting/keeper"
	"github.com/mande-labs/mande/v1/x/voting/types"
)

const (
	mandDenom  = "mand"
	stakeDenom = "stake"
)

func newMandCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(mandDenom, amt)
}

func newStakeCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(stakeDenom, amt)
}

type IntegrationTestSuite struct {
	suite.Suite

	app         *mandeapp.MandeApp
	ctx         sdk.Context
	queryClient banktypes.QueryClient

	msgServer types.MsgServer
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := mandeapp.Setup(suite.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	banktypes.RegisterQueryServer(queryHelper, app.BankKeeper)
	queryClient := banktypes.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx
	suite.queryClient = queryClient
	suite.msgServer = keeper.NewMsgServerImpl(suite.app.VotingKeeper)
}

func (suite *IntegrationTestSuite) TestCreateVote_Positive_Cast_Valid() {
	app, ctx, msgServer, keeper := suite.app, suite.ctx, suite.msgServer, suite.app.VotingKeeper
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)

	// ============= Case: 1 => addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1,
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver.AggregateVotesCasted)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.Require().Equal(uint64(100), voteBookEntry.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry.Negative)

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)

	// ============= Case: 2 => addr1 votes addr3 ===============
	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr3.String(),
		Count:    100,
		Mode:     1,
	}
	_, err1 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().NoError(err1)

	// validate aggregate vote counts
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver2, _ := keeper.GetAggregateVoteCount(ctx, addr3.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator1.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator1.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver2.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver2.AggregateVotesCasted)

	// validate vote book
	voteBookEntry1, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr3.String()))
	suite.Require().Equal(uint64(100), voteBookEntry1.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry1.Negative)

	// ============= Case: 3 => addr2 votes addr1 ===============
	sendAmt := sdk.NewCoins(newMandCoin(100))
	suite.Require().NoError(app.BankKeeper.SendCoins(ctx, addr1, addr2, sendAmt))
	msgCreateVote2 := &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     1,
	}
	_, err2 := msgServer.CreateVote(ctx, msgCreateVote2)
	suite.Require().NoError(err2)

	// validate aggregate vote counts
	aggregateVoteCountCreator3, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver4, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator3.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator3.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver4.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver4.AggregateVotesCasted)

	// validate vote book
	voteBookEntry2, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.Require().Equal(uint64(100), voteBookEntry2.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry2.Negative)
}

func (suite *IntegrationTestSuite) TestCreateVote_Positive_Cast_Invalid() {
	app, ctx, msgServer := suite.app, suite.ctx, suite.msgServer
	balances := sdk.NewCoins(newStakeCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)

	// ============= Case: 1 => addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1,
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().Error(err, "not enough MAND to use for voting")
}

func (suite *IntegrationTestSuite) TestCreateVote_Negative_Cast_Valid() {
	app, ctx, msgServer, keeper := suite.app, suite.ctx, suite.msgServer, suite.app.VotingKeeper
	balances := sdk.NewCoins(newMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)

	// ============= Case: 1 => addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver.AggregateVotesCasted)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.Require().Equal(uint64(0), voteBookEntry.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry.Negative)

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)

	// ============= Case: 2 => addr1 votes addr3 ===============
	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr3.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err1 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().NoError(err1)

	// validate aggregate vote counts
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver2, _ := keeper.GetAggregateVoteCount(ctx, addr3.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator1.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator1.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver2.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver2.AggregateVotesCasted)

	// validate vote book
	voteBookEntry1, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr3.String()))
	suite.Require().Equal(uint64(0), voteBookEntry1.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry1.Negative)

	// ============= Case: 3 => addr2 votes addr1 ===============
	sendAmt := sdk.NewCoins(newMandCoin(100))
	suite.Require().NoError(app.BankKeeper.SendCoins(ctx, addr1, addr2, sendAmt))
	msgCreateVote2 := &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err2 := msgServer.CreateVote(ctx, msgCreateVote2)
	suite.Require().NoError(err2)

	// validate aggregate vote counts
	aggregateVoteCountCreator3, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver4, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator3.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator3.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver4.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver4.AggregateVotesCasted)

	// validate vote book
	voteBookEntry2, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.Require().Equal(uint64(0), voteBookEntry2.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry2.Negative)
}

func (suite *IntegrationTestSuite) TestCreateVote_Negative_Cast_Invalid() {
	app, ctx, msgServer := suite.app, suite.ctx, suite.msgServer

	// =============== Case: 1 => addr1 votes addr2 ===============
	balances := sdk.NewCoins(newStakeCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)

	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().Error(err, "not enough MAND to use for voting")

	// =============== Case: 2 => addr3 votes addr4 ===============
	balances2 := sdk.NewCoins(newMandCoin(99))

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr3, balances2))

	addr4 := sdk.AccAddress("addr4_______________")
	acc4 := app.AccountKeeper.NewAccountWithAddress(ctx, addr4)
	app.AccountKeeper.SetAccount(ctx, acc4)

	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr3.String(),
		Receiver: addr4.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err1 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().Error(err1, "not enough MAND to use for voting")
}

func (suite *IntegrationTestSuite) TestCreateVote_Positive_Uncast_Valid() {
	app, ctx, msgServer, keeper := suite.app, suite.ctx, suite.msgServer, suite.app.VotingKeeper
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr2, balances))

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)

	// ============= cast vote in order to uncast ===============
	_, err := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    200,
		Mode:     1,
	})
	suite.Require().NoError(err)

	_, err1 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr3.String(),
		Count:    100,
		Mode:     1,
	})
	suite.Require().NoError(err1)

	// ============= Case: 1 => addr1 uncasts votes from addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     0,
	}
	_, err2 := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err2)

	// validate aggregate vote counts
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver1, _ := keeper.GetAggregateVoteCount(ctx, addr3.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver1.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver1.AggregateVotesCasted)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	voteBookEntry1, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr3.String()))
	suite.Require().Equal(uint64(100), voteBookEntry.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry.Negative)
	suite.Require().Equal(uint64(100), voteBookEntry1.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry.Negative)

	// ============= cast vote in order to uncast ===============
	_, err3 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    300,
		Mode:     1,
	})
	suite.Require().NoError(err3)

	// ============= Case: 2 => addr2 uncasts votes from addr1 ===============
	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     0,
	}
	_, err4 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().NoError(err4)

	// validate aggregate vote counts
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver2, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator1.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator1.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver2.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver2.AggregateVotesCasted)

	// validate vote book
	voteBookEntry2, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.Require().Equal(uint64(200), voteBookEntry2.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry2.Negative)
}

func (suite *IntegrationTestSuite) TestCreateVote_Positive_Uncast_Invalid() {
	app, ctx, msgServer := suite.app, suite.ctx, suite.msgServer
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr2, balances))

	// ============= Case: 1 => addr2 uncasts votes from addr1 =============
	_, err := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     0,
	})
	suite.Require().Error(err, "voting record not found")

	// ============= cast vote in order to uncast ===============
	_, err1 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    200,
		Mode:     1,
	})
	suite.Require().NoError(err1)

	_, err2 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    300,
		Mode:     0,
	})
	suite.Require().Error(err2, "you have not casted this many positive votes")
}

func (suite *IntegrationTestSuite) TestCreateVote_Negative_Uncast_Valid() {
	app, ctx, msgServer, keeper := suite.app, suite.ctx, suite.msgServer, suite.app.VotingKeeper
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr2, balances))

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)

	// ============= cast vote in order to uncast ===============
	_, err := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -200,
		Mode:     1,
	})
	suite.Require().NoError(err)

	_, err1 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr3.String(),
		Count:    -100,
		Mode:     1,
	})
	suite.Require().NoError(err1)

	// ============= Case: 1 => addr1 uncasts votes from addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     0,
	}
	_, err2 := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err2)

	// validate aggregate vote counts
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver1, _ := keeper.GetAggregateVoteCount(ctx, addr3.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver1.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver1.AggregateVotesCasted)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	voteBookEntry1, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr3.String()))
	suite.Require().Equal(uint64(0), voteBookEntry.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry.Negative)
	suite.Require().Equal(uint64(0), voteBookEntry1.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry.Negative)

	// ============= cast vote in order to uncast ===============
	_, err3 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    -300,
		Mode:     1,
	})
	suite.Require().NoError(err3)

	// ============= Case: 2 => addr2 uncasts votes from addr1 ===============
	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     0,
	}
	_, err4 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().NoError(err4)

	// validate aggregate vote counts
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver2, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator1.AggregateVotesCasted)
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator1.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver2.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver2.AggregateVotesCasted)

	// validate vote book
	voteBookEntry2, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.Require().Equal(uint64(0), voteBookEntry2.Positive)
	suite.Require().Equal(uint64(200), voteBookEntry2.Negative)
}

func (suite *IntegrationTestSuite) TestCreateVote_Negative_Uncast_Invalid() {
	app, ctx, msgServer := suite.app, suite.ctx, suite.msgServer
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr2, balances))

	// ============= Case: 1 => addr2 uncasts votes from addr1 =============
	_, err := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     0,
	})
	suite.Require().Error(err, "voting record not found")

	// ============= cast vote in order to uncast ===============
	_, err1 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -200,
		Mode:     1,
	})
	suite.Require().NoError(err1)

	_, err2 := msgServer.CreateVote(ctx, &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -300,
		Mode:     0,
	})
	suite.Require().Error(err2, "you have not casted this many negative votes")
}

func (suite *IntegrationTestSuite) TestCreateVote_Positive_Negative_Cast_Valid() {
	app, ctx, msgServer, keeper := suite.app, suite.ctx, suite.msgServer, suite.app.VotingKeeper
	balances := sdk.NewCoins(newMandCoin(300))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	app.AccountKeeper.SetAccount(ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	app.AccountKeeper.SetAccount(ctx, acc2)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, addr2, balances))

	// ============= Case: 1 => addr1 positive votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    200,
		Mode:     1,
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator.AggregateVotesReceived)
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver.AggregateVotesCasted)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.Require().Equal(uint64(200), voteBookEntry.Positive)
	suite.Require().Equal(uint64(0), voteBookEntry.Negative)

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(ctx, addr3)
	app.AccountKeeper.SetAccount(ctx, acc3)

	// ============= Case: 2 => addr1 negative votes addr3 ===============
	msgCreateVote1 := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr3.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err1 := msgServer.CreateVote(ctx, msgCreateVote1)
	suite.Require().NoError(err1)

	// validate aggregate vote counts
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	aggregateVoteCountReceiver2, _ := keeper.GetAggregateVoteCount(ctx, addr3.String())
	suite.Require().Equal(uint64(300), aggregateVoteCountCreator1.AggregateVotesCasted)
	suite.Require().Equal(uint64(0), aggregateVoteCountCreator1.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver2.AggregateVotesReceived)
	suite.Require().Equal(uint64(0), aggregateVoteCountReceiver2.AggregateVotesCasted)

	// validate vote book
	voteBookEntry1, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr1.String(), addr3.String()))
	suite.Require().Equal(uint64(0), voteBookEntry1.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry1.Negative)

	// ============= Case: 3 => addr2 negative votes addr1 ===============
	sendAmt := sdk.NewCoins(newMandCoin(100))
	suite.Require().NoError(app.BankKeeper.SendCoins(ctx, addr1, addr2, sendAmt))
	msgCreateVote2 := &types.MsgCreateVote{
		Creator:  addr2.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     1,
	}
	_, err2 := msgServer.CreateVote(ctx, msgCreateVote2)
	suite.Require().NoError(err2)

	// validate aggregate vote counts
	aggregateVoteCountCreator3, _ := keeper.GetAggregateVoteCount(ctx, addr2.String())
	aggregateVoteCountReceiver4, _ := keeper.GetAggregateVoteCount(ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator3.AggregateVotesCasted)
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator3.AggregateVotesReceived)
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver4.AggregateVotesReceived)
	suite.Require().Equal(uint64(300), aggregateVoteCountReceiver4.AggregateVotesCasted)

	// validate vote book
	voteBookEntry2, _ := keeper.GetVoteBook(ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.Require().Equal(uint64(0), voteBookEntry2.Positive)
	suite.Require().Equal(uint64(100), voteBookEntry2.Negative)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
