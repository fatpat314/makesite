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


func readDir() {
	dir := flag.String("dir", ".", "Find all .txt")
	file, err := os.Open(*dir)
	if err != nil{
		log.Fatal("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	// fmt.Println(list)
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
			// fmt.Println(name)
		}
		m += 1
	}
	fmt.Print("LIST: ", list)
}


	// fmt.Print(n)





func main() {

	readDir()








	// var files []string

	// dir := flag.String("dir", "makesite", "Find all .txt")
	// flag.Parse()

	// for _, file := range *dir{
	// 	// fmt.Println("hi")
	// 	fmt.Println(files)
	// 	fmt.Println(file)
	// 	return
	//
	// 	if filepath.Ext(*dir) == ".txt" {
	// 		files = append(files, filepath.Ext(*dir) )
	// 	}
	// }
	//
	// fmt.Println(files)

	// data, err := ioutil.ReadFile(*dir)
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// fmt.Println("Contents of file:", string(data))



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
	// Create new file
	create, _ := os.Create(newFile)
	err2 := t.Execute(create, content)
	if err2 != nil{
		panic(err2)
		}
}
