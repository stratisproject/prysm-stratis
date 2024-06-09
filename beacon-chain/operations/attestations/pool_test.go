package attestations

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
