package testing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/testing/assert"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

type getState func() (state.BeaconState, error)

func VerifyBeaconStateValidatorAtIndexReadOnlyHandlesNilSlice(t *testing.T, factory getState) {
	st, err := factory()
	require.NoError(t, err)

	_, err = st.ValidatorAtIndexReadOnly(0)
	assert.Equal(t, state.ErrNilValidatorsInState, err)
}
