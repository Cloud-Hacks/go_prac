package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	git "github.com/go-git/go-git/v5"
)

func main() {
	token := "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"

	fs := memfs.New()

	Info("git clone https://github.com/Cloud-Hacks/devtro-app")
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           "https://github.com/Cloud-Hacks/devtro-app",
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
		Depth:         1,
		SingleBranch:  true,
		Auth:          &http.BasicAuth{Username: "Cloud-Hacks", Password: token},
		Progress:      os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")
}
