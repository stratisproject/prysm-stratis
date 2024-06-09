package validator

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/builder"
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/operation"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/attestations"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/synccommittee"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/core"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/eth/rewards"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
	eth "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Validator service,
// providing RPC endpoints intended for validator clients.
type Server struct {
	HeadFetcher            blockchain.HeadFetcher
	TimeFetcher            blockchain.TimeFetcher
	SyncChecker            sync.Checker
	AttestationsPool       attestations.Pool
	PeerManager            p2p.PeerManager
	Broadcaster            p2p.Broadcaster
	Stater                 lookup.Stater
	OptimisticModeFetcher  blockchain.OptimisticModeFetcher
	SyncCommitteePool      synccommittee.Pool
	V1Alpha1Server         eth.BeaconNodeValidatorServer
	ChainInfoFetcher       blockchain.ChainInfoFetcher
	BeaconDB               db.HeadAccessDatabase
	BlockBuilder           builder.BlockBuilder
	OperationNotifier      operation.Notifier
	CoreService            *core.Service
	BlockRewardFetcher     rewards.BlockRewardsFetcher
	TrackedValidatorsCache *cache.TrackedValidatorsCache
	PayloadIDCache         *cache.PayloadIDCache
}
