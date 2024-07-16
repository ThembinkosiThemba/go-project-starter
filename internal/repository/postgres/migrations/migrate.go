package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ThembinkosiThemba/go-project-starter/internal/repository/postgres"
	custom_response "github.com/ThembinkosiThemba/go-project-starter/pkg/http"
	"github.com/gin-gonic/gin"
)

func MigrateEndPoint(c *gin.Context) {
	db, err := postgres.PostgresConn()
	if err != nil {
		return
	}

	err = Migrate(db)
	if err != nil {
		return
	}

	db.Close()

	custom_response.WriteJSON(c, http.StatusOK, 1, nil, "migrations ran successfully")
}

func Migrate(db *sql.DB) error {
	path := "internal/repository/postgres/migrations"

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(path, file.Name())
			query, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read migration script %s: %v", file.Name(), err)
			}

			sqlStmts := strings.Split(string(query), ";")
			for _, stmt := range sqlStmts {
				trimmedStmt := strings.TrimSpace(stmt)
				if len(trimmedStmt) > 0 {
					if _, err := db.Exec(trimmedStmt); err != nil {
						return fmt.Errorf("failed to apply migration script %s: %v", file.Name(), err)
					}
				}
			}

			log.Printf("Applied migration: %s\n", file.Name())
		}
	}

	log.Println("All migrations applied successfully")
	return nil
}
