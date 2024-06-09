package blockchain

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	testDB "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	doublylinkedtree "github.com/stratisproject/prysm-stratis/beacon-chain/forkchoice/doubly-linked-tree"
	forkchoicetypes "github.com/stratisproject/prysm-stratis/beacon-chain/forkchoice/types"
	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
	"github.com/stratisproject/prysm-stratis/time/slots"
)

func TestService_VerifyWeakSubjectivityRoot(t *testing.T) {
	beaconDB := testDB.SetupDB(t)

	b := util.NewBeaconBlock()
	b.Block.Slot = 1792480
	util.SaveBlock(t, context.Background(), beaconDB, b)
	r, err := b.Block.HashTreeRoot()
	require.NoError(t, err)

	blockEpoch := slots.ToEpoch(b.Block.Slot)
	tests := []struct {
		wsVerified     bool
		disabled       bool
		wantErr        error
		checkpt        *ethpb.Checkpoint
		finalizedEpoch primitives.Epoch
		name           string
	}{
		{
			name:     "nil root and epoch",
			disabled: true,
		},
		{
			name:           "not yet to verify, ws epoch higher than finalized epoch",
			checkpt:        &ethpb.Checkpoint{Root: bytesutil.PadTo([]byte{'a'}, 32), Epoch: blockEpoch},
			finalizedEpoch: blockEpoch - 1,
		},
		{
			name:           "can't find the block in DB",
			checkpt:        &ethpb.Checkpoint{Root: bytesutil.PadTo([]byte{'a'}, fieldparams.RootLength), Epoch: 1},
			finalizedEpoch: blockEpoch + 1,
			wantErr:        errWSBlockNotFound,
		},
		{
			name:           "can't find the block corresponds to ws epoch in DB",
			checkpt:        &ethpb.Checkpoint{Root: r[:], Epoch: blockEpoch - 2}, // Root belongs in epoch 1.
			finalizedEpoch: blockEpoch - 1,
			wantErr:        errWSBlockNotFoundInEpoch,
		},
		{
			name:           "can verify and pass",
			checkpt:        &ethpb.Checkpoint{Root: r[:], Epoch: blockEpoch},
			finalizedEpoch: blockEpoch + 1,
		},
		{
			name:           "equal epoch",
			checkpt:        &ethpb.Checkpoint{Root: r[:], Epoch: blockEpoch},
			finalizedEpoch: blockEpoch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wv, err := NewWeakSubjectivityVerifier(tt.checkpt, beaconDB)
			require.Equal(t, !tt.disabled, wv.enabled)
			require.NoError(t, err)
			fcs := doublylinkedtree.New()
			s := &Service{
				cfg:        &config{BeaconDB: beaconDB, WeakSubjectivityCheckpt: tt.checkpt, ForkChoiceStore: fcs},
				wsVerifier: wv,
			}
			require.NoError(t, fcs.UpdateFinalizedCheckpoint(&forkchoicetypes.Checkpoint{Epoch: tt.finalizedEpoch}))
			cp := s.cfg.ForkChoiceStore.FinalizedCheckpoint()
			err = s.wsVerifier.VerifyWeakSubjectivity(context.Background(), cp.Epoch)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, true, errors.Is(err, tt.wantErr))
			}
		})
	}
}
