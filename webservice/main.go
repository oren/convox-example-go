package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	User         string
	Password     string
	DatabaseName string
}

type AppConfig struct {
	Database DBConfig
}

var (
	ConfigFile = flag.String("config", "config.json", "Config file to load")
	Config     AppConfig
	db         *sql.DB
)

func init() {
	flag.Parse()

	ConfigBytes, err := ioutil.ReadFile(*ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(ConfigBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}

	db = SetupDB()
}

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
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", Config.Database.User, Config.Database.Password, os.Getenv("DATABASE_HOST"), Config.Database.DatabaseName)
	db, err := sql.Open("mysql", connection)
	PanicIf(err)

	return db
}
