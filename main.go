package main

import (
	"flag"
)

//scan given a apath craws it and its subfolders
//searching for Git repos

func scan(path string) {
	print("scan")
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
