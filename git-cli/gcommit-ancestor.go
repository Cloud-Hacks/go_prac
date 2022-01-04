package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Example how to resolve a commit revision (branch(origin/main) or (Any feature branch)) into its commit counterpart
// git merge-base --all origin/main origin/HEAD
func main() {
	dir := os.Args[1]
	// CheckIfError(err)

	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	revision := "origin/main"

	// Resolve revision into a sha1 commit, only some revisions are resolved
	// look at the doc to get more details
	Info("git rev-parse %s", revision)

	revHash, err := repo.ResolveRevision(plumbing.Revision(revision))
	// fmt.Println(revHash)
	CheckIfError(err)
	revCommit, err := repo.CommitObject(*revHash)
	fmt.Println(revCommit)
	CheckIfError(err)

	headRef, err := repo.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	headCommit, err := repo.CommitObject(headRef.Hash())
	CheckIfError(err)
	fmt.Println(headCommit)

	isAncestor, err := headCommit.IsAncestor(revCommit)

	CheckIfError(err)

	fmt.Printf("Is the HEAD an Ancestor of origin/main? : %v\n", isAncestor)
}
