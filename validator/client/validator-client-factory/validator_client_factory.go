package validator_client_factory

import (
	"github.com/stratisproject/prysm-stratis/config/features"
	beaconApi "github.com/stratisproject/prysm-stratis/validator/client/beacon-api"
	grpcApi "github.com/stratisproject/prysm-stratis/validator/client/grpc-api"
	"github.com/stratisproject/prysm-stratis/validator/client/iface"
	validatorHelpers "github.com/stratisproject/prysm-stratis/validator/helpers"
)

func NewValidatorClient(
	validatorConn validatorHelpers.NodeConnection,
	jsonRestHandler beaconApi.JsonRestHandler,
	opt ...beaconApi.ValidatorClientOpt,
) iface.ValidatorClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiValidatorClient(jsonRestHandler, opt...)
	} else {
		return grpcApi.NewGrpcValidatorClient(validatorConn.GetGrpcClientConn())
	}
}
