package logging

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
)

// BlobFields extracts a standard set of fields from a BlobSidecar into a logrus.Fields struct
// which can be passed to log.WithFields.
func BlobFields(blob blocks.ROBlob) logrus.Fields {
	return logrus.Fields{
		"slot":          blob.Slot(),
		"proposerIndex": blob.ProposerIndex(),
		"blockRoot":     fmt.Sprintf("%#x", blob.BlockRoot()),
		"parentRoot":    fmt.Sprintf("%#x", blob.ParentRoot()),
		"kzgCommitment": fmt.Sprintf("%#x", blob.KzgCommitment),
		"index":         blob.Index,
	}
}

// BlockFieldsFromBlob extracts the set of fields from a given BlobSidecar which are shared by the block and
// all other sidecars for the block.
func BlockFieldsFromBlob(blob blocks.ROBlob) logrus.Fields {
	return logrus.Fields{
		"slot":          blob.Slot(),
		"proposerIndex": blob.ProposerIndex(),
		"blockRoot":     fmt.Sprintf("%#x", blob.BlockRoot()),
		"parentRoot":    fmt.Sprintf("%#x", blob.ParentRoot()),
	}
}
