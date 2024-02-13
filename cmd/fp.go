/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/itzmeerkat/icmd/domain/format_domain"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// fpCmd represents the fp command
var fpCmd = &cobra.Command{
	Use:   "fp",
	Short: "try to format a json string",
	Long:  `don't even need to be a json string, just doing my best to make it readable'`,
	Run: func(cmd *cobra.Command, args []string) {
		var input format_domain.ByteArr
		input, _ = io.ReadAll(os.Stdin)
		res := input.HighlightKeywords().FormatJson()
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(fpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
