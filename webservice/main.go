package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db = SetupDB()

type User struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	http.HandleFunc("/users", getUsers)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT email, phone_number FROM users")

	PanicIf(err)
	defer rows.Close()

	var users []User
	var email, phone_number string
	for rows.Next() {
		err := rows.Scan(&email, &phone_number)
		PanicIf(err)
		users = append(users, User{email, phone_number})
	}

	res := struct {
		Users  []User
		Errors []string
	}{
		users,
		[]string{""},
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(res)

	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	connection := fmt.Sprintf("root:123@tcp(%s:3306)/users", os.Getenv("DATABASE_HOST"))
	db, err := sql.Open("mysql", connection)
	PanicIf(err)

	return db
}
