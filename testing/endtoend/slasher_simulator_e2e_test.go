package endtoend

import (
	"context"
	"os"
	"strconv"
	"testing"

	logTest "github.com/sirupsen/logrus/hooks/test"
	mock "github.com/stratisproject/prysm-stratis/beacon-chain/blockchain/testing"
	dbtest "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	mockslashings "github.com/stratisproject/prysm-stratis/beacon-chain/operations/slashings/mock"
	"github.com/stratisproject/prysm-stratis/beacon-chain/startup"
	mockstategen "github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen/mock"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/crypto/bls"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/require"
	slashersimulator "github.com/stratisproject/prysm-stratis/testing/slasher/simulator"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

type mockSyncChecker struct{}

func (c mockSyncChecker) Initialized() bool {
	return true
}

func (c mockSyncChecker) Syncing() bool {
	return false
}

func (c mockSyncChecker) Synced() bool {
	return true
}

func (c mockSyncChecker) Status() error {
	return nil
}

func (c mockSyncChecker) Resync() error {
	return nil
}

func (mockSyncChecker) IsSynced(_ context.Context) (bool, error) {
	return true, nil
}

func TestEndToEnd_SlasherSimulator(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	params.OverrideBeaconConfig(params.E2ETestConfig().Copy())

	hook := logTest.NewGlobal()
	ctx := context.Background()

	// Run for 10 epochs if not in long-running to confirm long-running has no issues.
	simulatorParams := slashersimulator.DefaultParams()
	var err error
	epochStr, longRunning := os.LookupEnv("E2E_EPOCHS")
	if longRunning {
		epochsToRun, err := strconv.Atoi(epochStr)
		require.NoError(t, err)
		simulatorParams.NumEpochs = uint64(epochsToRun)
	}

	slasherDB := dbtest.SetupSlasherDB(t)
	beaconState, err := util.NewBeaconState()
	require.NoError(t, err)

	// We setup validators in the beacon state along with their
	// private keys used to generate valid signatures in generated objects.
	validators := make([]*ethpb.Validator, simulatorParams.NumValidators)
	privKeys := make(map[primitives.ValidatorIndex]bls.SecretKey)
	for valIdx := range validators {
		privKey, err := bls.RandKey()
		require.NoError(t, err)
		privKeys[primitives.ValidatorIndex(valIdx)] = privKey
		validators[valIdx] = &ethpb.Validator{
			PublicKey:             privKey.PublicKey().Marshal(),
			WithdrawalCredentials: make([]byte, 32),
		}
	}
	err = beaconState.SetValidators(validators)
	require.NoError(t, err)

	mockChain := &mock.ChainService{State: beaconState}
	gen := mockstategen.NewService()
	gen.AddStateForRoot(beaconState, [32]byte{})

	gs := startup.NewClockSynchronizer()
	sim, err := slashersimulator.New(ctx, &slashersimulator.ServiceConfig{
		Params:                      simulatorParams,
		Database:                    slasherDB,
		StateNotifier:               &mock.MockStateNotifier{},
		HeadStateFetcher:            mockChain,
		AttestationStateFetcher:     mockChain,
		StateGen:                    gen,
		PrivateKeysByValidatorIndex: privKeys,
		SlashingsPool:               &mockslashings.PoolMock{},
		SyncChecker:                 mockSyncChecker{},
		ClockWaiter:                 gs,
		ClockSetter:                 gs,
	})
	require.NoError(t, err)
	sim.Start()
	err = sim.Stop()
	require.NoError(t, err)
	require.LogsDoNotContain(t, hook, "ERROR")
	require.LogsDoNotContain(t, hook, "Did not detect")
	require.LogsContain(t, hook, "Correctly detected simulated proposer slashing")
	require.LogsContain(t, hook, "Correctly detected simulated attester slashing")
}
