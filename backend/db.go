package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file:", err)
		return nil, err
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Print("Failed to connect to the database:", err)
		return nil, err
	}
	log.Println("Successfully connected to psqlDB")
	return db, nil
}
func InitTables(DB *sql.DB) error {
	queries := []string{
		`create table if not exists users (
			id serial primary key,
			name varchar(16) not null,
			email varchar(40) unique not null,
			password varchar(16) not null,
			created_at timestamp default current_timestamp
		)`,
		`create table if not exists folders (
			id serial primary key,
			name varchar(50) not null,
			user_id int references users(id),
			created_at timestamp default current_timestamp
		)`,
		`create table if not exists notes (
			id serial primary key,
			title varchar(100) not null,
			content text,
			user_id int references users(id),
			folder_id int references folders(id) default null,
			pinned boolean not null default false,
			featured boolean not null default false,
			created_at timestamp default current_timestamp
		)`,
	}
	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
func DeleteTables(DB *sql.DB) error {
	queries := []string{
		`drop table if exists notes;`,
		`drop table if exists folders;`,
		`drop table if exists users;`,
	}
	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
