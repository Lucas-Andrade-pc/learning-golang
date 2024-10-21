package main

import (
	"fmt"
	"project/note"
	"project/todo"
)

type saver interface {
	Save() error
}

type outputdisplay interface {
	saver
	Display()
}

func printSomethings(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Value - ", value)
	case float64:
		fmt.Println("Value - ", value)
	case string:
		fmt.Println("Value - ", value)
	}
}

func outputAndSave(data outputdisplay) {
	err := saveData(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Display()
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("error save file")
		return err
	}
	return nil
}

func main() {
	printSomethings(1)
	printSomethings("teste")
	printSomethings(1.5)
	file, err := getInfoNote("Info Title:")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := getInfoNote("Info Content:")
	if err != nil {
		fmt.Println(err)
		return
	}
	getTodo, err := getInfoNote("Info Todo:")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataNote, err := note.New(file, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataTodo, err := todo.New(getTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputAndSave(dataNote)
	outputAndSave(dataTodo)
}
