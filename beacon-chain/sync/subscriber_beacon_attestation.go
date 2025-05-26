package sync

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	"github.com/stratisproject/prysm-stratis/config/features"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/container/slice"
	eth "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/time/slots"
	"google.golang.org/protobuf/proto"
)

func (s *Service) committeeIndexBeaconAttestationSubscriber(_ context.Context, msg proto.Message) error {
	a, ok := msg.(eth.Att)
	if !ok {
		return fmt.Errorf("message was not type eth.Att, type=%T", msg)
	}

	if features.Get().EnableExperimentalAttestationPool {
		return s.cfg.attestationCache.Add(a)
	} else {
		exists, err := s.cfg.attPool.HasAggregatedAttestation(a)
		if err != nil {
			return errors.Wrap(err, "could not determine if attestation pool has this attestation")
		}
		if exists {
			return nil
		}
		return s.cfg.attPool.SaveUnaggregatedAttestation(a)
	}
}

func (*Service) persistentSubnetIndices() []uint64 {
	return cache.SubnetIDs.GetAllSubnets()
}

func (*Service) aggregatorSubnetIndices(currentSlot primitives.Slot) []uint64 {
	endEpoch := slots.ToEpoch(currentSlot) + 1
	endSlot := params.BeaconConfig().SlotsPerEpoch.Mul(uint64(endEpoch))
	var commIds []uint64
	for i := currentSlot; i <= endSlot; i++ {
		commIds = append(commIds, cache.SubnetIDs.GetAggregatorSubnetIDs(i)...)
	}
	return slice.SetUint64(commIds)
}

func (*Service) attesterSubnetIndices(currentSlot primitives.Slot) []uint64 {
	endEpoch := slots.ToEpoch(currentSlot) + 1
	endSlot := params.BeaconConfig().SlotsPerEpoch.Mul(uint64(endEpoch))
	var commIds []uint64
	for i := currentSlot; i <= endSlot; i++ {
		commIds = append(commIds, cache.SubnetIDs.GetAttesterSubnetIDs(i)...)
	}
	return slice.SetUint64(commIds)
}
