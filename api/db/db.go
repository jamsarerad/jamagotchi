package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "jamagotchi/models"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("postgres", "your_connection_string_here")
    if err != nil {
        panic(err)
    }
}

func RegisterUser(user models.User) {
    query := `INSERT INTO users (username, password) VALUES ($1, $2)`
    _, err := DB.Exec(query, user.Username, user.Password)
    if err != nil {
        // Handle error
    }
}


