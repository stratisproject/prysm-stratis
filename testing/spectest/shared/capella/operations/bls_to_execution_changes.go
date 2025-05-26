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

func blockWithBlsChange(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	c := &ethpb.SignedBLSToExecutionChange{}
	if err := c.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockCapella()
	b.Block.Body = &ethpb.BeaconBlockBodyCapella{BlsToExecutionChanges: []*ethpb.SignedBLSToExecutionChange{c}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunBLSToExecutionChangeTest(t *testing.T, config string) {
	common.RunBLSToExecutionChangeTest(t, config, version.String(version.Capella), blockWithBlsChange, sszToState)
}
