package validator_client_factory

import (
	"github.com/stratisproject/prysm-stratis/config/features"
	beaconApi "github.com/stratisproject/prysm-stratis/validator/client/beacon-api"
	grpcApi "github.com/stratisproject/prysm-stratis/validator/client/grpc-api"
	"github.com/stratisproject/prysm-stratis/validator/client/iface"
	nodeClientFactory "github.com/stratisproject/prysm-stratis/validator/client/node-client-factory"
	validatorHelpers "github.com/stratisproject/prysm-stratis/validator/helpers"
)

func NewBeaconChainClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.BeaconChainClient {
	grpcClient := grpcApi.NewGrpcBeaconChainClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiBeaconChainClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}

func NewPrysmBeaconClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.PrysmBeaconChainClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewPrysmBeaconChainClient(jsonRestHandler, nodeClientFactory.NewNodeClient(validatorConn, jsonRestHandler))
	} else {
		return grpcApi.NewGrpcPrysmBeaconChainClient(validatorConn.GetGrpcClientConn())
	}
}
