package cmd

import (
	"log"
	"os"

	"github.com/lileio/lile/v2"
)

var RootCmd = lile.BaseCommand("urban", "A gRPC based service that proxy Hacker News to share top ten stories and user information")

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}
