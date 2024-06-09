// Package stateutils contains useful tools for faster computation
// of state transitions using maps to represent validators instead
// of slices.
package stateutils

import (
	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// ValidatorIndexMap builds a lookup map for quickly determining the index of
// a validator by their public key.
func ValidatorIndexMap(validators []*ethpb.Validator) map[[fieldparams.BLSPubkeyLength]byte]primitives.ValidatorIndex {
	m := make(map[[fieldparams.BLSPubkeyLength]byte]primitives.ValidatorIndex, len(validators))
	if validators == nil {
		return m
	}
	for idx, record := range validators {
		if record == nil {
			continue
		}
		key := bytesutil.ToBytes48(record.PublicKey)
		m[key] = primitives.ValidatorIndex(idx)
	}
	return m
}
