package sync

import (
	"context"
	"testing"
	"time"

	logTest "github.com/sirupsen/logrus/hooks/test"
	mockChain "github.com/stratisproject/prysm-stratis/beacon-chain/blockchain/testing"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/signing"
	testingdb "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	doublylinkedtree "github.com/stratisproject/prysm-stratis/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/blstoexec"
	mockp2p "github.com/stratisproject/prysm-stratis/beacon-chain/p2p/testing"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	mockSync "github.com/stratisproject/prysm-stratis/beacon-chain/sync/initial-sync/testing"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/assert"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
	"github.com/stratisproject/prysm-stratis/time/slots"
)

func TestBroadcastBLSChanges(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	c := params.BeaconConfig()
	c.CapellaForkEpoch = c.BellatrixForkEpoch.Add(2)
	params.OverrideBeaconConfig(c)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	s := NewService(context.Background(),
		WithP2P(mockp2p.NewTestP2P(t)),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	var emptySig [96]byte
	s.cfg.blsToExecPool.InsertBLSToExecChange(&ethpb.SignedBLSToExecutionChange{
		Message: &ethpb.BLSToExecutionChange{
			ValidatorIndex:     10,
			FromBlsPubkey:      make([]byte, 48),
			ToExecutionAddress: make([]byte, 20),
		},
		Signature: emptySig[:],
	})

	capellaStart, err := slots.EpochStart(params.BeaconConfig().CapellaForkEpoch)
	require.NoError(t, err)
	s.broadcastBLSChanges(capellaStart + 1)
}

func TestRateBLSChanges(t *testing.T) {
	logHook := logTest.NewGlobal()
	params.SetupTestConfigCleanup(t)
	c := params.BeaconConfig()
	c.CapellaForkEpoch = c.BellatrixForkEpoch.Add(2)
	params.OverrideBeaconConfig(c)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	p1 := mockp2p.NewTestP2P(t)
	s := NewService(context.Background(),
		WithP2P(p1),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	beaconDB := testingdb.SetupDB(t)
	s.cfg.stateGen = stategen.New(beaconDB, doublylinkedtree.New())
	s.cfg.beaconDB = beaconDB
	s.initCaches()
	st, keys := util.DeterministicGenesisStateCapella(t, 256)
	s.cfg.chain = &mockChain.ChainService{
		ValidatorsRoot: [32]byte{'A'},
		Genesis:        time.Now().Add(-time.Second * time.Duration(params.BeaconConfig().SecondsPerSlot) * time.Duration(10)),
		State:          st,
	}

	for i := 0; i < 200; i++ {
		message := &ethpb.BLSToExecutionChange{
			ValidatorIndex:     primitives.ValidatorIndex(i),
			FromBlsPubkey:      keys[i+1].PublicKey().Marshal(),
			ToExecutionAddress: bytesutil.PadTo([]byte("address"), 20),
		}
		epoch := params.BeaconConfig().CapellaForkEpoch + 1
		domain, err := signing.Domain(st.Fork(), epoch, params.BeaconConfig().DomainBLSToExecutionChange, st.GenesisValidatorsRoot())
		assert.NoError(t, err)
		htr, err := signing.Data(message.HashTreeRoot, domain)
		assert.NoError(t, err)
		signed := &ethpb.SignedBLSToExecutionChange{
			Message:   message,
			Signature: keys[i+1].Sign(htr[:]).Marshal(),
		}

		s.cfg.blsToExecPool.InsertBLSToExecChange(signed)
	}

	require.Equal(t, false, p1.BroadcastCalled.Load())
	slot, err := slots.EpochStart(params.BeaconConfig().CapellaForkEpoch)
	require.NoError(t, err)
	s.broadcastBLSChanges(slot)
	time.Sleep(100 * time.Millisecond) // Need a sleep for the go routine to be ready
	require.Equal(t, true, p1.BroadcastCalled.Load())
	require.LogsDoNotContain(t, logHook, "could not")

	p1.BroadcastCalled.Store(false)
	time.Sleep(500 * time.Millisecond) // Need a sleep for the second batch to be broadcast
	require.Equal(t, true, p1.BroadcastCalled.Load())
	require.LogsDoNotContain(t, logHook, "could not")
}

func TestBroadcastBLSBatch_changes_slice(t *testing.T) {
	message := &ethpb.BLSToExecutionChange{
		FromBlsPubkey:      make([]byte, 48),
		ToExecutionAddress: make([]byte, 20),
	}
	signed := &ethpb.SignedBLSToExecutionChange{
		Message:   message,
		Signature: make([]byte, 96),
	}
	changes := make([]*ethpb.SignedBLSToExecutionChange, 200)
	for i := 0; i < len(changes); i++ {
		changes[i] = signed
	}
	p1 := mockp2p.NewTestP2P(t)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	s := NewService(context.Background(),
		WithP2P(p1),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	beaconDB := testingdb.SetupDB(t)
	s.cfg.stateGen = stategen.New(beaconDB, doublylinkedtree.New())
	s.cfg.beaconDB = beaconDB
	s.initCaches()
	st, _ := util.DeterministicGenesisStateCapella(t, 32)
	s.cfg.chain = &mockChain.ChainService{
		ValidatorsRoot: [32]byte{'A'},
		Genesis:        time.Now().Add(-time.Second * time.Duration(params.BeaconConfig().SecondsPerSlot) * time.Duration(10)),
		State:          st,
	}

	s.broadcastBLSBatch(s.ctx, &changes)
	require.Equal(t, 200-128, len(changes))
}
