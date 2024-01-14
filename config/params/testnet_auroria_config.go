package params

// UseAuroriaNetworkConfig uses the Auroria beacon chain specific network config.
func UseAuroriaNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.BootstrapNodes = []string{
		"enr:-LG4QBlMMXGM0Fhzhvo-VGJB7cQTGqpak90xanmCHfgSxdaLLTQRJRUfQjYEYi7QYmlBvpvh_01zSKvZj-Tg3-qG2GeGAY0I5-oxh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCvLv9SCgAKFP__________gmlkgnY0gmlwhM-axcyJc2VjcDI1NmsxoQIGuTO8PqkEDOW76FlFr6F3pwmLSCJM2_mwEYZe3kXOfIN1ZHCCMsg",
		"enr:-LG4QJ0d7LrA1tlTdjw4p0CCROZH24dsNpZcAja5BWZC2FcxTKKmYts-TBf-JBj599xNeNGZzHQQYCi7bwvSVrlHQJ2GAY0I663mh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCvLv9SCgAKFP__________gmlkgnY0gmlwhJB-yIqJc2VjcDI1NmsxoQMNFuK9zAIMUvQV5RDcaXEiRxNswWhnJVQX4q3cg1qPG4N1ZHCCMsg",
		"enr:-LG4QKJ1RYeQBSfJuks1Izqhk_Adl5j5fqvn9-Zpw4MiqqOFU3FJ_6AP2m-qoLLVixPTfVmm9Sof_FVYMklf9Jf28-aGAY0JAVxFh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCvLv9SCgAKFP__________gmlkgnY0gmlwhIDHP-OJc2VjcDI1NmsxoQL5xwf9tI-ajIU6_-vxCmHtm9He4Z-vOa0oMF3I8MC1SoN1ZHCCMsg",
		"enr:-LG4QOJEt68bobkI0hvwI1jhcRAEdsLaQi-nseuKyQKY6QAbEmJ2_mwqXrGm4Wgh2OFvH7R4aJfcW7994799cJMGhDmGAY0I_tWdh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCvLv9SCgAKFP__________gmlkgnY0gmlwhIrFfCuJc2VjcDI1NmsxoQK_fk2uDvKEmEh9Nv57hhyLsle9tiIXaQ1zRUzIDlcMxIN1ZHCCMsg",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// AuroriaConfig defines the config for the Testnet Auroria beacon chain.
func AuroriaConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = AuroriaName
	cfg.GenesisForkVersion = []byte{10, 0, 10, 20}
	cfg.DepositChainID = 205205
	cfg.DepositNetworkID = 205205
	cfg.AltairForkVersion = []byte{11, 0, 10, 20}
	cfg.BellatrixForkVersion = []byte{12, 0, 10, 20}
	cfg.CapellaForkVersion = []byte{13, 0, 10, 20}
	cfg.DenebForkVersion = []byte{14, 0, 10, 20}

	cfg.InitializeForkSchedule()
	return cfg
}
