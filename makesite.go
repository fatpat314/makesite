package main

// <<<<<<< HEAD
import (
	"io/ioutil"
	"fmt"
	"html/template"
	"os"
	"flag"
	"strings"

)

// Creating Contents structure
type Content struct {
	// Setting content structure as a string
	Content string
}

// ```What is the right way to make the functions modular? I am not sure how to call
// these functions when the main() does not accept input or return output ```
// Create main
func main() {
	// READ
	// """Reading file content"""
	fileContents, err := ioutil.ReadFile("first-post.txt")
	// if there is an error, panic
	if err != nil {
		// panic
		panic(err)
	}
	// print fileContent
	fmt.Print(string(fileContents))


	// TRANSPOSE
	file := flag.String("file", "post.html", "This flag represents the name of any `.txt` file in the same directory as your program.")
	flag.Parse()
	// Print file to the console
	fmt.Println(*file)
	// set extention to a .html
	extension := ".html"
	// split from "." and replate with extention
	newFile := strings.Split(*file, ".") [0] + extension


	// RENDER
	// set file path for ParseFile
	path := []string{
		"template.tmpl",
	}
	// set content to a string that is fileContent and point to Content struct
	content := Content{string(fileContents)}
	// Make the new template from parsed path
	t := template.Must(template.New("template.tmpl").ParseFiles(path...))
	// Create html
	create, _ := os.Create(newFile)
	err2 := t.Execute(create, content)
	if err2 != nil{
		panic(err2)
		}
}
