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

func blockWithWithdrawalRequest(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	wr := &enginev1.WithdrawalRequest{}
	if err := wr.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	er := &enginev1.ExecutionRequests{
		Withdrawals: []*enginev1.WithdrawalRequest{wr},
	}
	b := util.NewBeaconBlockElectra()
	b.Block.Body = &ethpb.BeaconBlockBodyElectra{ExecutionRequests: er}
	return blocks.NewSignedBeaconBlock(b)
}

func RunWithdrawalRequestTest(t *testing.T, config string) {
	common.RunWithdrawalRequestTest(t, config, version.String(version.Electra), blockWithWithdrawalRequest, sszToState)
}
