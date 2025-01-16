package main

import (
	"encoding/json"
	"httpServer/db"
	"log"
  "net/http"

	"github.com/gorilla/mux"
)

func fetchUsers(w http.ResponseWriter, r *http.Request) {

	var users []db.User          // Assuming your User struct is in db package
	result := db.DB.Find(&users) // Fetch users into the users slice

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users) // Encode and send users as JSON
}


func main() {
	db.InitDB()

	// Insert a new user
	// newUser := db.User{Name: "heyyy Doe", Email: "john@example.com", Age: 30}
	// result := db.DB.Create(&newUser) // Insert data into the users table
	// if result.Error != nil {
	// 	fmt.Println("Error inserting user:", result.Error)
	// } else {
	// 	fmt.Println("User created successfully:", newUser)
	// }

	router := mux.NewRouter()

	// Define routes before starting the server
	router.HandleFunc("/users", fetchUsers).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Getting the / route")
		w.Write([]byte("Welcome to the API"))
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
