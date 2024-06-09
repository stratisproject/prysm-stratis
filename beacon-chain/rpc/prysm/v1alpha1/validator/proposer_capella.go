package validator

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
)

// Sets the bls to exec data for a block.
func (vs *Server) setBlsToExecData(blk interfaces.SignedBeaconBlock, headState state.BeaconState) {
	if blk.Version() < version.Capella {
		return
	}
	if err := blk.SetBLSToExecutionChanges([]*ethpb.SignedBLSToExecutionChange{}); err != nil {
		log.WithError(err).Error("Could not set bls to execution data in block")
		return
	}
	changes, err := vs.BLSChangesPool.BLSToExecChangesForInclusion(headState)
	if err != nil {
		log.WithError(err).Error("Could not get bls to execution changes")
		return
	} else {
		if err := blk.SetBLSToExecutionChanges(changes); err != nil {
			log.WithError(err).Error("Could not set bls to execution changes")
			return
		}
	}
}
