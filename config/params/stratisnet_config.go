package params

// UseStratisNetworkConfig uses the Stratis beacon chain specific network config.
func UseStratisNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{}
	OverrideBeaconNetworkConfig(cfg)
}

// StratisConfig defines the config for the Stratis beacon chain.
func StratisConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.MinGenesisTime = 1655647200
	cfg.GenesisDelay = 1
	cfg.MinGenesisActiveValidatorCount = 1300
	cfg.ConfigName = StratisName
	cfg.GenesisForkVersion = []byte{0x20, 0x00, 0x00, 0x69}
	cfg.SecondsPerETH1Block = 14
	cfg.DepositChainID = 56200
	cfg.DepositNetworkID = 56200
	cfg.AltairForkEpoch = 0
	cfg.AltairForkVersion = []byte{0x20, 0x00, 0x00, 0x70}
	cfg.BellatrixForkEpoch = 0
	cfg.BellatrixForkVersion = []byte{0x20, 0x00, 0x00, 0x71}
	cfg.CapellaForkEpoch = 0
	cfg.CapellaForkVersion = []byte{0x20, 0x00, 0x00, 0x72}
	cfg.DenebForkEpoch = 0
	cfg.DenebForkVersion = []byte{0x20, 0x00, 0x00, 0x73}
	cfg.TerminalTotalDifficulty = "0"
	cfg.DepositContractAddress = "0x0000000000000000000000000000000000001001"
	cfg.MaxEffectiveBalance = 20000 * 1e9
	cfg.EjectionBalance = 10000 * 1e9
	cfg.ProposerBlockReward = 60 * 1e9

	// Changed for faster testing
	// TODO: remove once testing is finished
	//cfg.SecondsPerSlot = 3
	//cfg.SlotsPerEpoch = 2
	//cfg.Eth1FollowDistance = 1
	//cfg.EpochsPerEth1VotingPeriod = 1
	//cfg.MinValidatorWithdrawabilityDelay = 2
	//cfg.ShardCommitteePeriod = 2

	cfg.InitializeForkSchedule()
	return cfg
}
