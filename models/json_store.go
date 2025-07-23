package models

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	usersFile     = "data/users.json"
	superiorsFile = "data/superiors.json"
	users         []User
	superiors     []Superior
	mu            sync.Mutex
)

func init() {
	loadUsers()
	loadSuperiors()
}

func loadUsers() {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(usersFile)
	if err != nil {
		// If file doesn't exist, initialize with empty slice
		users = []User{}
		return
	}
	json.Unmarshal(data, &users)
}

func saveUsers() {
	data, _ := json.MarshalIndent(users, "", "  ")
	os.WriteFile(usersFile, data, 0644)
}

func loadSuperiors() {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(superiorsFile)
	if err != nil {
		// If file doesn't exist, initialize with empty slice
		superiors = []Superior{}
		return
	}
	json.Unmarshal(data, &superiors)
}

func saveSuperiors() {
	data, _ := json.MarshalIndent(superiors, "", "  ")
	os.WriteFile(superiorsFile, data, 0644)
}

// Get all users
func GetAllUsers() []User {
	mu.Lock()
	defer mu.Unlock()
	return users
}

// Add a new user
func AddUser(user User) {
	mu.Lock()
	defer mu.Unlock()
	users = append(users, user)
	saveUsers()
}

// Get all superiors
func GetAllSuperiors() []Superior {
	mu.Lock()
	defer mu.Unlock()
	return superiors
}

// Add a new superior
func AddSuperior(superior Superior) {
	mu.Lock()
	defer mu.Unlock()
	superiors = append(superiors, superior)
	saveSuperiors()
}

// Clear all users (for simplified updates)
func ClearUsers() {
	users = []User{}
	saveUsers()
}

// Clear all superiors (for simplified updates)
func ClearSuperiors() {
	superiors = []Superior{}
	saveSuperiors()
}
