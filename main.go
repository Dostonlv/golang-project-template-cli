package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

type (
	errMsg error
)
type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Golang Project Name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}
func (m model) Init() tea.Cmd {
	return textinput.Blink
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	globalModel.Input = m.textInput.Value()
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"What’s your Project Name?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

type GlobalModel struct {
	Input    string
	Database string
}

var globalModel GlobalModel

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	// var input string
	// fmt.Print("Enter a project name: ")
	// fmt.Scan(&input)

	// Create project directory
	os.Mkdir(globalModel.Input, 0777)

	// Create Makefile
	os.Create(globalModel.Input + "/Makefile")

	// Create .gitignore
	os.WriteFile(globalModel.Input+"/.gitignore", []byte("/.env\n/.DS_Store\n/.idea"), 0777)

	// Create cmd directory
	os.Mkdir(globalModel.Input+"/cmd", 0777)

	// Create main.go in cmd directory
	os.WriteFile(globalModel.Input+"/cmd/main.go", []byte("package main"), 0777)

	// Create models directory
	os.Mkdir(globalModel.Input+"/models", 0777)

	// Create common.go in models directory
	os.WriteFile(globalModel.Input+"/models/common.go", []byte("package models\n\ntype DefaultError struct {\n\tMessage string `json:\"message\"`\n}\n\ntype ErrorResponse struct {\n\tMessage string `json:\"message\"`\n\tCode    int    `json:\"code\"`\n}\ntype SuccessResponse struct {\n\tMessage string      `json:\"message\"`\n\tData    interface{} `json:\"data\"`\n}"), 0777)

	// Create config directory
	os.Mkdir(globalModel.Input+"/config", 0777)

	// Create config.go in config directory
	os.WriteFile(globalModel.Input+"/config/config.go", []byte("package config"), 0777)

	// Create api directory
	os.Mkdir(globalModel.Input+"/api", 0777)

	// Create api.go in api directory
	os.WriteFile(globalModel.Input+"/api/api.go", []byte("package api"), 0777)

	// Create handlers directory inside api
	os.Mkdir(globalModel.Input+"/api/handlers", 0777)

	// Create handlers.go in handlers directory
	os.WriteFile(globalModel.Input+"/api/handlers/handlers.go", []byte("package handlers"), 0777)

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
	os.Mkdir(globalModel.Input+"/migrations", 0777)

	// Create database-specific migrations directory
	os.Mkdir(globalModel.Input+"/migrations/"+database, 0777)

	if database == "postgres" {
		// Create up and down migration scripts
		os.WriteFile(globalModel.Input+"/migrations/"+database+"/01_create_tables.down.sql", nil, 0777)
		os.WriteFile(globalModel.Input+"/migrations/"+database+"/01_create_tables.up.sql", nil, 0777)

	}

	// Create storage directory
	os.Mkdir(globalModel.Input+"/storage", 0777)

	// Create storage.go in storage directory
	os.WriteFile(globalModel.Input+"/storage/storage.go", []byte("package storage"), 0777)

	// Create database-specific storage directory
	os.Mkdir(globalModel.Input+"/storage/"+database, 0777)

	// Create database.go in database-specific storage directory
	os.WriteFile(globalModel.Input+"/storage/"+database+"/"+database+".go", []byte("package "+database), 0777)

	// golang version check with go from os !Feature
	// os.WriteFile(input+"/os.sh", []byte("#!/bin/bash \ngo version"), 0777)
	// _, err := exec.Command("/bin/sh", "./"+input+"/os.sh").Output()
	// if err != nil {
	// 	fmt.Printf("error %s", err)
	// }
	fmt.Println("Project structure created successfully.")
}
