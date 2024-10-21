package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	file "importxml/read_file"
	"os"
	"strings"

	xj "github.com/basgys/goxml2json"
)

const caminho = "/home/lucas-irede/Downloads/output.xml"
const caminhojson = "test.json"

type HostProperties struct {
	//XMLName    xml.Name `xml:"robot"`
	Suite      []Tag  `xml:"suite"`
	Statistics []Tag2 `xml:"statistic"`
}

type Tag struct {
	Id     string   `xml:"id,attr"`
	Name   string   `xml:"name,attr"`
	Source string   `xml:"source,attr"`
	Status []string `xml:"status"`
	Test   []string
}

type Tag2 struct {
	Tag string `xml:"tag,attr"`
}

func SaveFile(name string, n string) error {

	//json, _ := json.Marshal(file)
	err := os.WriteFile(name, []byte(n), 0644)
	if err != nil {
		return errors.New("file write file")
	}
	return nil
}
func readFile() []byte {
	read, _ := file.ReadyOpenFile(caminhojson)
	json, _ := json.Marshal(read)
	return json
}

func main() {
	file, err := file.ReadyOpenFile(caminho)

	xml := strings.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	root := &xj.Node{}
	err = xj.NewDecoder(xml).Decode(root)
	if err != nil {
		panic(err)
	}
	// Then encode it in JSON
	var buf bytes.Buffer
	e := xj.NewEncoder(&buf)
	err = e.Encode(root)
	if err != nil {
		panic(err)
	}

	//fmt.Println(buf.String())
	SaveFile("test.json", buf.String())
	// // fmt.Println(file)
	var v HostProperties
	// err = xml.Unmarshal([]byte(file), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Print(v)
	// // fmt.Printf("v = %#v\n", v)
	// SaveFile("test.json", v)

	// // x := v.Suite
	// // fmt.Print(x)
}
