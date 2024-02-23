package main

import "math/rand"

type Account struct {
	ID   int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`

}

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		ID:   rand.Intn(1000000),
		FirstName: firstName,
		LastName: lastName,
		Email: email,
	}
}