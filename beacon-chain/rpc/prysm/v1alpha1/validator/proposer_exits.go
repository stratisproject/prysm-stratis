package validator

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func (vs *Server) getExits(head state.BeaconState, slot primitives.Slot) []*ethpb.SignedVoluntaryExit {
	exits, err := vs.ExitPool.ExitsForInclusion(head, slot)
	if err != nil {
		log.WithError(err).Error("Could not get exits")
		return []*ethpb.SignedVoluntaryExit{}
	}
	return exits
}
