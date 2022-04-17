package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	urlExample := "postgres://postgres:psql@localhost:5432/trade_app_2"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	err = conn.QueryRow(context.Background(), "select status from order_data where order_id=$1", 58).Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	var name1 string
	err = conn.QueryRow(context.Background(), "select status from order_data where order_id=$1", 44).Scan(&name1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	var name2 string
	err = conn.QueryRow(context.Background(), "select status from order_data where order_id=$1", 23).Scan(&name2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	var name3 string
	err = conn.QueryRow(context.Background(), "select status from order_data where order_id=$1", 33).Scan(&name3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Status", name)
	fmt.Println("Status", name1)
	fmt.Println("Status", name2)
	fmt.Println("Status", name3)
}
