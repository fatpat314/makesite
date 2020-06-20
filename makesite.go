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
	dir := flag.String("dir", ".", "Find all .txt")
	file, err := os.Open(*dir)
	if err != nil{
		log.Fatal("failed opening directory: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0)
	n := 0
	for _, name := range list {
		if strings.Split(name, "")[len(name)-1] != "t"  {
			list[n] = ""
		}
		n += 1
	}
	m := 0
	for _, name := range list{
		// fmt.Print(list)
		if strings.Split(name, ".")[0] == "" {
			list[m] = ""
		}
		m += 1
	}
	for _, name := range list{
		if name == ""{
			// do nothing
		}else{
		// Basically the same code as render, I would like to figure out
		// how to organize the inputs for RENDER so it can be modular and
		 // so I can begin to have all of my function calls in main.
		fmt.Print(name, "\n")
		path := []string{
			"template.tmpl",
		}
		extension := ".html"
		file := name

		fileContents, err := ioutil.ReadFile(name)
		if err != nil {
			// panic
			panic(err)
		}
		// split from "." and replate with extention
		newFile := strings.Split(file, ".") [0] + extension
		content := Content{string(fileContents)}
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
