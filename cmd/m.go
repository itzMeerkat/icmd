/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/itzmeerkat/icmd/domain/memo_domain"
	"github.com/itzmeerkat/icmd/domain/memo_domain/model"
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
		memo_domain.AddMemo(model.Memo{
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
}
