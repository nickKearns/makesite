package main

import (
	"context"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// Page is the struct of a page
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Body         string
}

func createFromTextFile(filepath string, translate bool) Page {

	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	fileContentsAsString := string(fileContents)

	if translate == true {
		fileContentsAsString = translateToFrench(string(fileContents))
	}

	removedExtension := strings.Split(filepath, ".txt")[0]

	return Page{
		TextFilePath: filepath,
		TextFileName: removedExtension,
		HTMLPagePath: removedExtension + ".html",
		Body:         fileContentsAsString,
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

func translateToFrench(textToTranslate string) string {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
		panic(err)
	}

	defer client.Close()

	translatedText, err := client.Translate(ctx,
		[]string{textToTranslate}, language.French,
		&translate.Options{
			Source: language.English,
			Format: translate.Text,
		})
	if err != nil {
		panic(err)
	}

	return translatedText[0].Text
}

func main() {

	var textFilePath string
	flag.StringVar(&textFilePath, "file", "", "Name or path to a text file")

	var dir string
	flag.StringVar(&dir, "dir", "", "The directory of text files")

	var translatePath string
	flag.StringVar(&translatePath, "translatePath", "", "translate to French and create an html file from a txt file")

	flag.Parse()

	// if textFilePath == "" {
	// 	panic("Missing the --file flag! Please supply one.")
	// }

	if textFilePath != "" {
		newPage := createFromTextFile(textFilePath, false)
		renderTemplateFromPage("template.tmpl", newPage)
	}

	if dir != "" {
		allTxtFiles, _ := filepath.Glob(dir + "/*.txt")
		for _, txtFile := range allTxtFiles {
			newPage := createFromTextFile(txtFile, false)
			renderTemplateFromPage("template.tmpl", newPage)
		}
	}

	if translatePath != "" {
		newPage := createFromTextFile(translatePath, true)
		renderTemplateFromPage("template.tmpl", newPage)
	}

}
