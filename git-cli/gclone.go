package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

// Basic example of how to clone a repository using clone options.
func main() {
	url := "https://github.com/Cloud-Hacks/cloud-native-pro.git"
	// Clone the given repository to the given directory
	Info("git clone %s", url)

	r, err := git.PlainClone("./tmp/foo", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	CheckIfError(err)
	// CheckArgs("https://github.com/src-d/go-git.git", "./tmp/foo")
	// url := os.Args[1]
	// directory := os.Args[2]

	// // Clone the given repository to the given directory
	// Info("git clone %s %s ", url, directory)

	// r, err := git.PlainClone(directory, false, &git.CloneOptions{
	// 	URL:      url,
	// 	Progress: os.Stdout,
	// })

	// CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
