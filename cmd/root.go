package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(func() {
		configDir, err := os.UserConfigDir()
		if err != nil {
			panic(err)
		}

		viper.AddConfigPath(configDir + "/pull-request")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	})
}

var rootCmd = &cobra.Command{
	Use:   "pull-request",
	Short: "Create GitHub pull request via the command line",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
