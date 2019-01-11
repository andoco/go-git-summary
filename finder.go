package main

import billy "gopkg.in/src-d/go-billy.v4"

type RepoFinder interface {
	Find(fs billy.Filesystem, root string) ([]string, error)
}

type DefaultRepoFinder struct{}

func (rf *DefaultRepoFinder) Find(fs billy.Filesystem, root string) ([]string, error) {
	return nil, nil
}
