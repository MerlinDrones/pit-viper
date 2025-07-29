package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"pit-viper/examples/module"
	"pit-viper/pkg/config"
)

// versionCmd represents the version command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Mock runner to proove config DI",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize config
		err := config.Init(viper.GetString("cfgFile"))
		if err != nil {
			cmd.PrintErrln("Failed to load config:", err)
			return
		}

		module := module.NewModule()
		fmt.Println("Current Module is configured as:")
		fmt.Println(module.String())
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
}
