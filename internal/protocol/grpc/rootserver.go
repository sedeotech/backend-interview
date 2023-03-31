package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/sedeotech/backend-interview/internal/config"
	"github.com/sedeotech/backend-interview/internal/services/cart"
	"github.com/sedeotech/backend-interview/internal/storage/postgres"
	"github.com/sedeotech/backend-interview/pkg/rpc/cartpb"
	"google.golang.org/grpc"
)

type RootServer struct {
	port   int
	server *grpc.Server
	ctx    context.Context
}

func NewRootServer(ctx context.Context, cfg config.Config) (*RootServer, error) {

	// Services initialisation
	///////////////////////////

	store, err := postgres.NewRepository(cfg)
	if err != nil {
		return nil, err
	}

	cartService := cart.NewCartService(store)
	//itemService := item.NewItemService(store)

	// Server registration
	///////////////////////
	s := grpc.NewServer()

	cartpb.RegisterCartServer(s, NewCartServer(cartService))
	//itempb.RegisterItemServer(s, NewItemServer(itemService))

	return &RootServer{
		port:   cfg.GRPCPort,
		server: s,
	}, nil
}

func (g *RootServer) Start() error {

	fmt.Println("start server", g.port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return err
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			//logger.Log.Warn("shutting down gRPC server...")

			g.server.GracefulStop()
			<-g.ctx.Done()
		}
	}()

	return g.server.Serve(lis)
}
