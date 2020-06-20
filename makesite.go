package main

// <<<<<<< HEAD
import (
	"io/ioutil"
	"fmt"
	"html/template"
	"os"
	"flag"
	"strings"
	"log"


)

// Creating Contents structure
type Content struct {
	// Setting content structure as a string
	Content string
}

// ```What is the right way to make the functions modular? I am not sure how to call
// these functions when the main() does not accept input or return output ```


func renDir() {
	// Set flag input default to this directory
	dir := flag.String("dir", ".", "Find all .txt")
	// Open the directory
	file, err := os.Open(*dir)
	// error
	if err != nil{
		log.Fatal("failed opening directory: %s", err)
	}
	defer file.Close()
	// Make a list of the files in the directory
	list, _ := file.Readdirnames(0)
	// set counter to 0
	n := 0
	// start looping over the names of the files in list
	for _, name := range list {
		// if the name of the file does not end in a "t", remove from list
		if strings.Split(name, "")[len(name)-1] != "t"  {
			// How do I actually delete this from the list?
			list[n] = ""
		}
		// increse counter for next itteration
		n += 1
	}
	// start another counter at 0
	m := 0
	// start looping over the names that are leftover in list
	for _, name := range list{
		// if there is nothing before the ".", remove from list
		if strings.Split(name, ".")[0] == "" {
			//How, delete, go, slice? ¯\_(ツ)_/¯
			list[m] = ""
		}
		// increse counter
		m += 1
	}
	// start looping over the list... again...
	for _, name := range list{
		// If the name of the file is nothing ""
		if name == ""{
			// do nothing
		}else{
		// This is basically the same code as render, I would like to figure out
		// how to organize the inputs for RENDER so it can be modular and
		 // so I can begin to have all of my function calls in main.
		fmt.Print(name, "\n")
		// set path for templating files
		path := []string{
			"template.tmpl",
		}

		extension := ".html"
		file := name
		// save file contents of each .txt folder
		fileContents, err := ioutil.ReadFile(name)
		// err
		if err != nil {
			// panic
			panic(err)
		}
		// split from "." and replate with new extention
		newFile := strings.Split(file, ".") [0] + extension
		// populate the contents of the new html template
		content := Content{string(fileContents)}
		// make the new template
		t := template.Must(template.New("template.tmpl").ParseFiles(path...))
		// Create new file
		create, _ := os.Create(newFile)
		err2 := t.Execute(create, content)
		if err2 != nil{
			panic(err2)
			}
		}
	}
}
// -------------------------------------------------------------------------//
// Okay, I really need some help figuring out how to make this more modular,
// This is getting ridiculous...
// ------------------------------------------------------------------------//

// func read(){
// 	fileContents, err := ioutil.ReadFile("first-post.txt")
// 	// if there is an error, panic
// 	if err != nil {
// 		// panic
// 		panic(err)
// 	}
// 	return
// }


func main() {

	renDir()


	// READ
	// """Reading file content"""
	fileContents, err := ioutil.ReadFile("first-post.txt")
	// if there is an error, panic
	if err != nil {
		// panic
		panic(err)
	}
	// print fileContent
	// fmt.Print(string(fileContents))

	// TRANSPOSE
	file := flag.String("file", "post.html", "This flag represents the name of any `.txt` file in the same directory as your program.")
	flag.Parse()
	// Print file to the console
	// fmt.Println(*file)

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
	// Create new file
	create, _ := os.Create(newFile)
	err2 := t.Execute(create, content)
	if err2 != nil{
		panic(err2)
		}
}
