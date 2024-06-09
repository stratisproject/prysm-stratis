package monitor

import (
	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// processExitsFromBlock logs the event when a tracked validators' exit was included in a block
func (s *Service) processExitsFromBlock(blk interfaces.ReadOnlyBeaconBlock) {
	s.RLock()
	defer s.RUnlock()
	for _, exit := range blk.Body().VoluntaryExits() {
		idx := exit.Exit.ValidatorIndex
		if s.trackedIndex(idx) {
			log.WithFields(logrus.Fields{
				"ValidatorIndex": idx,
				"Slot":           blk.Slot(),
			}).Info("Voluntary exit was included")
		}
	}
}

// processExit logs the event when tracked validators' exit was processed
func (s *Service) processExit(exit *ethpb.SignedVoluntaryExit) {
	idx := exit.Exit.ValidatorIndex
	s.RLock()
	defer s.RUnlock()
	if s.trackedIndex(idx) {
		log.WithFields(logrus.Fields{
			"ValidatorIndex": idx,
		}).Info("Voluntary exit was processed")
	}
}
