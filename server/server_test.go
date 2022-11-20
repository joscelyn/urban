package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/joscelyn/urban"
	"github.com/lileio/lile/v2"
)

var s = UrbanServer{}
var cli urban.UrbanClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		urban.RegisterUrbanServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = urban.NewUrbanClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
