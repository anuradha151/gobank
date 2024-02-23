package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(a *Account) error
	DeleteAccount(id int) error
	UpdateAccount(a *Account) error
	GetAccountById(id int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "user=root dbname=gobank password=root sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateAccount(a *Account) error {
	_, err := s.db.Exec("INSERT INTO accounts (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)", a.ID, a.FirstName, a.LastName, a.Email)
	return err
}