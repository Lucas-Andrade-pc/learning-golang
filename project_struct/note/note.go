package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

const nameFileNote = "note.json"

func (n Note) Save() error {
	json, _ := json.Marshal(n)
	err := os.WriteFile(nameFileNote, []byte(json), 0644)
	if err != nil {
		return errors.New("file write file")
	}
	return nil
}
func (n Note) Display() {
	fmt.Print(n.Title, n.Content, n.CreatedAt, "\n")
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title or content not value")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
