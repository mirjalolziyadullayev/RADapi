package handler

import (
	"fmt"
	"net/http"
	"time"
)


func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAccount(w,r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("\n'GET'-response sent to /account on", time.Now().Format(time.RFC850))
	case "POST":
		updateAccount(w,r)

		reqMessage := " 'POST'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"resuest": "%s"}`, reqMessage)
		fmt.Println("\n'POST'-response sent to /account on", time.Now().Format(time.RFC850))
	case "PUT":
		updateAccount(w,r)

		reqMessage := " 'POST'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'PUT'-response sent to /account on", time.Now().Format(time.RFC850))
	case "DELETE":
		deleteAccount(w,r)

		reqMessage := " 'POST'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'DELETE'-response sent to /account on", time.Now().Format(time.RFC850))
	}
}

func getAccount(w http.ResponseWriter, r *http.Request) {

}
func updateAccount(w http.ResponseWriter, r *http.Request) {

}
func deleteAccount(w http.ResponseWriter, r *http.Request) {

}