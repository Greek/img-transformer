package env

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/greek/img-transform/internal/lib/logging"
	"github.com/joho/godotenv"
)

type EnvData struct {
	S3_ACCESS_KEY string `env:"S3_ACCESS_KEY"`
	S3_SECRET_KEY string `env:"S3_SECRET_KEY"`
	S3_REGION     string `env:"S3_REGION"`
}

var cfg EnvData

// GetEnv loads environment variables from a .env file
func GetEnv() *EnvData {
	log := logging.BuildLogger("LoadEnv")

	err := godotenv.Load()
	if err != nil {
		log.Warn("env file does not exist")
	}

	cfg, err := env.ParseAs[EnvData]()

	return &cfg
}

// CheckEnv checks if the proper environment variables are present
func CheckEnv() {
	log := logging.BuildLogger("CheckEnv")

	log.Info("checking env variables")

	cfg := GetEnv()
	if len(cfg.S3_ACCESS_KEY) == 0 {
		log.Error("S3_ACCESS_KEY is not defined, exiting")
		os.Exit(1)
	}

	if len(cfg.S3_SECRET_KEY) == 0 {
		log.Error("S3_SECRET_KEY is not defined, exiting")
		os.Exit(1)
	}

	if len(cfg.S3_REGION) == 0 {
		log.Error("S3_REGION is not defined, exiting")
		os.Exit(1)
	}
}
