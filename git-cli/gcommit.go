package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Basic example of how to commit changes to the current branch to an existing
// repository.
func main() {
	// CheckArgs("./tmp/foo")
	// directory := os.Args[1]

	// Opens an already existing repository.
	r, err := git.PlainOpen("./tmp/foo")
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.
	Info("echo \"hello world!\" > README.md")
	filename := filepath.Join("./tmp/foo", "README.md")
	err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	CheckIfError(err)

	// Adds the new file to the staging area.
	Info("git add README.md")
	_, err = w.Add("README.md")
	CheckIfError(err)

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	Info("git commit -m \"adds my GO Exc\"")
	commit, err := w.Commit("adds my GO Exc", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Afzal",
			Email: "afz@k1.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}
