package params

import (
	"math"
	"time"

	fieldparams "github.com/prysmaticlabs/prysm/v5/config/fieldparams"
	"github.com/prysmaticlabs/prysm/v5/encoding/bytesutil"
)

// MainnetConfig returns the configuration to be used in the main network.
func MainnetConfig() *BeaconChainConfig {
	if mainnetBeaconConfig.ForkVersionSchedule == nil {
		mainnetBeaconConfig.InitializeForkSchedule()
	}
	return mainnetBeaconConfig
}

const (
	// Genesis Fork Epoch for the mainnet config.
	genesisForkEpoch      = 0
	mainnetDenebForkEpoch = 0
)

var mainnetNetworkConfig = &NetworkConfig{
	ETH2Key:                    "eth2",
	AttSubnetKey:               "attnets",
	SyncCommsSubnetKey:         "syncnets",
	MinimumPeersInSubnetSearch: 20,
	ContractDeploymentBlock:    0,
	BootstrapNodes: []string{
		"enr:-LG4QAhVE9DDGdVhgIBGkmEUn5UDIY-LxDNDZS770PDbyPQ3FVaoFG-MB5WCo4YClnUeKuXKQ_c3rFSqPdT3d-w44KWGAY44QOZ5h2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhJ1apOiJc2VjcDI1NmsxoQLnrfDF1KNaF8ETmtvnhMegUmRw9to2-g5Ep_sOakomVoN1ZHCCMsg",
		"enr:-LG4QDHOmijvtuQNU_vIjKTCfZGOkUcV9ffIPRDxxLlQf1lkOmBWwH6Wnk9eWkY6Jtm8huX01t-prOxqlmTpz5k5RPyGAY44RXp9h2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhIrJdNSJc2VjcDI1NmsxoQNeFWZAJl3caL2-VsrUq24d7bm-Xpym_0KIa155NW99N4N1ZHCCMsg",
		"enr:-LG4QKtaVfME32H-s6afk6BafSHxnWj9mmtyu1JTX1j-OTAAIMNOQ7NYBj0TBlr_-vEfskUuEE5V4Oh2F0FSKsLZ7eKGAY44SCLWh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhCUbHACJc2VjcDI1NmsxoQIPCLUCODPVF1cNLbyGtuhotZI8NIRmMQrNjthjdun3x4N1ZHCCMsg",
		"enr:-LG4QK1Q24htbhQlDcJnYoL_D717t1oxecvkbMcFliG7OzGxP5GC1eK5wpT7jXmYfqMKgJiKxfjZ4XZGScCAyetsgb2GAY44So0-h2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhB_cZsuJc2VjcDI1NmsxoQOnt_WuU-aFMXyeiaHDSu-yCWGPs1gurFEUoinUv3ObKYN1ZHCCMsg",
		"enr:-LG4QLTWEf1um4Uk7S3K8YQWtVhNwJOlvSPd-vTomwnc00jjegFAbi49GKAbzB-jrEi8W-5LclNmvGQJZRobyJoYaVSGAY44TUr5h2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhF5IeR-Jc2VjcDI1NmsxoQPngFsJKjN88KSYGyezJoqJGarMBk2R0LnMI-MkNHw8dIN1ZHCCMsg",
		"enr:-LG4QBfIaPSrSrVlodcTYTPfU76E6RIur-wwFjDmlpmr_EiLZ3L2gVJQaEOOud4dzFiQwJyX7XbiJzdYbZtdse_YEkyGAY44UfFkh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhG176jeJc2VjcDI1NmsxoQKpr_p-y12D1l69-lvwfmRQ_UrAcDux_-09xUlaamCERIN1ZHCCMsg",
		"enr:-LG4QPdxscerSfZsyMpYupaURxMYSOwiFPYAt-rQsEnm0cu-H9kqFHHFN6dlnVWX094fMYyb2ABPYs1CSVfEYEsktJ-GAY44VG6Dh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhFT3ngKJc2VjcDI1NmsxoQNK1iO25wPYMyBGsjSN-NIDepwMv1IsTolwc0amAfZSuYN1ZHCCMsg",
		"enr:-LG4QLOFijNWfoReZB0FnFPIgxIWww_2JQ6YPutO-8ZwoI0FQUhKlthcZ_RH1subwqUVMOGujBHCrEtXHc3ccoXPnJyGAY44VzSnh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAc-AxSCgAAAP__________gmlkgnY0gmlwhC7686GJc2VjcDI1NmsxoQMjNS3qva1dd9WFninKHufheG50ZyphlOhTFlr7DbArX4N1ZHCCMsg",
	},
}

