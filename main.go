package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/format/gitignore"
)

func main() {
	path := os.Args[1]

	repos, err := findRepos(path)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "finding repos at path %s", path))
	}

	for _, p := range repos {
		fmt.Printf("Checking repo %s\n", p)
		if err := checkRepoPath(p); err != nil {
			log.Fatal(errors.Wrapf(err, "checking repo at path %s", p))
		}
	}

}

// findRepos will find all the repo folders that are children of path.
func findRepos(root string) ([]string, error) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, errors.Wrapf(err, "reading dir %v", root)
	}

	repos := []string{}

	for _, file := range files {
		if file.IsDir() {
			tmpPath := path.Join(root, file.Name())
			if isGitRepo(tmpPath) {
				repos = append(repos, tmpPath)
			}
		}
	}

	return repos, nil
}

// isGitRepo checks if the folder at path p contains a git repo.
func isGitRepo(p string) bool {
	if _, err := os.Stat(path.Join(p, ".git")); !os.IsNotExist(err) {
		return true
	}
	return false
}

func checkRepoPath(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return errors.Wrapf(err, "opening repo at path %s", path)
	}

	fs := osfs.New("/")
	patterns, err := gitignore.LoadGlobalPatterns(fs)
	if err != nil {
		return errors.Wrap(err, "loading global patterns")
	}

	w, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, "getting worktree")
	}
	w.Excludes = patterns

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
