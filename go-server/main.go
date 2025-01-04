package main

import (
	"context"
	"fmt"
	"os"

	"github.com/0xBradock/go-srvr/server/matt"
)

func main() {
	ctx := context.Background()

	if err := matt.Run(ctx, os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