var mainnetBeaconConfig = &BeaconChainConfig{
	// Constants (Non-configurable)
	FarFutureEpoch:           math.MaxUint64,
	FarFutureSlot:            math.MaxUint64,
	BaseRewardsPerEpoch:      4,
	DepositContractTreeDepth: 32,
	GenesisDelay:             1,

	// Misc constant.
	TargetCommitteeSize:            128,
	MaxValidatorsPerCommittee:      2048,
	MaxCommitteesPerSlot:           64,
	MinPerEpochChurnLimit:          4,
	ChurnLimitQuotient:             1 << 16,
	ShuffleRoundCount:              90,
	MinGenesisActiveValidatorCount: 64,
	MinGenesisTime:                 1606824000, // Dec 1, 2020, 12pm UTC.
	TargetAggregatorsPerCommittee:  16,
	HysteresisQuotient:             4,
	HysteresisDownwardMultiplier:   1,
	HysteresisUpwardMultiplier:     5,

	// Gwei value constants.
	MinDepositAmount:          1 * 1e9,
	MaxEffectiveBalance:       20000 * 1e9,
	EjectionBalance:           10000 * 1e9,
	EffectiveBalanceIncrement: 1 * 1e9,
	ProposerBlockReward:       60 * 1e9,

	// Initial value constants.
	BLSWithdrawalPrefixByte:         byte(0),
	ETH1AddressWithdrawalPrefixByte: byte(1),
	ZeroHash:                        [32]byte{},

	// Time parameter constants.
	MinAttestationInclusionDelay:     1,
	SecondsPerSlot:                   15,
	SlotsPerEpoch:                    32,
	SqrRootSlotsPerEpoch:             5,
	MinSeedLookahead:                 1,
	MaxSeedLookahead:                 4,
	EpochsPerEth1VotingPeriod:        64,
	SlotsPerHistoricalRoot:           8192,
	MinValidatorWithdrawabilityDelay: 256,
	ShardCommitteePeriod:             256,
	MinEpochsToInactivityPenalty:     4,
	Eth1FollowDistance:               2048,

	// Fork choice algorithm constants.
	ProposerScoreBoost:              40,
	ReorgWeightThreshold:            20,
	ReorgParentWeightThreshold:      160,
	ReorgMaxEpochsSinceFinalization: 2,
	IntervalsPerSlot:                3,

	// Ethereum PoW parameters.
	DepositChainID:         105105, // Chain ID of eth1 mainnet.
	DepositNetworkID:       105105, // Network ID of eth1 mainnet.
	DepositContractAddress: "0x0000000000000000000000000000000000001001",

	// Validator params.
	RandomSubnetsPerValidator:         1 << 0,
	EpochsPerRandomSubnetSubscription: 1 << 8,

	// While eth1 mainnet block times are closer to 13s, we must conform with other clients in
	// order to vote on the correct eth1 blocks.
	//
	// Additional context: https://github.com/ethereum/consensus-specs/issues/2132
	// Bug prompting this change: https://github.com/prysmaticlabs/prysm/issues/7856
	// Future optimization: https://github.com/prysmaticlabs/prysm/issues/7739
	SecondsPerETH1Block: 14,

	// State list length constants.
	EpochsPerHistoricalVector: 65536,
	EpochsPerSlashingsVector:  8192,
	HistoricalRootsLimit:      16777216,
	ValidatorRegistryLimit:    1099511627776,

	// Reward and penalty quotients constants.
	BaseRewardFactor:               64,
	WhistleBlowerRewardQuotient:    512,
	ProposerRewardQuotient:         8,
	InactivityPenaltyQuotient:      67108864,
	MinSlashingPenaltyQuotient:     128,
	ProportionalSlashingMultiplier: 1,

	// Max operations per block constants.
	MaxProposerSlashings:             16,
	MaxAttesterSlashings:             2,
	MaxAttestations:                  128,
	MaxDeposits:                      16,
	MaxVoluntaryExits:                16,
	MaxWithdrawalsPerPayload:         16,
	MaxBlsToExecutionChanges:         16,
	MaxValidatorsPerWithdrawalsSweep: 16384,

	// BLS domain values.
	DomainBeaconProposer:              bytesutil.Uint32ToBytes4(0x00000000),
	DomainBeaconAttester:              bytesutil.Uint32ToBytes4(0x01000000),
	DomainRandao:                      bytesutil.Uint32ToBytes4(0x02000000),
	DomainDeposit:                     bytesutil.Uint32ToBytes4(0x03000000),
	DomainVoluntaryExit:               bytesutil.Uint32ToBytes4(0x04000000),
	DomainSelectionProof:              bytesutil.Uint32ToBytes4(0x05000000),
	DomainAggregateAndProof:           bytesutil.Uint32ToBytes4(0x06000000),
	DomainSyncCommittee:               bytesutil.Uint32ToBytes4(0x07000000),
	DomainSyncCommitteeSelectionProof: bytesutil.Uint32ToBytes4(0x08000000),
	DomainContributionAndProof:        bytesutil.Uint32ToBytes4(0x09000000),
	DomainApplicationMask:             bytesutil.Uint32ToBytes4(0x00000001),
	DomainApplicationBuilder:          bytesutil.Uint32ToBytes4(0x00000001),
	DomainBLSToExecutionChange:        bytesutil.Uint32ToBytes4(0x0A000000),
	DomainBlobSidecar:                 bytesutil.Uint32ToBytes4(0x0B000000),

	// Prysm constants.
	GweiPerEth:                     1000000000,
	BLSSecretKeyLength:             32,
	BLSPubkeyLength:                48,
	DefaultBufferSize:              10000,
	WithdrawalPrivkeyFileName:      "/shardwithdrawalkey",
	ValidatorPrivkeyFileName:       "/validatorprivatekey",
	RPCSyncCheck:                   1,
	EmptySignature:                 [96]byte{},
	DefaultPageSize:                250,
	MaxPeersToSync:                 15,
	SlotsPerArchivedPoint:          2048,
	GenesisCountdownInterval:       time.Minute,
	ConfigName:                     MainnetName,
	PresetBase:                     "mainnet",
	BeaconStateFieldCount:          21,
	BeaconStateAltairFieldCount:    24,
	BeaconStateBellatrixFieldCount: 25,
	BeaconStateCapellaFieldCount:   28,
	BeaconStateDenebFieldCount:     31,

	// Slasher related values.
	WeakSubjectivityPeriod:          54000,
	PruneSlasherStoragePeriod:       10,
	SlashingProtectionPruningEpochs: 512,

	// Weak subjectivity values.
	SafetyDecay: 10,

	// Fork related values.
	GenesisEpoch:         genesisForkEpoch,
	GenesisForkVersion:   []byte{10, 0, 0, 0},
	AltairForkVersion:    []byte{11, 0, 0, 0},
	AltairForkEpoch:      0,
	BellatrixForkVersion: []byte{12, 0, 0, 0},
	BellatrixForkEpoch:   0,
	CapellaForkVersion:   []byte{13, 0, 0, 0},
	CapellaForkEpoch:     0,
	DenebForkVersion:     []byte{14, 0, 0, 0},
	DenebForkEpoch:       0,

	// New values introduced in Altair hard fork 1.
	// Participation flag indices.
	TimelySourceFlagIndex: 0,
	TimelyTargetFlagIndex: 1,
	TimelyHeadFlagIndex:   2,

	// Incentivization weight values.
	TimelySourceWeight: 14,
	TimelyTargetWeight: 26,
	TimelyHeadWeight:   14,
	SyncRewardWeight:   2,
	ProposerWeight:     8,
	WeightDenominator:  64,

	// Validator related values.
	TargetAggregatorsPerSyncSubcommittee: 16,
	SyncCommitteeSubnetCount:             4,

	// Misc values.
	SyncCommitteeSize:            512,
	InactivityScoreBias:          4,
	InactivityScoreRecoveryRate:  16,
	EpochsPerSyncCommitteePeriod: 256,

	// Updated penalty values.
	InactivityPenaltyQuotientAltair:         3 * 1 << 24, //50331648
	MinSlashingPenaltyQuotientAltair:        64,
	ProportionalSlashingMultiplierAltair:    2,
	MinSlashingPenaltyQuotientBellatrix:     32,
	ProportionalSlashingMultiplierBellatrix: 3,
	InactivityPenaltyQuotientBellatrix:      1 << 24,

	// Light client
	MinSyncCommitteeParticipants: 1,
	MaxRequestLightClientUpdates: 128,

	// Bellatrix
	TerminalBlockHashActivationEpoch: 18446744073709551615,
	TerminalBlockHash:                [32]byte{},
	TerminalTotalDifficulty:          "0", // Estimated: Sept 15, 2022
	EthBurnAddressHex:                "0x0000000000000000000000000000000000000000",
	DefaultBuilderGasLimit:           uint64(30000000),

	// Mevboost circuit breaker
	MaxBuilderConsecutiveMissedSlots: 3,
	MaxBuilderEpochMissedSlots:       5,
	// Execution engine timeout value
	ExecutionEngineTimeoutValue: 8, // 8 seconds default based on: https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md#core

	// Subnet value
	BlobsidecarSubnetCount: 6,

	MaxPerEpochActivationChurnLimit:  8,
	MinEpochsForBlobsSidecarsRequest: 4096,
	MaxRequestBlobSidecars:           768,
	MaxRequestBlocksDeneb:            128,

	// Values related to networking parameters.
	GossipMaxSize:                   10 * 1 << 20, // 10 MiB
	MaxChunkSize:                    10 * 1 << 20, // 10 MiB
	AttestationSubnetCount:          64,
	AttestationPropagationSlotRange: 32,
	MaxRequestBlocks:                1 << 10, // 1024
	TtfbTimeout:                     5,
	RespTimeout:                     10,
	MaximumGossipClockDisparity:     500,
	MessageDomainInvalidSnappy:      [4]byte{00, 00, 00, 00},
	MessageDomainValidSnappy:        [4]byte{01, 00, 00, 00},
	MinEpochsForBlockRequests:       33024, // MIN_VALIDATOR_WITHDRAWABILITY_DELAY + CHURN_LIMIT_QUOTIENT / 2 (= 33024, ~5 months)
	EpochsPerSubnetSubscription:     256,
	AttestationSubnetExtraBits:      0,
	AttestationSubnetPrefixBits:     6,
	SubnetsPerNode:                  2,
	NodeIdBits:                      256,
}

