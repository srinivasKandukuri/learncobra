/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srinivasKandukuri/learncobraapp/pkg/stringer"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Aliases: []string{"rev"},
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reverse called")
		res := stringer.Reverse(args[0])
        fmt.Println(res)
	},
}

func init() {
	reverseCmd.Flags().BoolP("detaills", "d", false, "details message for toggle")
	rootCmd.AddCommand(reverseCmd)
}
