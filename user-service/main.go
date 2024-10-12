package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users []User
var nextID int

func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    _ = json.NewDecoder(r.Body).Decode(&newUser)
    newUser.ID = nextID
    nextID++
    users = append(users, newUser)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, user := range users {
        if user.ID == id {
            users = append(users[:index], users[index+1:]...)
            break
        }
    }
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    router := mux.NewRouter()

    // Routes
    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/users/{id}", getUser).Methods("GET")
    router.HandleFunc("/users", createUser).Methods("POST")
    router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

    // Seed some data
    users = append(users, User{ID: 1, Name: "Alice", Email: "alice@example.com"})
    users = append(users, User{ID: 2, Name: "Bob", Email: "bob@example.com"})
    nextID = 3

    log.Println("Server starting on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

