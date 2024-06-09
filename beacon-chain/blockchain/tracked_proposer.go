package blockchain

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/helpers"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/config/features"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
)

// trackedProposer returns whether the beacon node was informed, via the
// validators/prepare_proposer endpoint, of the proposer at the given slot.
// It only returns true if the tracked proposer is present and active.
func (s *Service) trackedProposer(st state.ReadOnlyBeaconState, slot primitives.Slot) (cache.TrackedValidator, bool) {
	if features.Get().PrepareAllPayloads {
		return cache.TrackedValidator{Active: true}, true
	}
	id, err := helpers.BeaconProposerIndexAtSlot(s.ctx, st, slot)
	if err != nil {
		return cache.TrackedValidator{}, false
	}
	val, ok := s.cfg.TrackedValidatorsCache.Validator(id)
	if !ok {
		return cache.TrackedValidator{}, false
	}
	return val, val.Active
}
