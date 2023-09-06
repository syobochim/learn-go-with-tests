package blogposts_test

import (
	"learn-go/reading-files/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hora")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("hot %d posts, wanted %d posts", len(posts), len(fs))
	}
}
