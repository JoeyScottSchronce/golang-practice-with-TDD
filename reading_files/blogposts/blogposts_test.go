package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/JoeyScottSchronce/golang-practice-with-TDD/reading_files/blogposts"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
	}
}
