package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"strconv"
)

type User struct {
    Id    int    `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

var users = []User{
    {
        Id:    1,
        Name:  "Shuvo Chakrabrati",
        Age:   24,
        Email: "shuvo@gmail.com",
    },
    {
        Id:    2,
        Name:  "Habib Ullah",
        Age:   24,
        Email: "habib@gmail.com",
    },
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/health", healthHandler)
    mux.HandleFunc("POST /createUser", createUserHandler)
    mux.HandleFunc("GET /user", getUserHandler)
    // mux.HandleFunc("GET /user/{id}", getSingleUserHandler)
	mux.HandleFunc("GET /user/{id}", getSingleUserHandler)

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
    var newUser User

    err := json.NewDecoder(r.Body).Decode(&newUser)  // ← err capture করো
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Error decoding user", err)  // ← lowercase p
        return
    }

    newUser.Id = len(users) + 1
    users = append(users, newUser)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)  // ← created user টা return করো
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)  // ← সঠিক syntax, users (plural)
}


func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
    idParam := r.PathValue("id")
    fmt.Printf("the value of id is %v and the type is %T\n", idParam, idParam)

    id, err := strconv.Atoi(idParam)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "invalid user id")
        return
    }

    for _, user := range users {  // ← uer → user
        if user.Id == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(user)
            return
        }
    }

    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintln(w, "user not found")
}