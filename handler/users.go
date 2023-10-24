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
		fmt.Println("\n'POST'-response sent on", time.Now().Format(time.RFC850))
	case "DELETE":
		deleteUser(w,r)

		reqMessage := " 'DELETE'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'DELETE'-response sent on", time.Now().Format(time.RFC850))
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
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser models.User
	json.NewDecoder(r.Body).Decode(&deleteUser)

	var userData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userData)

	for i := 0; i < len(userData); i++ {
		if userData[i].EmailUsername == deleteUser.EmailUsername && userData[i].Password == deleteUser.Password {
			fmt.Println("\n_____________________________________________")
			fmt.Println("Found and deleted user at", time.Now().Format(time.RFC850))
			fmt.Println("_____________________________________________")
			fmt.Println("____________________________")
			fmt.Println("User ID was: ", userData[i].Id)
			fmt.Println("____________________________")
			fmt.Println("First name was: ", userData[i].Firstname)
			fmt.Println("____________________________")
			fmt.Println("Last name was: ", userData[i].Lastname)
			fmt.Println("____________________________")
			fmt.Println("Email or username was: ", userData[i].EmailUsername)
			fmt.Println("____________________________")
			fmt.Println("Password was: ", userData[i].Password)
			fmt.Println("____________________________")
			userData = append(userData[:i], userData[i+1:]... )
		}
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/users.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Println("--------------------------------")
	fmt.Println("Operation Completed")
	fmt.Println("--------------------------------")
}