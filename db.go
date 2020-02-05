package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type JOB struct {
	ID        int
	JobId     int
	TypeID    int
	Text      string
	CreatedAt string
	UpdatedAt string
}

func db(dbUser string, dbPass string, hostName string) {
	dns := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=center sslmode=disable connect_timeout=5", hostName, dbUser, dbPass)
	db, err := sql.Open("postgres", dns)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	log.Println("result: ===========")

	if err != nil {
		panic(err)
	}

	var text string
	row := db.QueryRowContext(ctx, "select text from writing_job limit 1")
	err = row.Scan(&text)
	log.Println(text)

	err = db.PingContext(ctx)
	if err != nil {
		log.Println("==== timeout expired ====")
		log.Println(err)
	}

}
