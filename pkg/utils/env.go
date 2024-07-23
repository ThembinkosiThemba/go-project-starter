package utils

import (
	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Warn("Please make sure you set environment file")
	}
}
