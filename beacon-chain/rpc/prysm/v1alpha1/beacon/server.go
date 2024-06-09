// Package beacon defines a gRPC beacon service implementation, providing
// useful endpoints for checking fetching chain-specific data such as
// blocks, committees, validators, assignments, and more.
package beacon

import (
	"context"
	"time"

	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	blockfeed "github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/block"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/operation"
	statefeed "github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/state"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/execution"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/attestations"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/slashings"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/core"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum beacon chain.
type Server struct {
	BeaconDB                    db.ReadOnlyDatabase
	Ctx                         context.Context
	ChainStartFetcher           execution.ChainStartFetcher
	HeadFetcher                 blockchain.HeadFetcher
	CanonicalFetcher            blockchain.CanonicalFetcher
	FinalizationFetcher         blockchain.FinalizationFetcher
	DepositFetcher              cache.DepositFetcher
	BlockFetcher                execution.POWBlockFetcher
	GenesisTimeFetcher          blockchain.TimeFetcher
	StateNotifier               statefeed.Notifier
	BlockNotifier               blockfeed.Notifier
	AttestationNotifier         operation.Notifier
	Broadcaster                 p2p.Broadcaster
	AttestationsPool            attestations.Pool
	SlashingsPool               slashings.PoolManager
	ChainStartChan              chan time.Time
	ReceivedAttestationsBuffer  chan *ethpb.Attestation
	CollectedAttestationsBuffer chan []*ethpb.Attestation
	StateGen                    stategen.StateManager
	SyncChecker                 sync.Checker
	ReplayerBuilder             stategen.ReplayerBuilder
	OptimisticModeFetcher       blockchain.OptimisticModeFetcher
	CoreService                 *core.Service
}
