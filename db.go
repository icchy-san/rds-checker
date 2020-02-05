package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
	dns := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=center sslmode=disable", hostName, dbUser, dbPass)
	db, err := sql.Open("postgres", dns)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select text from writing_job limit 3")
	var jobs []JOB
	for rows.Next() {
		var job JOB
		rows.Scan(&job.Text)
		jobs = append(jobs, job)
	}

	for _, b := range jobs {
		log.Println(b.Text)
	}
}
