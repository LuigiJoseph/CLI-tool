package main

import (
	"flag"
	"fmt"
)

//scan given a apath craws it and its subfolders
//searching for Git repos

func scan(path string) {
	fmt.Printf("Found folders: \n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

// stats generates a nice graph of your Git Contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
