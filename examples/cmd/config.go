package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"pit-viper/pkg/config"
)

// versionCmd represents the version command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Displays pit-viper version information",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize config
		err := config.Init(viper.GetString("cfgFile"))
		if err != nil {
			cmd.PrintErrln("Failed to load config:", err)
			return
		}
		fmt.Println("------- Begin Config -------")
		fmt.Print(config.Toml())
		fmt.Println("------- End Config -------")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
