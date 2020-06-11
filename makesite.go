package main

// <<<<<<< HEAD
import (
	"io/ioutil"
	"fmt"
	"html/template"
	"os"

)


type Contents struct {
	Content string
}

func main() {
	// """Writing file content"""
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))

	content := Contents{"first-post.txt"}

	t := template.Must(template.New("temp.html").ParseFiles("template.tmpl"))
	err2 := t.Execute(os.Stdout, content)
	if err2 != nil{
		panic(err)
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



func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}







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
