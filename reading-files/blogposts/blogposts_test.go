package blogposts_test

import (
	"errors"
	"io/fs"
	"learn-go/reading-files/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hora")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("hot %d posts, wanted %d posts", len(posts), len(fs))
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
