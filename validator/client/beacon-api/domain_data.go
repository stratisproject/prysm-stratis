package beacon_api

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/signing"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/network/forks"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func (c beaconApiValidatorClient) getDomainData(ctx context.Context, epoch primitives.Epoch, domainType [4]byte) (*ethpb.DomainResponse, error) {
	// Get the fork version from the given epoch
	fork, err := forks.Fork(epoch)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get fork version for epoch %d", epoch)
	}

	// Get the genesis validator root
	genesis, err := c.genesisProvider.GetGenesis(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get genesis info")
	}

	if !validRoot(genesis.GenesisValidatorsRoot) {
		return nil, errors.Errorf("invalid genesis validators root: %s", genesis.GenesisValidatorsRoot)
	}

	genesisValidatorRoot, err := hexutil.Decode(genesis.GenesisValidatorsRoot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode genesis validators root")
	}

	signatureDomain, err := signing.Domain(fork, epoch, domainType, genesisValidatorRoot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compute signature domain")
	}

	return &ethpb.DomainResponse{SignatureDomain: signatureDomain}, nil
}
