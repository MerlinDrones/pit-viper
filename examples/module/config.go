package module

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ModuleConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func NewModuleConfig() *ModuleConfig {
	cfg := &ModuleConfig{}

	if sub := viper.Sub("module"); sub != nil {
		if err := sub.Unmarshal(cfg); err != nil {
			log.WithError(err).Error("failed to unmarshal module config")
		}
		return cfg
	}

	// fallback if viper.Sub("module") == nil
	if err := viper.UnmarshalKey("module", cfg); err != nil {
		log.WithError(err).Error("failed to unmarshal module config from root")
	}

	return cfg
}

func (mc *ModuleConfig) RegisterDefaults(v *viper.Viper) {
	v.SetDefault("module.host", "localhost")
	v.SetDefault("module.port", 8080)
}

func (mc *ModuleConfig) String() string {
	b, _ := json.Marshal(mc)
	return string(b)
}
