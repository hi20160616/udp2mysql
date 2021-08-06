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

	// UDPReceiver
	s, err := server.NewUDPReceiver(configs.V.RemoteAddr, 1024)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		log.Println("UDP Server start.")
		return s.Start(ctx)
	})
	g.Go(func() error {
		defer log.Println("UDP Server stop done.")
		<-ctx.Done() // wait for stop signal
		log.Println("UDP Server stop now...")
		return s.Stop(ctx)
	})

	// gRPC
	gs, err := server.NewGRPCServer()
	if err != nil {
		log.Printf("%v", err)
	}
	g.Go(func() error {
		log.Println("gRPC Server start.")
		return gs.Start(ctx)
	})
	g.Go(func() error {
		defer log.Println("gRPC Server stop done.")
		<-ctx.Done()
		log.Println("gRPC Server stop now...")
		return gs.Stop(ctx)
	})

	// Web
	ws := server.NewWebServer()
	g.Go(func() error {
		log.Println("Web Server start.")
		return ws.Start(ctx)
	})
	g.Go(func() error {
		defer log.Println("Web Server stop done.")
		<-ctx.Done()
		log.Println("Web Server stop now...")
		return ws.Stop(ctx)
	})

	// Graceful stop
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
