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

func blockWithProposerSlashing(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	ps := &ethpb.ProposerSlashing{}
	if err := ps.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockCapella()
	b.Block.Body = &ethpb.BeaconBlockBodyCapella{ProposerSlashings: []*ethpb.ProposerSlashing{ps}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunProposerSlashingTest(t *testing.T, config string) {
	common.RunProposerSlashingTest(t, config, version.String(version.Capella), blockWithProposerSlashing, sszToState)
}
