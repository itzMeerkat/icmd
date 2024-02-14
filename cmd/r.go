/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/itzmeerkat/icmd/acl"
	"github.com/itzmeerkat/icmd/domain/git_repo_domain"
	"github.com/itzmeerkat/icmd/domain/memo_domain"
	"github.com/itzmeerkat/icmd/domain/page_render_domain"
	"github.com/spf13/cobra"
)

var wetRun bool

// renderCmd represents the render command
var rCmd = &cobra.Command{
	Use:   "r",
	Short: "render memos to single HTML page",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		page := acl.NewPageContent(memo_domain.GetAllWishlist(), memo_domain.GetAllTodo(), memo_domain.GetAllPlain())
		newPage := page_render_domain.RenderPage(page)
		git_repo_domain.UpdateRemote(newPage, wetRun)
	},
}

func init() {

	rCmd.Flags().BoolVar(&wetRun, "wetRun", false, "set this flag to actually update git repo")
	rCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		err := command.Flags().MarkHidden("todo")
		if err != nil {
			panic(err)
		}
		err = command.Flags().MarkHidden("wishlist")
		if err != nil {
			panic(err)
		}
		command.Parent().HelpFunc()(command, strings)
	})
	mCmd.AddCommand(rCmd)
}
