package app

import (
	wasmkeeper "github.com/JackalLabs/jackal-wasmd/x/wasm/keeper"
)

const (
	// DefaultJackalInstanceCost is initially set the same as in wasmd
	DefaultJackalInstanceCost uint64 = 60_000
	// DefaultJackalCompileCost set to a large number for testing
	DefaultJackalCompileCost uint64 = 3
)

// JackalGasRegisterConfig is defaults plus a custom compile amount
func JackalGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultJackalInstanceCost
	gasConfig.CompileCost = DefaultJackalCompileCost

	return gasConfig
}

func NewJackalWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(JackalGasRegisterConfig())
}
