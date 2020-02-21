package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Repo struct {
	*sql.DB
}

func NewRepo(host, port, database, username, password string) *Repo {
	conStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s", host, port, database, username, password)
	db, err := sql.Open("postgresql", conStr)
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{db}
}
