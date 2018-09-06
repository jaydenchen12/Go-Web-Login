package main;

import (
  "fmt"
  "net/http"
  "encoding/json"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/lib/pq"
)

type Authorizations struct {
  Username string `json:"username", db:"username"`
  Password string `json:"password", db:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request){
  creds := &Authorizations{}
  err := json.NewDecoder(r.Body).Decode(creds)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  hashedPass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 5)
  if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPass));
    err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
}

func Login(w http.ResponseWriter, r *http.Request){
  creds := &Authorizations{}
  err := json.NewDecoder(r.Body).Decode(creds)
  fmt.Println("Server is listening on port 8080")
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  password := db.QueryRow("select password from users where username=$1", creds.Username)
  if err != nil{
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  storedAuthorizations := &Authorizations{}
  err = password.Scan(&storedAuthorizations.Password)
  if err != nil {
    // No user found
    if err == sql.ErrNoRows {
      w.WriteHeader(http.StatusUnauthorized)
      return
    }

    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if err = bcrypt.CompareHashAndPassword([]byte(storedAuthorizations.Password), []byte(creds.Password));
  err != nil {
    w.WriteHeader(http.StatusUnauthorized)
  }
}

func Default(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hi there, Welcome to the Go web login")
}
