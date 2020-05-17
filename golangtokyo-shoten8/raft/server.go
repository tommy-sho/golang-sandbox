package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
)

type Store interface {
	Lookup(key string) (string, bool)
	Save(key, value string) error
}

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type handler struct {
	store Store
}

func (h *handler) Get(c *gin.Context) {
	key := c.Param("key")
	v, ok := h.store.Lookup(key)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{key: v})
}

func (h *handler) Put(c *gin.Context) {
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := h.store.Save(req.Key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		req.Key: req.Value,
	})

}

type Server struct {
	server http.Server
}

func New(port int, kv Store) *Server {
	h := &handler{
		store: kv,
	}

	r := gin.Default()
	r.GET("/:key", h.Get)
	r.PUT("/", h.Put)

	return &Server{
		server: http.Server{
			Addr:    fmt.Sprintf(":%s", strconv.Itoa(port)),
			Handler: r,
		},
	}
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return s.server.ListenAndServe()
	})

	<-ctx.Done()
	sCtx, sCancel := context.WithTimeout(
		context.Background(), 10*time.Second)
	defer sCancel()

	if err := s.server.Shutdown(sCtx); err != nil {
		return err
	}

	return eg.Wait()
}
