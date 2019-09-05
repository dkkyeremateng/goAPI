package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// New returns new HTTP server
func New(engine *gin.Engine, serverAddress string) *http.Server {
	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		// TLSConfig:    tlsConfig,
		Handler: engine,
	}
	return srv
}
