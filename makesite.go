package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

// Content for body of txt file
type Content struct {
	Body string
}

func main() {

	content, err := ioutil.ReadFile("first-post.txt")
	contentAsString := string(content)
	contentStruct := Content{contentAsString}
	// if err != nil {
	// 	// A common use of `panic` is to abort if a function returns an error
	// 	// value that we don’t know how to (or want to) handle. This example
	// 	// panics if we get an unexpected error when creating a new file.
	// 	panic(err)
	// }
	// fmt.Print(string(fileContents))
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(os.Stdout, contentStruct)
	if err != nil {
		panic(err)
	}
}
