package stateutil

import (
	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/math"
)

// UnrealizedCheckpointBalances returns the total current active balance, the
// total previous epoch correctly attested for target balance, and the total
// current epoch correctly attested for target balance. It takes the current and
// previous epoch participation bits as parameters so implicitly only works for
// beacon states post-Altair.
func UnrealizedCheckpointBalances(cp, pp []byte, validators ValReader, currentEpoch primitives.Epoch) (uint64, uint64, uint64, error) {
	targetIdx := params.BeaconConfig().TimelyTargetFlagIndex
	activeBalance := uint64(0)
	currentTarget := uint64(0)
	prevTarget := uint64(0)
	if len(cp) < validators.Len() || len(pp) < validators.Len() {
		return 0, 0, 0, errors.New("participation does not match validator set")
	}

	valLength := validators.Len()
	for i := 0; i < valLength; i++ {
		v, err := validators.At(i)
		if err != nil {
			return 0, 0, 0, err
		}
		activeCurrent := v.ActivationEpoch <= currentEpoch && currentEpoch < v.ExitEpoch
		if activeCurrent {
			activeBalance, err = math.Add64(activeBalance, v.EffectiveBalance)
			if err != nil {
				return 0, 0, 0, err
			}
		}
		if v.Slashed {
			continue
		}
		if activeCurrent && ((cp[i]>>targetIdx)&1) == 1 {
			currentTarget, err = math.Add64(currentTarget, v.EffectiveBalance)
			if err != nil {
				return 0, 0, 0, err
			}
		}
		activePrevious := v.ActivationEpoch < currentEpoch && currentEpoch <= v.ExitEpoch
		if activePrevious && ((pp[i]>>targetIdx)&1) == 1 {
			prevTarget, err = math.Add64(prevTarget, v.EffectiveBalance)
			if err != nil {
				return 0, 0, 0, err
			}
		}
	}
	activeBalance, prevTarget, currentTarget = ensureLowerBound(activeBalance, prevTarget, currentTarget)
	return activeBalance, prevTarget, currentTarget, nil
}

func ensureLowerBound(activeCurrEpoch, prevTargetAttested, currTargetAttested uint64) (uint64, uint64, uint64) {
	ebi := params.BeaconConfig().EffectiveBalanceIncrement
	if ebi > activeCurrEpoch {
		activeCurrEpoch = ebi
	}
	if ebi > prevTargetAttested {
		prevTargetAttested = ebi
	}
	if ebi > currTargetAttested {
		currTargetAttested = ebi
	}
	return activeCurrEpoch, prevTargetAttested, currTargetAttested
}
