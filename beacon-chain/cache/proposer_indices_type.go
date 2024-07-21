package cache

import (
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
)

// ProposerIndices defines the cached struct for proposer indices.
type ProposerIndices struct {
	BlockRoot       [32]byte
	ProposerIndices []primitives.ValidatorIndex
}
