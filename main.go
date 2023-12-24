package main

import (
	"fmt"

	"os"
)

func main() {

	var input string
	fmt.Print("Enter a project name: ")
	fmt.Scan(&input)

	// Create project directory
	os.Mkdir(input, 0777)

	// Create Makefile
	os.Create(input + "/Makefile")

	// Create .gitignore
	os.WriteFile(input+"/.gitignore", []byte("/.env\n/.DS_Store\n/.idea"), 0777)

	// Create cmd directory
	os.Mkdir(input+"/cmd", 0777)

	// Create main.go in cmd directory
	os.WriteFile(input+"/cmd/main.go", []byte("package main"), 0777)

	// Create models directory
	os.Mkdir(input+"/models", 0777)

	// Create common.go in models directory
	os.WriteFile(input+"/models/common.go", []byte("package models\n\ntype DefaultError struct {\n\tMessage string `json:\"message\"`\n}\n\ntype ErrorResponse struct {\n\tMessage string `json:\"message\"`\n\tCode    int    `json:\"code\"`\n}\ntype SuccessResponse struct {\n\tMessage string      `json:\"message\"`\n\tData    interface{} `json:\"data\"`\n}"), 0777)

	// Create config directory
	os.Mkdir(input+"/config", 0777)

	// Create config.go in config directory
	os.WriteFile(input+"/config/config.go", []byte("package config"), 0777)

	// Create api directory
	os.Mkdir(input+"/api", 0777)

	// Create api.go in api directory
	os.WriteFile(input+"/api/api.go", []byte("package api"), 0777)

	// Create handlers directory inside api
	os.Mkdir(input+"/api/handlers", 0777)

	// Create handlers.go in handlers directory
	os.WriteFile(input+"/api/handlers/handlers.go", []byte("package handlers"), 0777)

	var database string
	fmt.Print("Enter a database name (1 for PostgreSQL, 2 for MongoDB): ")
	fmt.Scan(&database)
	if database == "1" {
		database = "postgres"
	}
	if database == "2" {
		database = "mongodb"
	}

	// Create migrations directory
	os.Mkdir(input+"/migrations", 0777)

	// Create database-specific migrations directory
	os.Mkdir(input+"/migrations/"+database, 0777)

	if database == "postgres" {
		// Create up and down migration scripts
		os.WriteFile(input+"/migrations/"+database+"/01_create_tables.down.sql", nil, 0777)
		os.WriteFile(input+"/migrations/"+database+"/01_create_tables.up.sql", nil, 0777)

	}

	// Create storage directory
	os.Mkdir(input+"/storage", 0777)

	// Create storage.go in storage directory
	os.WriteFile(input+"/storage/storage.go", []byte("package storage"), 0777)

	// Create database-specific storage directory
	os.Mkdir(input+"/storage/"+database, 0777)

	// Create database.go in database-specific storage directory
	os.WriteFile(input+"/storage/"+database+"/"+database+".go", []byte("package "+database), 0777)

	// golang version check with go from os !Feature
	// os.WriteFile(input+"/os.sh", []byte("#!/bin/bash \ngo version"), 0777)
	// _, err := exec.Command("/bin/sh", "./"+input+"/os.sh").Output()
	// if err != nil {
	// 	fmt.Printf("error %s", err)
	// }
	fmt.Println("Project structure created successfully.")
}
