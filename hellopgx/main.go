package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Hello struct {
	Id   int
	Nama string
}

func hello(hello Hello) string {
	return fmt.Sprint("id: ", hello.Id, " name: ", hello.Nama)
}

func FetchDataList(conn *pgx.Conn, limit int) ([]Hello, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, nama FROM hello LIMIT $1", limit)
	if err != nil {
		return nil, err
	}

	resultList := []Hello{}

	for rows.Next() {
		var hello = Hello{}
		var err = rows.Scan(&hello.Id, &hello.Nama)
		if err != nil {
			return nil, err
		}

		resultList = append(resultList, hello)
	}

	return resultList, nil
}

func main() {
	fmt.Println("Hello PGX")

	urlExample := "postgres://sts:Awesome123!@localhost:14555/hellopgx"

	conn, err := pgx.Connect(context.Background(), urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	fmt.Println("Database connection success !")

	list, errDb := FetchDataList(conn, 2)
	if errDb != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch data: %v\n", errDb)

		os.Exit(1)
	}

	fmt.Println("Get content hello : ")
	for idx, hello := range list {
		fmt.Println("hello[", idx, "] => ", hello)
	}
}
