package blocks_test

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/beacon-chain/core/blocks"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	"github.com/stratisproject/prysm-stratis/testing/assert"
)

func TestGenesisBlock_InitializedCorrectly(t *testing.T) {
	stateHash := bytesutil.PadTo([]byte{0}, 32)
	b1 := blocks.NewGenesisBlock(stateHash)

	assert.NotNil(t, b1.Block.ParentRoot, "Genesis block missing ParentHash field")
	assert.DeepEqual(t, b1.Block.StateRoot, stateHash, "Genesis block StateRootHash32 isn't initialized correctly")
}
