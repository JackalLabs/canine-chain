package emconfig

import (
	"time"

	emconfigerrors "github.com/jackalLabs/canine-chain/x/ethercoat/liteemserver/configerrors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/viper"
)

// Config defines the server's top level configuration
type Config struct {
	// ETH-INSERT START
	JSONRPC JSONRPCConfig `mapstructure:"json-rpc"`
	// ETH-INSERT END
}

// ETH-INSERT START
// JSONRPCConfig defines configuration for the eth-like RPC server.
type JSONRPCConfig struct {
	// API defines a list of JSON-RPC namespaces that should be enabled
	API []string `mapstructure:"api"`
	// Address defines the HTTP server to listen on
	Address string `mapstructure:"address"`
	// Enable defines if the EVM RPC server should be enabled.
	Enable bool `mapstructure:"enable"`
	// HTTPTimeout is the read/write timeout of http json-rpc server.
	HTTPTimeout time.Duration `mapstructure:"http-timeout"`
	// HTTPIdleTimeout is the idle timeout of http json-rpc server.
	HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`
}

// ETH-INSERT END

// DefaultConfig returns server's default configuration.
func DefaultConfig() *Config {
	return &Config{
		// ETH-INSERT START
		JSONRPC: JSONRPCConfig{
			API:             []string{"eth"},
			Address:         "0.0.0.0:8545",
			Enable:          true,
			HTTPTimeout:     30 * time.Second,
			HTTPIdleTimeout: 120 * time.Second,
		},
		// ETH-INSERT END
	}
}

// ValidateBasic returns an error if min-gas-prices field is empty in BaseConfig. Otherwise, it returns nil.
func (c Config) ValidateBasic() error {
	//ETH-INSERT START
	if c.JSONRPC.Enable && len(c.JSONRPC.API) == 0 {
		return sdkerrors.Wrapf(emconfigerrors.ErrBlankAPINamespace, "cannot enable JSON-RPC without defining any API namespaces")
	}
	if c.JSONRPC.HTTPTimeout < 0 {
		return sdkerrors.Wrapf(emconfigerrors.ErrNegativeHTTPTimeout, "JSON-RPC HTTP timeout duration cannot be negative")
	}
	if c.JSONRPC.HTTPIdleTimeout < 0 {
		return sdkerrors.Wrapf(emconfigerrors.ErrNegativeHTTPIdleTimeout, "JSON-RPC HTTP idle timeout duration cannot be negative")
	}
	//ETH-INSERT END

	return nil
}

// GetConfig returns a fully parsed Config object.
func GetConfig(v *viper.Viper) Config {
	globalLabelsRaw := v.Get("telemetry.global-labels").([]interface{})
	globalLabels := make([][]string, 0, len(globalLabelsRaw))
	for _, glr := range globalLabelsRaw {
		labelsRaw := glr.([]interface{})
		if len(labelsRaw) == 2 {
			globalLabels = append(globalLabels, []string{labelsRaw[0].(string), labelsRaw[1].(string)})
		}
	}
	return Config{
		//ETH-INSERT START
		JSONRPC: JSONRPCConfig{
			API:             v.GetStringSlice("json-rpc.api"),
			Address:         v.GetString("json-rpc.address"),
			Enable:          v.GetBool("json-rpc.enable"),
			HTTPTimeout:     v.GetDuration("json-rpc.http-timeout"),
			HTTPIdleTimeout: v.GetDuration("json-rpc.http-idle-timeout"),
		},
		//ETH-INSERT END
	}
}
