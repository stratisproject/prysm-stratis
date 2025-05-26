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

func blockWithVoluntaryExit(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	e := &ethpb.SignedVoluntaryExit{}
	if err := e.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockDeneb()
	b.Block.Body = &ethpb.BeaconBlockBodyDeneb{VoluntaryExits: []*ethpb.SignedVoluntaryExit{e}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunVoluntaryExitTest(t *testing.T, config string) {
	common.RunVoluntaryExitTest(t, config, version.String(version.Deneb), blockWithVoluntaryExit, sszToState)
}
