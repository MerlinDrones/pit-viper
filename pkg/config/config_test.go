package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"pit-viper/pkg/config"
)

func resetEnv(keys ...string) {
	for _, k := range keys {
		_ = os.Unsetenv(k)
	}
}

func TestInit_WithDefaults(t *testing.T) {
	resetEnv("PIT_MODULE_HOST", "PIT_MODULE_PORT")
	viper.Reset()

	err := config.Init("")
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	host := viper.GetString("module.host")
	port := viper.GetInt("module.port")

	if host != "localhost" {
		t.Errorf("expected host = localhost, got %s", host)
	}
	if port != 8080 {
		t.Errorf("expected port = 8080, got %d", port)
	}
}

func TestInit_WithConfigFile(t *testing.T) {
	viper.Reset()

	tmp := t.TempDir()
	configFile := filepath.Join(tmp, "config.toml")
	content := `
[module]
host = "from-file"
port = 9999
`
	if err := os.WriteFile(configFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	_ = os.Setenv("cfgFile", configFile)
	defer resetEnv("cfgFile")

	err := config.Init(configFile)
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if got := viper.GetString("module.host"); got != "from-file" {
		t.Errorf("expected host from config file, got %s", got)
	}
	if got := viper.GetInt("module.port"); got != 9999 {
		t.Errorf("expected port from config file, got %d", got)
	}
}

func TestInit_WithEnvVars(t *testing.T) {
	viper.Reset()
	resetEnv("PIT_MODULE_HOST", "PIT_MODULE_PORT")

	_ = os.Setenv("PIT_MODULE_HOST", "envhost")
	_ = os.Setenv("PIT_MODULE_PORT", "3030")
	defer resetEnv("PIT_MODULE_HOST", "PIT_MODULE_PORT")

	err := config.Init("")
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if got := viper.GetString("module.host"); got != "envhost" {
		t.Errorf("expected env var override host = envhost, got %s", got)
	}
	if got := viper.GetInt("module.port"); got != 3030 {
		t.Errorf("expected env var override port = 3030, got %d", got)
	}
}
