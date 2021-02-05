package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// Content for body of txt file
type Page struct {
	TextFilePath string
	TextFileName string
	HTMlPagePath string
	Body         string
}

func createFromTextFile(filepath string) Page {

	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	removedExetension := strings.Split(filepath, ".txt")[0]

	return Page{
		TextFilePath: filepath,
		TextFileName: removedExetension,
		HTMlPagePath: removedExetension + ".html",
		Body:         string(fileContents),
	}

}

func renderTemplateFromPage(templateFilePath string, page Page) {

	t := template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	w, err := os.Create(page.HTMlPagePath)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, templateFilePath, page)
	if err != nil {
		panic(err)
	}

}

func main() {

	var textFilePath string
	flag.StringVar(&textFilePath, "file", "", "Name or path to a text file")
	flag.Parse()

	if textFilePath == "" {
		panic("Missing the --file flag! Please supply one.")
	}

	newPage := createFromTextFile(textFilePath)
	renderTemplateFromPage("template.tmpl", newPage)

}
