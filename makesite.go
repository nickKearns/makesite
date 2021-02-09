package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Page is the struct of a page
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Body         string
}

func createFromTextFile(filepath string) Page {

	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	removedExtension := strings.Split(filepath, ".txt")[0]

	return Page{
		TextFilePath: filepath,
		TextFileName: removedExtension,
		HTMLPagePath: removedExtension + ".html",
		Body:         string(fileContents),
	}

}

func renderTemplateFromPage(templateFilePath string, page Page) {

	t := template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	w, err := os.Create(page.HTMLPagePath)
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

	var dir string
	flag.StringVar(&dir, "dir", "", "The directory of text files")

	flag.Parse()

	// if textFilePath == "" {
	// 	panic("Missing the --file flag! Please supply one.")
	// }

	if textFilePath != "" {
		newPage := createFromTextFile(textFilePath)
		renderTemplateFromPage("template.tmpl", newPage)
	}

	if dir != "" {
		allTxtFiles, _ := filepath.Glob(dir + "/*.txt")
		for _, txtFile := range allTxtFiles {
			newPage := createFromTextFile(txtFile)
			renderTemplateFromPage("template.tmpl", newPage)
		}
	}
}
