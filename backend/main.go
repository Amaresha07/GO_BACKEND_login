package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var connection *sql.DB

func main() {
	// Connect to the database
	var err error
	connection, err = sql.Open("postgres", "postgresql://demo_du6c_user:1f2qZEeQmeYIpBR5EOI5G2m4uoPDIdC6@dpg-ctls88aj1k6c73d2gti0-a.oregon-postgres.render.com/demo_du6c")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer connection.Close()

	// Define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)

	// Wrap the mux with CORS middleware
	handler := corsMiddleware(mux)

	// Start the server
	log.Println("Server is running on port 8000...")
	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatalf("Unable to start the server: %v", err)
	}
}

type User struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var loginUser User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check user in the database
	var userID string
	err := connection.QueryRow("SELECT id FROM users WHERE email=$1 AND password=$2", loginUser.Email, loginUser.Password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
			log.Println("Database error:", err)
		}
		return
	}

	// Respond with a success message
	response := map[string]string{
		"message": "Already registered",
		"user_id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func register(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	_, err := connection.Exec("INSERT INTO users(id, email, password) VALUES($1, $2, $3)",
		uuid.NewString(), newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		log.Println("Database error:", err)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
