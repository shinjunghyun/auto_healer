package config

import (
	"auto_healer/configs"
	"auto_healer/internal/pkg/env"
	log "logger"
	"sync"
)

var (
	LoggerConfig     log.LoggerOptions
	loggerConfigOnce sync.Once
)

type Logger struct {
	Level    log.LoggerLevel
	LogStyle string
}

func DefaultLoggerConfigFromEnv() log.LoggerOptions {
	loggerConfigOnce.Do(func() {
		LoggerConfig = log.LoggerOptions{
			Level: log.GetLogLevelFromString(env.GetEnv("DEBUG_LEVEL", configs.DEBUG_LEVEL)),
			Style: log.LoggerStyle(env.GetEnv("LOG_STYLE", configs.LOG_STYLE)),
		}

		if env.GetEnvAsBool("LOG_FILE_WRITE", configs.LOG_FILE_WRITE) {
			LoggerConfig.FileOption = &log.FileOptions{
				FileName:    env.GetEnv("POD_NAME", configs.SERVICE_NAME),
				MaxSizeMB:   env.GetEnvAsInt("LOG_MAX_SIZE_MB", configs.LOG_MAX_SIZE_MB),
				MaxBackups:  env.GetEnvAsInt("LOG_MAX_BACKUPS", configs.LOG_MAX_BACKUPS),
				MaxAgeDays:  env.GetEnvAsInt("LOG_MAX_AGE_DAYS", configs.LOG_MAX_AGE_DAYS),
				Compression: env.GetEnvAsBool("LOG_COMPRESSION", configs.LOG_COMPRESSION),
			}
		}
	})
	return LoggerConfig
}
