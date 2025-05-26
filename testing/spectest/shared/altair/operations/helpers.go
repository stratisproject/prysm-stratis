package operations

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	state_native "github.com/stratisproject/prysm-stratis/beacon-chain/state/state-native"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func sszToState(b []byte) (state.BeaconState, error) {
	base := &ethpb.BeaconStateAltair{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return state_native.InitializeFromProtoAltair(base)
}

func sszToBlock(b []byte) (interfaces.SignedBeaconBlock, error) {
	base := &ethpb.BeaconBlockAltair{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: base})
}
