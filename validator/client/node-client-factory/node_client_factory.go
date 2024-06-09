package validator_client_factory

import (
	"github.com/stratisproject/prysm-stratis/config/features"
	beaconApi "github.com/stratisproject/prysm-stratis/validator/client/beacon-api"
	grpcApi "github.com/stratisproject/prysm-stratis/validator/client/grpc-api"
	"github.com/stratisproject/prysm-stratis/validator/client/iface"
	validatorHelpers "github.com/stratisproject/prysm-stratis/validator/helpers"
)

func NewNodeClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.NodeClient {
	grpcClient := grpcApi.NewNodeClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewNodeClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}
