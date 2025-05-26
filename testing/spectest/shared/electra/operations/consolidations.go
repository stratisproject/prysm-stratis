package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	enginev1 "github.com/stratisproject/prysm-stratis/proto/engine/v1"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	common "github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/operations"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func blockWithConsolidation(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	cr := &enginev1.ConsolidationRequest{}
	if err := cr.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	er := &enginev1.ExecutionRequests{
		Consolidations: []*enginev1.ConsolidationRequest{cr},
	}
	b := util.NewBeaconBlockElectra()
	b.Block.Body = &ethpb.BeaconBlockBodyElectra{ExecutionRequests: er}
	return blocks.NewSignedBeaconBlock(b)
}

func RunConsolidationTest(t *testing.T, config string) {
	common.RunConsolidationTest(t, config, version.String(version.Electra), blockWithConsolidation, sszToState)
}
