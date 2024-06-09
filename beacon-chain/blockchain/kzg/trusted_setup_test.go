package kzg

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
