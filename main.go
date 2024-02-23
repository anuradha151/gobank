package main

import "fmt"

func main() {
	store := NewPostgresStorage()

	fmt.Println(store)

	server := NewAPIServer(":8080", store)
	server.Run()

}
