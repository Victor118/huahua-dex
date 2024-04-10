package testutil

import (
	"encoding/json"

	cometbfttypes "github.com/cometbft/cometbft/abci/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/stretchr/testify/suite"

	"github.com/ChihuahuaChain/chihuahua/app"
)

var (
	// TestOwnerAddress defines a reusable bech32 address for testing purposes
	TestOwnerAddress = "neutron17dtl0mjt3t77kpuhg2edqzjpszulwhgzcdvagh"

	TestInterchainID = "owner_id"

	// provider-consumer connection takes connection-0
	ConnectionOne = "connection-1"

	// TestVersion defines a reusable interchainaccounts version string for testing purposes
	TestVersion = string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: ConnectionOne,
		HostConnectionId:       ConnectionOne,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))
)

func init() {
	// ibctesting.DefaultTestingAppInit = SetupTestingApp()
	app.GetDefaultConfig()
}

type IBCConnectionTestSuite struct {
	suite.Suite
	Coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	ChainProvider *ibctesting.TestChain
	ChainA        *ibctesting.TestChain
	ChainB        *ibctesting.TestChain
}

// SetupTestingApp initializes the IBC-go testing application
func SetupTestingApp(initValUpdates []cometbfttypes.ValidatorUpdate) func() (ibctesting.TestingApp, map[string]json.RawMessage) {
	return func() (ibctesting.TestingApp, map[string]json.RawMessage) {
		encoding := app.MakeEncodingConfig()
		db := dbm.NewMemDB()
		testApp := app.New(
			log.NewNopLogger(),
			db,
			nil,
			false,
			map[int64]bool{},
			app.DefaultNodeHome,
			0,
			app.GetEnabledProposals(),
			sims.EmptyAppOptions{},
			nil,
		)

		// we need to set up a TestInitChainer where we can redefine MaxBlockGas in ConsensusParamsKeeper
		testApp.SetInitChainer(testApp.TestInitChainer)
		// and then we manually init baseapp and load states
		testApp.LoadLatest()

		genesisState := app.NewDefaultGenesisState(testApp.AppCodec())

		// NOTE ibc-go/v7/testing.SetupWithGenesisValSet requires a staking module
		// genesisState or it panics. Feed a minimum one.
		genesisState[stakingtypes.ModuleName] = encoding.Marshaler.MustMarshalJSON(
			&stakingtypes.GenesisState{
				Params: stakingtypes.Params{BondDenom: sdk.DefaultBondDenom},
			},
		)

		return testApp, genesisState
	}
}
