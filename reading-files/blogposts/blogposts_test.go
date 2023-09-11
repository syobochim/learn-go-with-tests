package blogposts_test

import (
	"errors"
	"io/fs"
	"learn-go/reading-files/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("hot %d posts, wanted %d posts", len(posts), len(fs))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
