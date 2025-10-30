package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	noteTitle     string
	noteContent   string
	noteTimestamp time.Time
}

func New() *Note {
	return &Note{}
}

func (note *Note) SetNoteTitle(title string) {
	note.noteTitle = title
}

func (note *Note) SetNoteContent(content string) {
	note.noteContent = content
}

func (note *Note) SetNoteTimestamp() {
	note.noteTimestamp = time.Now()
}

func (note Note) GetNoteTitle() string {
	return note.noteTitle
}

func (note Note) GetNoteContent() string {
	return note.noteContent
}

func (note Note) GetNoteTimestamp() time.Time {
	return note.noteTimestamp
}

func (note Note) marshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"title":     note.GetNoteTitle(),
		"content":   note.GetNoteContent(),
		"timestamp": note.GetNoteTimestamp(),
	}

	return json.Marshal(data)
}

func (note Note) Save() {
	var fileName string = strings.ReplaceAll(note.GetNoteTitle(), " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fmt.Sprintf("%v.json", fileName)

	json, _ := note.marshalJSON()
	os.WriteFile(fileName, json, 0644)
}
