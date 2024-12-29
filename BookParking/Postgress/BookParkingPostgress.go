package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB() {
	var err error
	dsn := "postgres://admin:admin@localhost:5432/BookParking"
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	log.Println("Connected to the database.")
}

func CloseDB() {
	db.Close()
}

func CreateUser(name, email string) (int, error) {
	var id int
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(context.Background(), query, name, email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUserByID(id int) (string, string, error) {
	var name, email string
	query := `SELECT name, email FROM users WHERE id = $1`
	err := db.QueryRow(context.Background(), query, id).Scan(&name, &email)
	if err != nil {
		return "", "", err
	}
	return name, email, nil
}

func UpdateUser(id int, name, email string) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := db.Exec(context.Background(), query, name, email, id)
	return err
}

func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(context.Background(), query, id)
	return err
}

func main() {
	InitDB()
	defer CloseDB()

	// Create a user
	userID, err := CreateUser("John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("User created with ID: %d\n", userID)

	// Read user details
	name, email, err := GetUserByID(userID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	fmt.Printf("Fetched User - Name: %s, Email: %s\n", name, email)

	// Update user
	err = UpdateUser(userID, "John Updated", "john.updated@example.com")
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("User updated successfully.")

	/*
		// Delete user
		err = DeleteUser(userID)
		if err != nil {
			log.Fatalf("Error deleting user: %v", err)
		}
		fmt.Println("User deleted successfully.")
	*/
}
