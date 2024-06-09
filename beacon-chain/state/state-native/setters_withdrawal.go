package state_native

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/state-native/types"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/runtime/version"
)

// SetNextWithdrawalIndex sets the index that will be assigned to the next withdrawal.
func (b *BeaconState) SetNextWithdrawalIndex(i uint64) error {
	if b.version < version.Capella {
		return errNotSupported("SetNextWithdrawalIndex", b.version)
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	b.nextWithdrawalIndex = i
	b.markFieldAsDirty(types.NextWithdrawalIndex)
	return nil
}

// SetNextWithdrawalValidatorIndex sets the index of the validator which is
// next in line for a partial withdrawal.
func (b *BeaconState) SetNextWithdrawalValidatorIndex(i primitives.ValidatorIndex) error {
	if b.version < version.Capella {
		return errNotSupported("SetNextWithdrawalValidatorIndex", b.version)
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	b.nextWithdrawalValidatorIndex = i
	b.markFieldAsDirty(types.NextWithdrawalValidatorIndex)
	return nil
}
