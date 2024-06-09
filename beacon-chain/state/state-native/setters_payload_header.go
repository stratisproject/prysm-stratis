package state_native

import (
	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/state-native/types"
	consensusblocks "github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	enginev1 "github.com/stratisproject/prysm-stratis/proto/engine/v1"
	_ "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
)

// SetLatestExecutionPayloadHeader for the beacon state.
func (b *BeaconState) SetLatestExecutionPayloadHeader(val interfaces.ExecutionData) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.version < version.Bellatrix {
		return errNotSupported("SetLatestExecutionPayloadHeader", b.version)
	}

	switch header := val.Proto().(type) {
	case *enginev1.ExecutionPayload:
		latest, err := consensusblocks.PayloadToHeader(val)
		if err != nil {
			return errors.Wrap(err, "could not convert payload to header")
		}
		b.latestExecutionPayloadHeader = latest
		b.markFieldAsDirty(types.LatestExecutionPayloadHeader)
		return nil
	case *enginev1.ExecutionPayloadCapella:
		latest, err := consensusblocks.PayloadToHeaderCapella(val)
		if err != nil {
			return errors.Wrap(err, "could not convert payload to header")
		}
		b.latestExecutionPayloadHeaderCapella = latest
		b.markFieldAsDirty(types.LatestExecutionPayloadHeaderCapella)
		return nil
	case *enginev1.ExecutionPayloadDeneb:
		latest, err := consensusblocks.PayloadToHeaderDeneb(val)
		if err != nil {
			return errors.Wrap(err, "could not convert payload to header")
		}
		b.latestExecutionPayloadHeaderDeneb = latest
		b.markFieldAsDirty(types.LatestExecutionPayloadHeaderDeneb)
		return nil
	case *enginev1.ExecutionPayloadHeader:
		b.latestExecutionPayloadHeader = header
		b.markFieldAsDirty(types.LatestExecutionPayloadHeader)
		return nil
	case *enginev1.ExecutionPayloadHeaderCapella:
		b.latestExecutionPayloadHeaderCapella = header
		b.markFieldAsDirty(types.LatestExecutionPayloadHeaderCapella)
		return nil
	case *enginev1.ExecutionPayloadHeaderDeneb:
		b.latestExecutionPayloadHeaderDeneb = header
		b.markFieldAsDirty(types.LatestExecutionPayloadHeaderDeneb)
		return nil
	default:
		return errors.New("value must be an execution payload header")
	}
}
