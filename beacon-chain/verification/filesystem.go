package verification

import (
	"github.com/spf13/afero"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func VerifiedROBlobFromDisk(fs afero.Fs, root [32]byte, path string) (blocks.VerifiedROBlob, error) {
	encoded, err := afero.ReadFile(fs, path)
	if err != nil {
		return VerifiedROBlobError(err)
	}
	s := &ethpb.BlobSidecar{}
	if err := s.UnmarshalSSZ(encoded); err != nil {
		return VerifiedROBlobError(err)
	}
	ro, err := blocks.NewROBlobWithRoot(s, root)
	if err != nil {
		return VerifiedROBlobError(err)
	}
	return blocks.NewVerifiedROBlob(ro), nil
}
