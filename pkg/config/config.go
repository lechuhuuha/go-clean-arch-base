package config

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/viper"

	"github.com/lechuhuuha/oneshield/constant"
	"github.com/lechuhuuha/oneshield/pkg/utils"
)

var (
	NumCPU  = runtime.NumCPU()
	Version = "dev"
)

type (
	Config struct {
		Env *Env `json:"env" mapstructure:"env" validate:"required"`
		Log *Log `json:"log" mapstructure:"log" validate:"required"`
	}

	Secret struct {
		AccessToken      string        `json:"access_token" mapstructure:"access_token" validate:"required"`
		AccessTokenTTL   time.Duration `json:"access_token_ttl" mapstructure:"access_token_ttl" validate:"required"`
		GlobalSessionTTL time.Duration `json:"global_session_ttl" mapstructure:"global_session_ttl" validate:"required"`
	}

	Env struct {
		IsStaging   bool   `json:"is_staging" mapstructure:"is_staging"`
		Test        string `json:"test"`
		TelegramKey string `json:"telegram_key" mapstructure:"telegram_key"`
	}
	Log struct {
		Path       string `json:"path" mapstructure:"path" validate:"required"`
		MaxSize    int    `json:"max_size" mapstructure:"max_size" validate:"required,gt=0"`
		MaxAge     int    `json:"max_age" mapstructure:"max_age" validate:"required,gt=0"`
		MaxBackups int    `json:"max_backups" mapstructure:"max_backups" validate:"required,gt=0"`
	}

	Services struct {
		SharedServiceAddr string `json:"shared_service_addr" mapstructure:"shared_service_addr"`
	}

	Backup struct {
		MaxKeepDeletedUserBackupDays int64 `json:"max_keep_backup_days" mapstructure:"max_keep_backup_days" validate:"required,gt=1"`
	}
)

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/oneshield/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}
	viper.SetDefault("services.shared_service_addr", constant.HTTPAddress)
	viper.SetDefault("secret.global_session_ttl", 24*time.Hour)
	viper.SetDefault("db.driver_name", "mysql")
	viper.SetDefault("backup.max_keep_backup_days", constant.MaxKeepBackupDays)
	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	utils.SetTimezone()
	return &c, nil
}
