package main

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/lechuhuuha/oneshield/pkg/config"
	"github.com/lechuhuuha/oneshield/pkg/validate"
)

func main() {
	totalCPUToUse := runtime.NumCPU() - 1
	runtime.GOMAXPROCS(totalCPUToUse)
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	if err := validate.Validate(cfg); err != nil {
		panic(err)
	}

	out := &lumberjack.Logger{
		Filename:   cfg.Log.Path,
		MaxSize:    cfg.Log.MaxSize, // megabytes
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge, // days
	}
	defer func(out *lumberjack.Logger) {
		_ = out.Close()
	}(out)

	log.Logger = log.Output(&zerolog.ConsoleWriter{
		Out:        out,
		NoColor:    true,
		TimeFormat: time.RFC3339,
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}).With().Caller().Logger()

	if !cfg.Env.IsStaging {
		log.Logger = log.Logger.Level(zerolog.WarnLevel)
		gin.SetMode(gin.ReleaseMode)
	}
	app, cleanup, err := NewService(cfg)
	if err != nil {
		panic(err)
	}

	defer cleanup()

	if err = app.Run(); err != nil {
		log.Info().Err(err).Msg("failed to run app")
	}

	time.Sleep(time.Second)
}
