package handlers

import (
    "net/http"
    "encoding/json"
    "models"
    "db"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User // Assume User is a struct you've defined
    _ = json.NewDecoder(r.Body).Decode(&user)
    
    // Register the user in the database
    db.RegisterUser(user)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

