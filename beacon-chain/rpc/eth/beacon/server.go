// Package beacon defines a gRPC beacon service implementation,
// following the official API standards https://ethereum.github.io/beacon-apis/#/.
// This package includes the beacon and config endpoints.
package beacon

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	blockfeed "github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/block"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/operation"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/execution"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/attestations"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/blstoexec"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/slashings"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/voluntaryexits"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/core"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
	eth "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	BeaconDB                      db.ReadOnlyDatabase
	ChainInfoFetcher              blockchain.ChainInfoFetcher
	GenesisTimeFetcher            blockchain.TimeFetcher
	BlockReceiver                 blockchain.BlockReceiver
	BlockNotifier                 blockfeed.Notifier
	OperationNotifier             operation.Notifier
	Broadcaster                   p2p.Broadcaster
	AttestationsPool              attestations.Pool
	SlashingsPool                 slashings.PoolManager
	VoluntaryExitsPool            voluntaryexits.PoolManager
	StateGenService               stategen.StateManager
	Stater                        lookup.Stater
	Blocker                       lookup.Blocker
	HeadFetcher                   blockchain.HeadFetcher
	TimeFetcher                   blockchain.TimeFetcher
	OptimisticModeFetcher         blockchain.OptimisticModeFetcher
	V1Alpha1ValidatorServer       eth.BeaconNodeValidatorServer
	SyncChecker                   sync.Checker
	CanonicalHistory              *stategen.CanonicalHistory
	ExecutionPayloadReconstructor execution.PayloadReconstructor
	FinalizationFetcher           blockchain.FinalizationFetcher
	BLSChangesPool                blstoexec.PoolManager
	ForkchoiceFetcher             blockchain.ForkchoiceFetcher
	CoreService                   *core.Service
}
