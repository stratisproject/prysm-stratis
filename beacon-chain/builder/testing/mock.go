package testing

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/api/client/builder"
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	v1 "github.com/stratisproject/prysm-stratis/proto/engine/v1"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/time/slots"
)

// Config defines a config struct for dependencies into the service.
type Config struct {
	BeaconDB db.HeadAccessDatabase
}

// MockBuilderService to mock builder.
type MockBuilderService struct {
	HasConfigured         bool
	Payload               *v1.ExecutionPayload
	PayloadCapella        *v1.ExecutionPayloadCapella
	PayloadDeneb          *v1.ExecutionPayloadDeneb
	BlobBundle            *v1.BlobsBundle
	ErrSubmitBlindedBlock error
	Bid                   *ethpb.SignedBuilderBid
	BidCapella            *ethpb.SignedBuilderBidCapella
	BidDeneb              *ethpb.SignedBuilderBidDeneb
	RegistrationCache     *cache.RegistrationCache
	ErrGetHeader          error
	ErrRegisterValidator  error
	Cfg                   *Config
}

// Configured for mocking.
func (s *MockBuilderService) Configured() bool {
	return s.HasConfigured
}

// SubmitBlindedBlock for mocking.
func (s *MockBuilderService) SubmitBlindedBlock(_ context.Context, b interfaces.ReadOnlySignedBeaconBlock) (interfaces.ExecutionData, *v1.BlobsBundle, error) {
	switch b.Version() {
	case version.Bellatrix:
		w, err := blocks.WrappedExecutionPayload(s.Payload)
		if err != nil {
			return nil, nil, errors.Wrap(err, "could not wrap payload")
		}
		return w, nil, s.ErrSubmitBlindedBlock
	case version.Capella:
		w, err := blocks.WrappedExecutionPayloadCapella(s.PayloadCapella, big.NewInt(0))
		if err != nil {
			return nil, nil, errors.Wrap(err, "could not wrap capella payload")
		}
		return w, nil, s.ErrSubmitBlindedBlock
	case version.Deneb:
		w, err := blocks.WrappedExecutionPayloadDeneb(s.PayloadDeneb, big.NewInt(0))
		if err != nil {
			return nil, nil, errors.Wrap(err, "could not wrap deneb payload")
		}
		return w, s.BlobBundle, s.ErrSubmitBlindedBlock
	default:
		return nil, nil, errors.New("unknown block version for mocking")
	}
}

// GetHeader for mocking.
func (s *MockBuilderService) GetHeader(_ context.Context, slot primitives.Slot, _ [32]byte, _ [48]byte) (builder.SignedBid, error) {
	if slots.ToEpoch(slot) >= params.BeaconConfig().DenebForkEpoch || s.BidDeneb != nil {
		return builder.WrappedSignedBuilderBidDeneb(s.BidDeneb)
	}
	if slots.ToEpoch(slot) >= params.BeaconConfig().CapellaForkEpoch || s.BidCapella != nil {
		return builder.WrappedSignedBuilderBidCapella(s.BidCapella)
	}
	w, err := builder.WrappedSignedBuilderBid(s.Bid)
	if err != nil {
		return nil, errors.Wrap(err, "could not wrap capella bid")
	}
	return w, s.ErrGetHeader
}

// RegistrationByValidatorID returns either the values from the cache or db.
func (s *MockBuilderService) RegistrationByValidatorID(ctx context.Context, id primitives.ValidatorIndex) (*ethpb.ValidatorRegistrationV1, error) {
	if s.RegistrationCache != nil {
		return s.RegistrationCache.RegistrationByIndex(id)
	}
	if s.Cfg.BeaconDB != nil {
		return s.Cfg.BeaconDB.RegistrationByValidatorID(ctx, id)
	}
	return nil, cache.ErrNotFoundRegistration
}

// RegisterValidator for mocking.
func (s *MockBuilderService) RegisterValidator(context.Context, []*ethpb.SignedValidatorRegistrationV1) error {
	return s.ErrRegisterValidator
}
