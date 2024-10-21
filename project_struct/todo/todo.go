package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

const nameFileTodo = "todo.json"

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (t Todo) Save() error {

	json, _ := json.Marshal(t)
	err := os.WriteFile(nameFileTodo, []byte(json), 0644)
	if err != nil {
		return errors.New("file write file")
	}
	return nil
}
func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("todo not valid")
	}
	return Todo{
		Text: content,
	}, nil
}
