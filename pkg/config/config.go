package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/merlindrones/pit-viper/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate go run ../../tools/gen_config_list/main.go

// DefaultModules is overwritten by code generation. Do not edit manually.
var DefaultModules []IConfig

// IConfig represents a module-specific config with default registration support.
type IConfig interface {
	RegisterDefaults(*viper.Viper)
	String() string
}

// Init loads and initializes all config modules.
func Init(configPath string) error {
	v := viper.GetViper()

	v.SetConfigType("toml")
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetConfigName(pkg.APP_NAME)
		viper.AddConfigPath(path.Join(home, pkg.APP_NAME))    // Looks in ~/.APP_NAME
		viper.AddConfigPath(".")                              // Local dir
		viper.AddConfigPath(pkg.APP_NAME)                     // Looks in ./APP_NAME
		viper.AddConfigPath(path.Join("/etc/", pkg.APP_NAME)) // Looks in /etc/APP_NAME
		viper.AddConfigPath(home)                             // Looks in $HOME
	}
	//Setup Environment Variable Handling
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(pkg.APP_ENV_PREFIX) // optional scoping

	// Register all module defaults
	SetDefaults(v, DefaultModules)

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	} else {
		if e, ok := err.(viper.ConfigParseError); ok { //Config Parsing Error
			fmt.Printf("error parsing config file: %v\n", e)
			//log.Fatal("Exiting!")
		} else if _, ok := err.(*fs.PathError); ok {
			fmt.Printf("Config File Specified at %v Not Found.  Continuing with defaults.\n", configPath)
		} else if _, ok := err.(viper.ConfigFileNotFoundError); ok { // Config file not found; Use defaults
			fmt.Printf("Config File Specified at %v Not Found.  Continuing with defaults.\n", configPath)
		}
	}

	// Replace global viper instance (optional, could keep local `v`)
	//viper.Set("config", v)

	return nil
}

// SetDefaults ensures all module default values are registered.
func SetDefaults(v *viper.Viper, modules []IConfig) {
	for _, m := range modules {
		m.RegisterDefaults(v)
	}
}

func Json() (config string) {
	cb, err := json.MarshalIndent(viper.AllSettings(), "", "   ")
	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(-1)
	}
	return string(cb)
}

func Toml() (config string) {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(viper.AllSettings())
	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(-1)
	}
	return buf.String()
}
