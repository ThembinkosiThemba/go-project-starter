package mysql

import (
	"database/sql"
	"log"
	"os"

	"github.com/ThembinkosiThemba/go-project-starter/pkg/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

func MySqlConn() (*sql.DB, error) {
	// remember to set the env variable on your file
	dbUrl := os.Getenv("MY_SQL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
		return nil, err
	}

	logger.Info("connected to mysql")

	return db, nil
}
