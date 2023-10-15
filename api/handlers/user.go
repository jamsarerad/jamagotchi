package handlers

import (
  "net/http"
  "encoding/json"
  "jamagotchi/models"
  "jamagotchi/db"
  "github.com/golang-jwt/jwt"
)

func Register(w http.ResponseWriter, r *http.Request) {
  var user models.User // Assume User is a struct you've defined
  _ = json.NewDecoder(r.Body).Decode(&user)

  // Connect to database
  database := db.ConnectToUsers()

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})

  // Parse username and password from request body
  var credentials struct {
    Username     string `json:"username"`
    Password     string `json:"password"`
    Email        string `json:"email"`
    LeetCodeUser string `json:"leetcode_username"`
  }
  _ = json.NewDecoder(r.Body).Decode(&credentials)

  // Register the user in the database
  _, err := db.RegisterUser(database, credentials.Username, credentials.Email, credentials.Password, credentials.LeetCodeUser)
  if err != nil {
    http.Error(w, "Failed to register", http.StatusServiceUnavailable)
  }

  // Create a new token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": credentials.Username,
    })

  // Sign and get the complete encoded token as a string
  tokenString, _ := token.SignedString([]byte("your_secret_key"))

  // Return the token
  w.Header().Set("Authorization", "Bearer "+tokenString)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {

}
