package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct

type note struct {
	NoteTitle   string    `json:"title"`
	NoteContent string    `json:"content"`
	TimeStamp   time.Time `json:"timestamp"`
}

func new(title string, content string) note {
	return note{
		NoteTitle:   title,
		NoteContent: content,
		TimeStamp:   time.Now(),
	}
}

func (note note) getNoteTitle() string {
	return note.NoteTitle
}

func (note note) getNoteContent() string {
	return note.NoteContent
}

func (note note) save() {
	var fileName string = strings.ReplaceAll(note.NoteTitle, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fmt.Sprintf("%v.json", fileName)

	json, _ := json.Marshal(note)
	os.WriteFile(fileName, json, 0644)
}

// Utilities

func promptUser(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var userData string = scanner.Text()
	return userData
}

// Main logic

func main() {
	var noteTitle string = promptUser("Please enter a title for your note:")
	var noteContent string = promptUser("Now please enter your note:")

	var newNote note = new(noteTitle, noteContent)

	fmt.Println(newNote.getNoteTitle())
	fmt.Println(newNote.getNoteContent())

	newNote.save()
}
