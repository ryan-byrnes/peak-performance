package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
)

func main() {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	var password string = os.Getenv("DB_PASS")
	var dbname string = os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}