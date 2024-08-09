package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
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

//scanGitFolders returns a list of subfolders of 'folder' ending with '.git'
//Returns the base folder of the repo, the .git folder parent
//Recursively searches in the subfolders by passing as existsting 'folders' slice.

func scanGitFolders(folders []string, folder string) []string {
	//trim the last slash '/'
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	//optimise this later by checking if it's a git folder from the beginning
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "name_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}
	return folders
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
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
