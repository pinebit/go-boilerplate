package http

import (
	net_http "net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router interface {
	Handler() net_http.Handler
}

type router struct {
	engine *gin.Engine
}

var (
	promHelloWorldCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "hello_world",
		Help: "Number of querying 'hello world' endpoint",
	})
)

func NewRouter(logger logger.Logger, config *config.Config) Router {
	if config.DevMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gr := gin.New()

	gr.Use(ginzap.Ginzap(logger.Zap(), time.RFC3339, true))

	gr.Use(ginzap.RecoveryWithZap(logger.Zap(), true))

	gr.GET("/", func(c *gin.Context) {
		promHelloWorldCounter.Inc()
		c.String(200, "Hello world!")
	})

	gr.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	return &router{
		engine: gr,
	}
}

func (r router) Handler() net_http.Handler {
	return r.engine
}
