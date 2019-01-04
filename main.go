package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	path := os.Args[1]

	if err := checkRepoPath(path); err != nil {
		log.Fatal(errors.Wrapf(err, "checking repo at path", path))
	}
}

func checkRepoPath(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return errors.Wrapf(err, "opening repo at path %s", path)
	}

	w, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, "getting worktree")
	}

	status, err := w.Status()
	if err != nil {
		return errors.Wrap(err, "getting status")
	}

	if !status.IsClean() {
		fmt.Printf("DIRTY WORKTREE!\n")
		fmt.Printf("%v", status)
	}

	return nil
}
