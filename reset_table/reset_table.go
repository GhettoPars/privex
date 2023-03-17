package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db = make(map[string]string)

func main() {
	conn_string := os.Getenv("DB_CONN")
	if conn_string == "" {
		panic("No db connection string")
	}

	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	tag, err := conn.Exec(context.Background(), "DROP TABLE messages")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(tag)

	tag, err = conn.Exec(context.Background(), `CREATE TABLE messages (
		message_id   BIGSERIAL PRIMARY KEY,
		user_id integer NOT NULL,
		message_text text NOT NULL,
		message_type text NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	  )`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(tag)

	// var greeting string
	// err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(greeting)
}
