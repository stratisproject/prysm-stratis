package node

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/builder"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db/filesystem"
	"github.com/stratisproject/prysm-stratis/beacon-chain/execution"
	"github.com/stratisproject/prysm-stratis/consensus-types/primitives"
)

// Option for beacon node configuration.
type Option func(bn *BeaconNode) error

// WithBlockchainFlagOptions includes functional options for the blockchain service related to CLI flags.
func WithBlockchainFlagOptions(opts []blockchain.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.blockchainFlagOpts = opts
		return nil
	}
}

// WithExecutionChainOptions includes functional options for the execution chain service related to CLI flags.
func WithExecutionChainOptions(opts []execution.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.executionChainFlagOpts = opts
		return nil
	}
}

// WithBuilderFlagOptions includes functional options for the builder service related to CLI flags.
func WithBuilderFlagOptions(opts []builder.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.builderOpts = opts
		return nil
	}
}

// WithBlobStorage sets the BlobStorage backend for the BeaconNode
func WithBlobStorage(bs *filesystem.BlobStorage) Option {
	return func(bn *BeaconNode) error {
		bn.BlobStorage = bs
		return nil
	}
}

// WithBlobStorageOptions appends 1 or more filesystem.BlobStorageOption on the beacon node,
// to be used when initializing blob storage.
func WithBlobStorageOptions(opt ...filesystem.BlobStorageOption) Option {
	return func(bn *BeaconNode) error {
		bn.BlobStorageOptions = append(bn.BlobStorageOptions, opt...)
		return nil
	}
}

// WithBlobRetentionEpochs sets the blobRetentionEpochs value, used in kv store initialization.
func WithBlobRetentionEpochs(e primitives.Epoch) Option {
	return func(bn *BeaconNode) error {
		bn.blobRetentionEpochs = e
		return nil
	}
}
