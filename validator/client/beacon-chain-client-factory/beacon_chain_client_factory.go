package validator_client_factory

import (
	"github.com/stratisproject/prysm-stratis/config/features"
	beaconApi "github.com/stratisproject/prysm-stratis/validator/client/beacon-api"
	grpcApi "github.com/stratisproject/prysm-stratis/validator/client/grpc-api"
	"github.com/stratisproject/prysm-stratis/validator/client/iface"
	nodeClientFactory "github.com/stratisproject/prysm-stratis/validator/client/node-client-factory"
	validatorHelpers "github.com/stratisproject/prysm-stratis/validator/helpers"
)

func NewChainClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.ChainClient {
	grpcClient := grpcApi.NewGrpcChainClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiChainClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}

func NewPrysmChainClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.PrysmChainClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewPrysmChainClient(jsonRestHandler, nodeClientFactory.NewNodeClient(validatorConn, jsonRestHandler))
	} else {
		return grpcApi.NewGrpcPrysmChainClient(validatorConn.GetGrpcClientConn())
	}
}
