package cmd

import (
	"github.com/bobasensei/babu/cmd/fetch"
	"github.com/bobasensei/babu/cmd/get"
	"github.com/bobasensei/babu/cmd/list"
	"github.com/bobasensei/babu/cmd/store"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "babu",
		Short: "crawl and collect results from wikipedia",
	}
	cmd.AddCommand(fetch.Cmd())
	cmd.AddCommand(get.Cmd())
	cmd.AddCommand(list.Cmd())
	cmd.AddCommand(store.Cmd())
	return cmd
}
