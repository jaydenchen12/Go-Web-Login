package main;

import (
  "fmt"
  "net/http"
  "database/sql"
  "log"
  _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
  http.HandleFunc("/login", Login)
  http.HandleFunc("/signup", Signup)
  http.HandleFunc("/", Default)

  initDB()
  fmt.Println("Server is listening on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB(){
  var err error
  db, err = sql.Open("postgres", "dbname=login_creds sslmode=disable")
  if err != nil {
    panic(err)
  }
}
