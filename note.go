package main

import (
	"fmt"
	"time"
)

// Struct

type note struct {
	noteTitle   string
	noteContent string
	timeStamp   time.Time
}

func new(title string, content string) note {
	return note{
		noteTitle:   title,
		noteContent: content,
		timeStamp:   time.Now(),
	}
}

func (note note) getNoteTitle() string {
	return "dummy text"
}

func (note note) getNoteContent() string {
	return "dummy text"
}

// Utilities

func promptUser(prompt string) string {
	fmt.Println(prompt)
	var userData string
	fmt.Scan(&userData)
	return userData
}

// Main logic

func main() {
	var noteTitle string = promptUser("Please enter a title for your note:")
	var noteContent string = promptUser("Now please enter your note:")

	var newNote note = new(noteTitle, noteContent)

	fmt.Println(newNote.getNoteTitle())
	fmt.Println(newNote.getNoteContent())
}
