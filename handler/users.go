package handler

import (
	"RADserver/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllUsers(w,r)	

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("\n'GET'-response sent on", time.Now().Format(time.RFC850))
	case "POST":
		createUser(w,r)

		reqMessage := " 'POST'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)

	case "DELETE":
		// deleteUser(w,r)

	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var usersData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &usersData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersData)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	var userData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userData)


	newUser.Id = len(userData)+1
	userData = append(userData, newUser)

	res, _ := json.Marshal(userData)
	os.WriteFile("db/users.json", res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Created new user at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("User ID is: ", newUser.Id)
	fmt.Println("____________________________")
	fmt.Println("First name is: ", newUser.Firstname)
	fmt.Println("____________________________")
	fmt.Println("Last name is: ", newUser.Lastname)
	fmt.Println("____________________________")
	fmt.Println("Email or username is: ", newUser.EmailUsername)
	fmt.Println("____________________________")
	fmt.Println("Password is: ", newUser.Password)
	fmt.Println("____________________________")
} 