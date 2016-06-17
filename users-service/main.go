package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db = SetupDB()

func main() {
	http.HandleFunc("/users", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT email, phone_number, description FROM books")

	PanicIf(err)
	defer rows.Close()

	var email, phone_number string
	for rows.Next() {
		err := rows.Scan(&email, &phone_number)
		PanicIf(err)
		fmt.Printf("Email: %s\nPhone: %s\n\n", email, phone_number)
	}

	io.WriteString(w, "hello, world!\n")
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// dials the database, returning any error
func SetupDB() *sql.DB {
	// connection := fmt.Sprintf("user_service:123@%s@db/users", os.Getenv("DATABASE_HOST"))
	// db, err := sql.Open("mysql", connection)
	db, err := sql.Open("mysql", "root:123@db/users")
	PanicIf(err)

	return db
}