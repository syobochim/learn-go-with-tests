package blogposts_test

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFile{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hora")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("hot %d posts, wanted %d posts", len(posts), len(fs))
	}
}
