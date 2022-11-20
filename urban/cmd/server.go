package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lileio/lile/v2"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server that proxy hacker news",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			lile.Run()
		}()

		<-c
		lile.Shutdown()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
