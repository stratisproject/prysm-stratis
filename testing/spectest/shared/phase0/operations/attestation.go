package operations

import (
	"testing"

	b "github.com/stratisproject/prysm-stratis/beacon-chain/core/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	common "github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/operations"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func blockWithAttestation(attestationSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	att := &ethpb.Attestation{}
	if err := att.UnmarshalSSZ(attestationSSZ); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlock()
	b.Block.Body = &ethpb.BeaconBlockBody{Attestations: []*ethpb.Attestation{att}}
	return blocks.NewSignedBeaconBlock(b)
}

// RunAttestationTest executes "operations/attestation" tests.
func RunAttestationTest(t *testing.T, config string) {
	common.RunAttestationTest(t, config, version.String(version.Phase0), blockWithAttestation, b.ProcessAttestationsNoVerifySignature, sszToState)
}