// MainnetTestConfig provides a version of the mainnet config that has a different name
// and a different fork choice schedule. This can be used in cases where we want to use config values
// that are consistent with mainnet, but won't conflict or cause the hard-coded genesis to be loaded.
func MainnetTestConfig() *BeaconChainConfig {
	mn := MainnetConfig().Copy()
	mn.ConfigName = MainnetTestName
	FillTestVersions(mn, 128)
	return mn
}

// FillTestVersions replaces the fork schedule in the given BeaconChainConfig with test values, using the given
// byte argument as the high byte (common across forks).
func FillTestVersions(c *BeaconChainConfig, b byte) {
	c.GenesisForkVersion = make([]byte, fieldparams.VersionLength)
	c.AltairForkVersion = make([]byte, fieldparams.VersionLength)
	c.BellatrixForkVersion = make([]byte, fieldparams.VersionLength)
	c.CapellaForkVersion = make([]byte, fieldparams.VersionLength)
	c.DenebForkVersion = make([]byte, fieldparams.VersionLength)

	c.GenesisForkVersion[fieldparams.VersionLength-1] = b
	c.AltairForkVersion[fieldparams.VersionLength-1] = b
	c.BellatrixForkVersion[fieldparams.VersionLength-1] = b
	c.CapellaForkVersion[fieldparams.VersionLength-1] = b
	c.DenebForkVersion[fieldparams.VersionLength-1] = b

	c.GenesisForkVersion[0] = 0
	c.AltairForkVersion[0] = 1
	c.BellatrixForkVersion[0] = 2
	c.CapellaForkVersion[0] = 3
	c.DenebForkVersion[0] = 4
}
