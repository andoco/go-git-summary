package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

func TestSpec(t *testing.T) {
	Convey("Given a root folder", t, func() {
		fs := memfs.New()
		So(fs.MkdirAll("/tmp/repo1", 0777), ShouldBeNil)
		s := filesystem.NewStorage(fs, cache.NewObjectLRUDefault())
		So(s.Init(), ShouldBeNil)
		So(fs.MkdirAll("/tmp/foo/repo2", 0777), ShouldBeNil)

		finder := &DefaultRepoFinder{}

		Convey("Should find nested git repo folders", func() {
			repos, err := finder.Find(fs, "/tmp")
			So(err, ShouldBeNil)
			So(repos, ShouldHaveLength, 2)
		})
	})
}
