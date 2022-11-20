package main

import (
	_ "net/http/pprof"

	"github.com/joscelyn/urban"
	"github.com/joscelyn/urban/server"
	"github.com/joscelyn/urban/urban/cmd"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile/v2"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.UrbanServer{}

	lile.Name("urban")
	lile.Server(func(g *grpc.Server) {
		urban.RegisterUrbanServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
