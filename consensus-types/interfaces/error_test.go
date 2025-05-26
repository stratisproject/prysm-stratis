package interfaces

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestNewInvalidCastError(t *testing.T) {
	err := NewInvalidCastError(version.Phase0, version.Electra)
	require.Equal(t, true, errors.Is(err, ErrInvalidCast))
}
