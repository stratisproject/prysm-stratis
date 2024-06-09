// Package debug defines a gRPC beacon service implementation,
// following the official API standards https://ethereum.github.io/beacon-apis/#/.
// This package includes the beacon and config endpoints.
package debug

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	BeaconDB              db.ReadOnlyDatabase
	HeadFetcher           blockchain.HeadFetcher
	Stater                lookup.Stater
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	ForkFetcher           blockchain.ForkFetcher
	ForkchoiceFetcher     blockchain.ForkchoiceFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
	ChainInfoFetcher      blockchain.ChainInfoFetcher
}
