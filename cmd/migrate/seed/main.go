package main

import (
	"log"

	"github.com/Intro0/goSocial/internal/db"
	"github.com/Intro0/goSocial/internal/env"
	"github.com/Intro0/goSocial/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
