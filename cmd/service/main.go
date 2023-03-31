package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sedeotech/backend-interview/internal/config"
	"github.com/sedeotech/backend-interview/internal/protocol/grpc"
)

func main() {

	if err := run(); err != nil {

		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {

	ctx := context.Background()

	// Parse the configuration
	cfg := config.Parse()

	g, err := grpc.NewRootServer(ctx, cfg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		g.Start()

	}

	return err
}
