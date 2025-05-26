package iface

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stratisproject/prysm-stratis/consensus-types/validator"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

var ErrNotSupported = errors.New("endpoint not supported")

type ValidatorCount struct {
	Status string
	Count  uint64
}

// PrysmChainClient defines an interface required to implement all the prysm specific custom endpoints.
type PrysmChainClient interface {
	ValidatorCount(context.Context, string, []validator.Status) ([]ValidatorCount, error)
	ValidatorPerformance(context.Context, *ethpb.ValidatorPerformanceRequest) (*ethpb.ValidatorPerformanceResponse, error)
}
