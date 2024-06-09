package simulator

import (
	"testing"

	mock "github.com/stratisproject/prysm-stratis/beacon-chain/blockchain/testing"
	dbtest "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	mockstategen "github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen/mock"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/crypto/bls"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func setupService(t *testing.T, params *Parameters) *Simulator {
	slasherDB := dbtest.SetupSlasherDB(t)
	beaconState, err := util.NewBeaconState()
	require.NoError(t, err)

	// We setup validators in the beacon state along with their
	// private keys used to generate valid signatures in generated objects.
	validators := make([]*ethpb.Validator, params.NumValidators)
	privKeys := make(map[primitives.ValidatorIndex]bls.SecretKey)
	for valIdx := range validators {
		privKey, err := bls.RandKey()
		require.NoError(t, err)
		privKeys[primitives.ValidatorIndex(valIdx)] = privKey
		validators[valIdx] = &ethpb.Validator{
			PublicKey:             privKey.PublicKey().Marshal(),
			WithdrawalCredentials: make([]byte, 32),
		}
	}
	err = beaconState.SetValidators(validators)
	require.NoError(t, err)

	gen := mockstategen.NewService()
	gen.AddStateForRoot(beaconState, [32]byte{})
	return &Simulator{
		srvConfig: &ServiceConfig{
			Params:                      params,
			Database:                    slasherDB,
			AttestationStateFetcher:     &mock.ChainService{State: beaconState},
			PrivateKeysByValidatorIndex: privKeys,
			StateGen:                    gen,
		},
	}
}
