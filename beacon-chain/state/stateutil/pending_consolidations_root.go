package stateutil

import (
	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/encoding/ssz"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func PendingConsolidationsRoot(slice []*ethpb.PendingConsolidation) ([32]byte, error) {
	return ssz.SliceRoot(slice, fieldparams.PendingConsolidationsLimit)
}
