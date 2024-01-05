package params

// UseAuroriaNetworkConfig uses the Auroria beacon chain specific network config.
func UseAuroriaNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.BootstrapNodes = []string{
		"enr:-MK4QOiposnqkGSkVcVXeS2RvaMOgLOWFRxJUGyIBQE5y7OITeshOx324-ziiP5rMXtq2UQzEEzBeVCV9x9qkZchGWKGAYx-0AIKh2F0dG5ldHOIAAAAAwAAAACEZXRoMpAZsrCWIAAAc___________gmlkgnY0gmlwhM-axcyJc2VjcDI1NmsxoQNtJ-08vT80nEn18osacUCGm7n7cC_AnQWe0lD6jSKp9YhzeW5jbmV0cw-DdGNwgjLIg3VkcIIu4A",
		"enr:-MK4QC0RZEk8vnHDV_r545j_fr0sMJiFDTg5XzjvMGvNpWC8Z7CxHXYIk8v5QQv98hag9GhwGfsS4i0gMgpuTOq0ddOGAYx-0Ebwh2F0dG5ldHOIgAEAAAAAAACEZXRoMpAZsrCWIAAAc___________gmlkgnY0gmlwhJB-yIqJc2VjcDI1NmsxoQJc7jQivE8v8maM__IlAQjRlk3fUjQ6g28LDm_yvuFTVIhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
		"enr:-MK4QHG-l9mhI4DxC9--wHLLzovKe_lIk7r01pQQph-u3Vk0WbKuBsiAFjsbu_DklyZGhSmt9f2s9LAAJfihMSG7VJGGAYyRVdKVh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAZsrCWIAAAc___________gmlkgnY0gmlwhIDHP-OJc2VjcDI1NmsxoQKgwqbQpccs8JHWpMolTsmXoMBr4YfGh-GX_8nRBIalgIhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
		"enr:-MK4QMxmEX5gwI0jXHxfYTZLgvzeC3_EIoiafqOse9h6xHRzIJZE8hw0qiuDAGtx4qFR4ZRq02Vm-JXY-gIbgGTJFXmGAYyRVsuYh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAZsrCWIAAAc___________gmlkgnY0gmlwhIrFfCuJc2VjcDI1NmsxoQKPDJLts_efCOK3I-us6SQCiqWdg8kY7LimC7etC8KF1ohzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
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
