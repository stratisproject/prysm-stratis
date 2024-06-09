package slasher

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	logTest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stratisproject/prysm-stratis/async/event"
	mock "github.com/stratisproject/prysm-stratis/beacon-chain/blockchain/testing"
	dbtest "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	"github.com/stratisproject/prysm-stratis/beacon-chain/startup"
	mockSync "github.com/stratisproject/prysm-stratis/beacon-chain/sync/initial-sync/testing"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
	"github.com/stratisproject/prysm-stratis/time/slots"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)

	m.Run()
}

func TestService_StartStop_ChainInitialized(t *testing.T) {
	slasherDB := dbtest.SetupSlasherDB(t)
	hook := logTest.NewGlobal()
	beaconState, err := util.NewBeaconState()
	require.NoError(t, err)
	currentSlot := primitives.Slot(4)
	require.NoError(t, beaconState.SetSlot(currentSlot))
	mockChain := &mock.ChainService{
		State: beaconState,
		Slot:  &currentSlot,
	}
	gs := startup.NewClockSynchronizer()
	srv, err := New(context.Background(), &ServiceConfig{
		IndexedAttestationsFeed: new(event.Feed),
		BeaconBlockHeadersFeed:  new(event.Feed),
		StateNotifier:           &mock.MockStateNotifier{},
		Database:                slasherDB,
		HeadStateFetcher:        mockChain,
		SyncChecker:             &mockSync.Sync{IsSyncing: false},
		ClockWaiter:             gs,
	})
	require.NoError(t, err)
	go srv.Start()
	time.Sleep(time.Millisecond * 100)
	var vr [32]byte
	require.NoError(t, gs.SetClock(startup.NewClock(time.Now(), vr)))
	time.Sleep(time.Millisecond * 100)
	srv.attsSlotTicker = &slots.SlotTicker{}
	srv.blocksSlotTicker = &slots.SlotTicker{}
	srv.pruningSlotTicker = &slots.SlotTicker{}
	require.NoError(t, srv.Stop())
	require.NoError(t, srv.Status())
	require.LogsContain(t, hook, "received chain initialization")
}
