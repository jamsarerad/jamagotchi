package routes

import (
	"jamagotchi/handlers"
	"jamagotchi/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// User-related endpoints
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Game-related endpoints wrapped with AuthMiddleware
	game := r.PathPrefix("/api/game").Subrouter()
	game.Use(middlewares.AuthMiddleware)

	game.HandleFunc("/roll", handlers.RollForItems).Methods("POST")
	game.HandleFunc("/user", handlers.GetUserDetails).Methods("GET")
	game.HandleFunc("/purchase", handlers.PurchaseItems).Methods("POST")
	game.HandleFunc("/inventory", handlers.ListInventory).Methods("GET")
	game.HandleFunc("/equip", handlers.EquipAccessory).Methods("POST")

	return r
}

