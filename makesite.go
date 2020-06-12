package main

// <<<<<<< HEAD
import (
	"io/ioutil"
	"fmt"
	"html/template"
	"os"
	"flag"

)

// Creating Contents structure
type Content struct {
	// Setting content structure as a string
	Content string
}
// Create main
func main() {

	file := flag.String("file", ".txt", "This flag represents the name of any `.txt` file in the same directory as your program.")
	flag.Parse()
	fmt.Println(*file)
	// output := flag.Bool("output", false, "Should there be output?")
	// input := flag.String("input", "file.csv", "the pather to the input file")
	// flag.Parse()
	// fmt.Println(*output)
	// fmt.Println(*input)

	// """Reading file content"""
	fileContents, err := ioutil.ReadFile("first-post.txt")
	// if there is an error, panic
	if err != nil {
		// panic
		panic(err)
	}
	// print fileContent
	fmt.Print(string(fileContents))

	// set file path for ParseFile
	path := []string{
		"template.tmpl",
	}
	// set content to a string that is fileContent and point to Content struct
	content := Content{string(fileContents)}
	// Make the new template from parsed path
	t := template.Must(template.New("template.tmpl").ParseFiles(path...))
	// Create html
	create, _ := os.Create(*file)
	err2 := t.Execute(create, content)
	if err2 != nil{
		panic(err2)
		}
	}



	// fileContents, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(string(fileContents))
	//
	// t := template.Must(template.New("template.tmpl").ParseFiles("first-post.txt"))
	// err2 := t.Execute(os.Stdout, content)
	// if err2 != nil {
	// 	panic(err)
	// }





// func render(){
//
// // ToDo: Look up arguments for template.New function in documentation
// 	t := template.Must(template.New("template.tmpl").ParseFiles("first-post.txt"))
// 	err2 := t.Execute(os.Stdout, content)
// 	if err2 != nil {
// 		panic(err)
// 	}
// }



// func readFile() string {
// 	fileContents, err := ioutil.ReadFile("first-post.txt")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return string(fileContents)
// }







// func temp_write(){
// 	paths := []string{
// 		"template.tmpl",
// 	}
// 	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
// 	err := t.Execute(os.Stdout, template)
// 	if err != nil {
// 		panic(err)
// 		}
// }


// func main() {
// 	bytesToWrite := []byte("hello\ngo\n")
// 	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func renderTemplate()
// =======
// import "fmt"
//
// func main() {
// 	fmt.Println("Hello, world!")
// }
// >>>>>>> 9514ac8a2c135a448a2b15a4b246dcd5d59ee7bf
