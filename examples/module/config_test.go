package module_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"pit-viper/examples/module"
)

func resetEnv(keys ...string) {
	for _, k := range keys {
		_ = os.Unsetenv(k)
	}
}

func TestNewModuleConfig_Defaults(t *testing.T) {
	viper.Reset()
	viper.SetDefault("module.host", "localhost")
	viper.SetDefault("module.port", 8080)

	cfg := module.NewModuleConfig()

	if cfg.Host != "localhost" {
		t.Errorf("expected default host = localhost, got %s", cfg.Host)
	}
	if cfg.Port != 8080 {
		t.Errorf("expected default port = 8080, got %d", cfg.Port)
	}
}

func TestNewModuleConfig_Overrides(t *testing.T) {
	viper.Reset()
	viper.Set("module.host", "override-host")
	viper.Set("module.port", 1234)

	cfg := module.NewModuleConfig()

	if cfg.Host != "override-host" {
		t.Errorf("expected overridden host = override-host, got %s", cfg.Host)
	}
	if cfg.Port != 1234 {
		t.Errorf("expected overridden port = 1234, got %d", cfg.Port)
	}
}

func TestRegisterDefaults(t *testing.T) {
	v := viper.New()
	cfg := &module.ModuleConfig{}
	cfg.RegisterDefaults(v)

	if got := v.GetString("module.host"); got != "localhost" {
		t.Errorf("expected module.host = localhost, got %s", got)
	}
	if got := v.GetInt("module.port"); got != 8080 {
		t.Errorf("expected module.port = 8080, got %d", got)
	}
}
