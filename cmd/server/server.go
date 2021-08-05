package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/udp2mysql/configs"
	"github.com/hi20160616/udp2mysql/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	s, err := server.NewUDPReceiver(configs.V.RemoteAddr, 1<<10)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		log.Println("UDP Server start.")
		return s.Start(ctx)
	})
	g.Go(func() error {
		defer log.Printf("UDP Server stop done.")
		<-ctx.Done() // wait for stop signal
		log.Print("UDP Server stop now...")
		return s.Stop(ctx)
	})
	// Elegant stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
