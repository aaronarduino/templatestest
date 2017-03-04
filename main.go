package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

type MainTemplate struct {
	b bytes.Buffer
}

func (m *MainTemplate) exec(baseFile string, data PageData) {
	t := template.New(baseFile)

	t, err := t.ParseFiles("templates/" + baseFile)
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(&m.b, data)
	if err != nil {
		log.Println(err)
	}
}

func (m *MainTemplate) writeToBytes() []byte {
	return m.b.Bytes()
}

func (m *MainTemplate) writeToString() string {
	return m.b.String()
}

type PageData struct {
	Message string
	Title   string
}

func main() {
	data := PageData{Message: "Hello", Title: "title"}

	mainTemp := MainTemplate{}
	mainTemp.exec("index.html", data)
	out := mainTemp.writeToBytes()

	err := ioutil.WriteFile("index.html", out, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done.")
	fmt.Println()
	fmt.Println(mainTemp.writeToString())
}
