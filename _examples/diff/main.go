package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// Basic example of how to diff w
func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]
	
	Info("git clone %s", directory)

	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	status, err := w.Status()
	CheckIfError(err)

	fmt.Println(status.String())

	ref, err := r.Head()
	CheckIfError(err)

	fmt.Println(ref.Hash())

	fmt.Println(r.Object(plumbing.AnyObject, ref.Hash()))

	commitObj, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	
	fmt.Println(commitObj.Files())
	
	

}
