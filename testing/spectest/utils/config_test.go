package utils

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestConfig(t *testing.T) {
	require.NoError(t, SetConfig(t, "minimal"))
	require.Equal(t, primitives.Slot(8), params.BeaconConfig().SlotsPerEpoch)
	require.NoError(t, SetConfig(t, "mainnet"))
	require.Equal(t, primitives.Slot(32), params.BeaconConfig().SlotsPerEpoch)
}
