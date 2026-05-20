package main

import (
	    "encoding/json"  
	"fmt"
	"net/http"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

	var user = []User{
		{
			Id: 1,
			Name: "Shuvo Chakrabrati",
			Age: 24,
			Email: "shuvo@gmail.com",
		},
		{
			Id: 2,
			Name: "Habib Ullah",
			Age: 24,
			Email: "habib@gmail.com",
		},
	}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /createUser", createUserHandler)
	mux.HandleFunc("GET /user", getUserHandler)

	fmt.Println("Server is running at port 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Server error", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up and running!")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.	StatusMethodNotAllowed)
	// 	fmt.Fprintln(w, "Method not allowed")
	// 	return
	// }
	// fmt.Fprintln(w, "User created successfully!")

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.FPrintln(w, "Error decoding user", err)
		return
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
// users, _:=	json.Marshal(user)
// w.Write(users)

json.NewEncoder(w).encoder.Encode(user)

}