package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	file, err := os.Open("../.env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", string(cBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
