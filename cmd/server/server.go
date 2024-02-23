package server

import (
	"Test_task_solidgate_junior/cmd/api"
	"Test_task_solidgate_junior/config"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
	"time"
)

type Server struct {
	srv *http.Server
	api api.Handle
}

func NewServer() Server {
	return Server{}
}

func (s Server) Start(ctx context.Context) error {
	// Get configs
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("can not parse config: %s", err)
	}

	s.api = api.NewApi()

	s.srv = &http.Server{
		Addr:    cfg.Server.Address,
		Handler: s.registerRoutes(),
	}

	// Start http server
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	select {
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		return s.srv.Shutdown(timeout)
	}
}

func (s Server) registerRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		runtime.Gosched()
		c.JSON(http.StatusOK, "pong")
	})

	postsRoutes := router.Group("/card")
	{
		postsRoutes.POST("/validate", s.api.ValidateCard)
	}

	return router
}
