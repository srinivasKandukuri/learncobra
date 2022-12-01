/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package stringer

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "learncobraapp",
	Short: "A brief description of your application",
	Long: `A longer `,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig(){
	configFile := ".cobra-cli-samples.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("LEARNCOBRA")
	
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}

}
