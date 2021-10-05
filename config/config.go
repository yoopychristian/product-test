package configuration

import (
	"fmt"
	"time"

	fx "product-test/functions"
	adt "product-test/repo-adaptor"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

//ServiceConfiguration service configuration which to be extract from env vars
type ServiceConfiguration struct {
	App AppConfig
}

type PortalConfiguration struct {
	App AppConfig
}

//AppConfig standard configuration for both service and repository
type AppConfig struct {
	Debug      bool
	Timezone   string
	Port       string
	Location   *time.Location `anonymous:"true"`
	LogPath    string
	Name       string
	NetTimeOut time.Duration
}

//ServiceContext context of service
type ServiceContext struct {
	Config  ServiceConfiguration
	Adaptor adt.RepositoryAdaptor
	Log     *zap.Logger
}

type PortalContext struct {
	Config  PortalConfiguration
	Adaptor adt.RepositoryAdaptor
	Log     *zap.Logger
}

//RepositoryContext context of repository
type RepositoryContext struct {
	Config  RepositoryConfiguration
	Adaptor adt.RepositoryAdaptor
	DB      *gorm.DB
	Log     *zap.Logger
}

//RepositoryConfiguration configuration collection for repositories
type RepositoryConfiguration struct {
	App AppConfig
	DB  DBConfig
}

//EnvDBConfig database configuration which to be extract from env vars
type DBConfig struct {
	Host           string
	Port           string
	Schema         string
	DBName         string
	Username       string
	Password       string
	Logging        bool
	SessionName    string
	ConnectTimeOut int
	MaxOpenConn    int
	MaxIdleConn    int
}

// func GetPortalConfiguration() (PortalConfiguration, error) {
// 	cfg := PortalConfiguration{
// 		App: AppConfig{
// 			LogPath:    fx.EnvString("LOG_PATH"),
// 			Debug:      fx.EnvBool("DEBUG"),
// 			Timezone:   fx.EnvString("TIMEZONE"),
// 			Port:       fx.EnvString("PORT"),
// 			Name:       fx.EnvString("APP_NAME"),
// 			NetTimeOut: time.Duration(fx.EnvInt("NET_TIMEOUT")) * time.Second,
// 		},
// 	}

// 	//default port
// 	if cfg.App.Port == "" {
// 		cfg.App.Port = "8080"
// 	}

// 	//load location
// 	var err error
// 	cfg.App.Location, err = time.LoadLocation(cfg.App.Timezone)
// 	if err != nil {
// 		return PortalConfiguration{}, fmt.Errorf("can't load location :" + cfg.App.Timezone + ",error :" + err.Error())
// 	}

// 	return cfg, nil
// }

// func GetServiceConfiguration() (ServiceConfiguration, error) {
// 	cfg := ServiceConfiguration{
// 		App: AppConfig{
// 			LogPath:    fx.EnvString("LOG_PATH"),
// 			Debug:      fx.EnvBool("DEBUG"),
// 			Timezone:   fx.EnvString("TIMEZONE"),
// 			Port:       fx.EnvString("PORT"),
// 			Name:       fx.EnvString("APP_NAME"),
// 			NetTimeOut: time.Duration(fx.EnvInt("NET_TIMEOUT")) * time.Second,
// 		},
// 	}

// 	//default port
// 	if cfg.App.Port == "" {
// 		cfg.App.Port = "8080"
// 	}

// 	//load location
// 	var err error
// 	cfg.App.Location, err = time.LoadLocation(cfg.App.Timezone)
// 	if err != nil {
// 		return ServiceConfiguration{}, fmt.Errorf("can't load location :" + cfg.App.Timezone + ",error :" + err.Error())
// 	}

// 	return cfg, nil
// }

func GetRepositoryConfiguration() (RepositoryConfiguration, error) {
	cfg := RepositoryConfiguration{
		App: AppConfig{
			LogPath:  fx.EnvString("LOG_PATH"),
			Debug:    fx.EnvBool("DEBUG"),
			Timezone: fx.EnvString("TIMEZONE"),
			Port:     fx.EnvString("PORT"),
			Name:     fx.EnvString("APP_NAME"),
		},
		DB: DBConfig{
			Host:           fx.EnvString("DB_HOST"),
			DBName:         fx.EnvString("DB_NAME"),
			Username:       fx.EnvString("DB_USERNAME"),
			Password:       fx.EnvString("DB_PASSWORD"),
			Logging:        fx.EnvBool("DB_LOGGING"),
			Port:           fx.EnvString("DB_PORT"),
			Schema:         fx.EnvString("DB_SCHEMA"),
			SessionName:    fx.EnvString("DB_SESSION_NAME"),
			ConnectTimeOut: fx.EnvInt("DB_CONNECT_TIMEOUT"),
			MaxOpenConn:    fx.EnvInt("DB_MAX_OPEN_CONN"),
			MaxIdleConn:    fx.EnvInt("DB_MAX_IDLE_CONN"),
		},
	}

	//default port
	if cfg.App.Port == "" {
		cfg.App.Port = "8081"
	}

	//default logging path
	if cfg.App.LogPath == "" {
		cfg.App.LogPath = "/logs/"
	}

	//default db connection time out
	if cfg.DB.ConnectTimeOut == 0 {
		cfg.DB.ConnectTimeOut = 30
	}

	//default db maximum open connection
	if cfg.DB.MaxOpenConn == 0 {
		cfg.DB.MaxOpenConn = 50
	}

	//default db maximum idle connection
	if cfg.DB.MaxIdleConn == 0 {
		cfg.DB.MaxIdleConn = 10
	}

	//load location
	var err error
	cfg.App.Location, err = time.LoadLocation(cfg.App.Timezone)
	if err != nil {
		return RepositoryConfiguration{}, fmt.Errorf("can't load location %s (%w)", cfg.App.Timezone, err)
	}

	return cfg, nil
}
