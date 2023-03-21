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
	}
	defer conn.Close(context.Background())

	tag, err := conn.Exec(context.Background(), "DROP TABLE messages")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(tag)

	tag, err = conn.Exec(context.Background(), "DROP TABLE users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
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
	}
	fmt.Println(tag)

	tag, err = conn.Exec(context.Background(), `CREATE TABLE users (
		user_id BIGSERIAL PRIMARY KEY,
		user_name text NOT NULL,
		user_role text NOT NULL,
		email text NOT NULL,
		password text NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	  )`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(tag)

}
