package main

import (
	"fmt"
	"net/http"
)

type User struct{
	Id int
	Name string
	Age int
	Email string
}

	var user = []User{
		{
			id: 1,
			name: "Shuvo Chakrabrati",
			age: 24,
			email: "shuvo@gmail.com"
		},
		{
			id: 2,
			name: "Habib Ullah",
			age: 24,
			email: "habib@gmail.com"
		}
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
	fmt.Fprintln(w, "User created successfully!")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
users, _:=	json.Marshal(user)
w.Write(users)
}