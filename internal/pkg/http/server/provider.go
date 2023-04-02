package server

import (
	"airbnb-messaging-be/internal/pkg/http/server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Instance = "HTTP Server"

type Options struct {
	config.Config
	Router *gin.Engine
}

type Server struct {
	address string
	server  *http.Server
	Options
}

func NewServer(options Options) *Server {
	return &Server{Options: options}
}
