/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/itzmeerkat/icmd/domain/memo_domain"
	"strings"

	"github.com/spf13/cobra"
)

var isTodo bool
var isWishlist bool

// mCmd represents the m command
var mCmd = &cobra.Command{
	Use:   "m",
	Short: "all about memos",
	Long:  `i m whatever content u want to input [-flags]`,
	Run: func(cmd *cobra.Command, args []string) {
		c := strings.Join(args, " ")
		memo_domain.AddMemo(memo_domain.Memo{
			Content:    c,
			IsWishlist: isWishlist,
			IsTodo:     isTodo,
			Status:     0,
		})
	},
}

func init() {
	rootCmd.AddCommand(mCmd)
	mCmd.PersistentFlags().BoolVarP(&isTodo, "todo", "t", false, "if its a todo entry")
	mCmd.PersistentFlags().BoolVarP(&isWishlist, "wishlist", "w", false, "if its a wishlist entry")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
