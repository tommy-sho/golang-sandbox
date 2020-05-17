package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/tommy-sho/golang-sandbox/golangtokyo-shoten8/raft/raftalg"

	"github.com/tommy-sho/golang-sandbox/golangtokyo-shoten8/raft/store"
)

func main() {
	port := flag.Int("port", 3000, "key-value server port")
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	flag.Parse()

	ra := raftalg.New(*id, strings.Split(*cluster, ","))
	s := store.New(ra)
	srv := New(*port, s)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return ra.Run(ctx)
	})

	eg.Go(func() error {
		return s.RunCommitReader(ctx)
	})

	eg.Go(func() error {
		return srv.Run(ctx)
	})

	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("received signal")
		cancel()
	case <-ctx.Done():
	}

	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

	log.Println("done")

	log.Println(srv.Run(context.Background()))
}
