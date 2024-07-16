package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PostgresConn() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// remember to set the env variable on your file
	dbUrl := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
		return nil, err
	}

	logger.Info("connected to postgres")

	return db, nil
}
