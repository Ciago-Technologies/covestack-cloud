// kan-agent polls the assigned account's read-only work endpoint.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ciagotech/kan-agent/internal/agent"
	"github.com/ciagotech/kan-agent/internal/client"
	"github.com/ciagotech/kan-agent/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	workClient, err := client.New(cfg)
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return agent.Run(ctx, workClient)
}
