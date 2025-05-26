package types

import (
	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	consensus_blocks "github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// Checkpoint is an array version of ethpb.Checkpoint. It is used internally in
// forkchoice, while the slice version is used in the interface to legacy code
// in other packages
type Checkpoint struct {
	Epoch primitives.Epoch
	Root  [fieldparams.RootLength]byte
}

// BlockAndCheckpoints to call the InsertOptimisticChain function
type BlockAndCheckpoints struct {
	Block               consensus_blocks.ROBlock
	JustifiedCheckpoint *ethpb.Checkpoint
	FinalizedCheckpoint *ethpb.Checkpoint
}
