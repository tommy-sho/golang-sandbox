package main

import (
	"context"
	"flag"
	"log"

	"github.com/tommy-sho/golang-sandbox/golangtokyo-shoten8/raft/store"
)

func main() {
	port := flag.Int("port", 3000, "key-value server port")
	flag.Parse()

	s := store.New()
	srv := New(*port, s)
	log.Println(srv.Run(context.Background()))
}
