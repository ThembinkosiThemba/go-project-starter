package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres/migrations"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PostgresConn() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	// remember to set the env variable on your file
	dbUrl := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
		return nil
	}

	if err = migrations.Migrate(db); err != nil {
		log.Println("failed to perform migrations", err)
	}
	log.Println("connected to postgres")

	return db
}
