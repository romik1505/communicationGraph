package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/romik1505/communicationGraph/internal/app/store"
)

// NewPostgresConnection Подключение к бд
func NewPostgresConnection() store.Storage {
	pgCon := os.Getenv("PGDSN")
	log.Println("pgcon: ", pgCon)
	conn, err := sql.Open("postgres", pgCon)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return store.Storage{
		DB: conn,
	}
}
