package file

import (
	"encoding/xml"
	"errors"
	"os"
)

type Robot struct {
	XMLName xml.Name `xml:"robot"`
	Tags    []Tag    `xml:"kw"`
}

type Tag struct {
	Key   string `xml:"test"`
	Value string `xml:"kw"`
}

func ReadyOpenFile(dataFile string) (string, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return "", errors.New("failed open file")
	}
	return string(data), nil

}
