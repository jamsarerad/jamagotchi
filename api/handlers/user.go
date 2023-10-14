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

  // Register the user in the database
  db.RegisterUser(user)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})

  // Parse username and password from request body
  var credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }
  _ = json.NewDecoder(r.Body).Decode(&credentials)

  // Check username and password against database (in this case, a map)
  storedPassword, ok := users[credentials.Username]
  if !ok {
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
  }

  // Compare hashed passwords
  if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password)); err != nil {
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
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
