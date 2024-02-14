/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/itzmeerkat/icmd/domain/memo_domain"
	"github.com/itzmeerkat/icmd/domain/memo_domain/model"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all memos",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var res []model.Memo
		if isWishlist == true && isTodo == false {
			res = memo_domain.GetAllWishlist()
		} else if isWishlist == false && isTodo == true {
			res = memo_domain.GetAllTodo()
		} else if isWishlist == false && isTodo == false {
			res = memo_domain.GetAllPlain()
		} else {
			panic("Don't query wishlist and todo at once")
		}
		t := table.NewWriter()
		for _, e := range res {
			t.AppendRow(table.Row{e.DisplayId, e.Content})
		}
		fmt.Print(t.Render())
	},
}

func init() {
	mCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
