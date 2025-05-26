package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	common "github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/operations"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func blockWithSyncAggregate(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	sa := &ethpb.SyncAggregate{}
	if err := sa.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockBellatrix()
	b.Block.Body = &ethpb.BeaconBlockBodyBellatrix{SyncAggregate: sa}
	return blocks.NewSignedBeaconBlock(b)
}

func RunSyncCommitteeTest(t *testing.T, config string) {
	common.RunSyncCommitteeTest(t, config, version.String(version.Bellatrix), blockWithSyncAggregate, sszToState)
}
