package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	cachecontrol "go.eigsys.de/gin-cachecontrol/v2"

	"github.com/lechuhuuha/oneshield/constant"
	"github.com/lechuhuuha/oneshield/pkg/config"
	"github.com/lechuhuuha/oneshield/pkg/sperrors"
	"github.com/lechuhuuha/oneshield/transports/http/handlers/middleware"
	"github.com/lechuhuuha/oneshield/transports/http/handlers/private"
	"github.com/lechuhuuha/oneshield/transports/http/handlers/response"
)

var (
	ProviderSet = wire.NewSet(NewServer, private.ProviderSet)
)

type Server struct {
	app               *gin.Engine
	svr               *http.Server
	sharedServiceAddr string
	isStaging         bool
}

func NewServer(
	conf *config.Config,
) *Server {
	app := gin.New()
	svr := &http.Server{
		Handler:           app.Handler(),
		ReadTimeout:       time.Hour,
		ReadHeaderTimeout: time.Second * 5,
		WriteTimeout:      time.Hour,
		IdleTimeout:       time.Hour,
		MaxHeaderBytes:    1 << 20,
		Addr:              ":28091",
	}
	app.Use(
		gin.LoggerWithWriter(log.Logger.Level(zerolog.DebugLevel)),
		gin.RecoveryWithWriter(log.Logger, func(c *gin.Context, err any) {
			response.HandleError(
				c,
				http.StatusInternalServerError,
				response.FormatError(
					"ERROR",
					"Internal Server Error",
					fmt.Sprintf("%v", err),
					"",
					sperrors.CodeInternalServer,
					nil,
				),
			)
		}),
	)

	if err := app.SetTrustedProxies([]string{
		"127.0.0.1",
	}); err != nil {
		panic(err)
	}

	return &Server{
		app:       app,
		svr:       svr,
		isStaging: conf.Env.IsStaging,
	}
}

func (s *Server) initRoutes() {
	s.app.Use(
		cachecontrol.New(cachecontrol.Config{
			Public:    true,
			MaxAge:    cachecontrol.Duration(time.Minute * 10),
			Immutable: true,
		}),
	)
	prom := middleware.NewPrometheus("zmonitor")
	s.app.Use(prom.HandlerFunc())

	s.privateRoutes()
	s.instrumentRoutes(prom)
}
func (s *Server) instrumentRoutes(prom *middleware.Prometheus) {
	prom.Use(s.app.Group("/"))
}

func (s *Server) privateRoutes() {
	_ = NewRouter(s.app.Group(
		"/api/private/",
	))
}

func (s *Server) Start() error {
	s.initRoutes()
	log.Info().Msgf("http server is listening on %s", constant.HTTPAddress)

	return s.svr.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.svr.Close()
}
