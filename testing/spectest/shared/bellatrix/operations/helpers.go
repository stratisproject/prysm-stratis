package operations

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	state_native "github.com/stratisproject/prysm-stratis/beacon-chain/state/state-native"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func sszToState(b []byte) (state.BeaconState, error) {
	base := &ethpb.BeaconStateBellatrix{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return state_native.InitializeFromProtoBellatrix(base)
}

func sszToBlock(b []byte) (interfaces.SignedBeaconBlock, error) {
	base := &ethpb.BeaconBlockBellatrix{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockBellatrix{Block: base})
}

func sszToBlockBody(b []byte) (interfaces.ReadOnlyBeaconBlockBody, error) {
	base := &ethpb.BeaconBlockBodyBellatrix{}
	if err := base.UnmarshalSSZ(b); err != nil {
		return nil, err
	}
	return blocks.NewBeaconBlockBody(base)
}
