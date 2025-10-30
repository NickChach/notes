package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NickChach/notes/note"
)

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

	newNote := note.New()
	newNote.SetNoteTitle(noteTitle)
	newNote.SetNoteContent(noteContent)
	newNote.SetNoteTimestamp()

	fmt.Println(newNote.GetNoteTitle())
	fmt.Println(newNote.GetNoteContent())

	newNote.Save()
}
