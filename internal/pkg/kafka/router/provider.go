package router

import (
	"airbnb-messaging-be/internal/pkg/kafka/router/config"
	"context"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const Instance string = "Kafka Router"

type EventHandler func(ctx context.Context, msg *kafka.Message)

type Handler struct {
	Topic   string
	Handler EventHandler
}

type Options struct {
	config.Config
}

type Router struct {
	Options
	basePath string
	Handlers []Handler
}

func NewRouter(options Options) *Router {
	// set default separator
	separator := options.Separator
	if separator == "" {
		separator = "."
	}

	return &Router{
		Options: options,
	}
}

func (r *Router) Group(relativePath string) *Router {
	return &Router{
		basePath: r.calculateAbsolutePath(relativePath),
	}
}

func (r *Router) Listen(relativePath string, handler EventHandler) {
	absolutePath := r.calculateAbsolutePath(relativePath)
	r.Handlers = append(r.Handlers, Handler{
		Topic:   absolutePath,
		Handler: handler,
	})
}

func (r *Router) calculateAbsolutePath(relativePath string) string {
	return r.joinPaths(r.basePath, relativePath)
}

func (r *Router) joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	return strings.Join([]string{absolutePath, relativePath}, r.Separator)
}
