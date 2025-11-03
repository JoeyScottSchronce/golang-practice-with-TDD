package templates_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/JoeyScottSchronce/golang-practice-with-TDD/templates"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = templates.Post{
			Title:       "Good Morning Everyone",
			Body:        "This is the original post on my new blog",
			Description: "This is a description of my first post",
			Tags:        []string{"go", "tdd", "yay"},
		}
	)

	postRendering, err := templates.NewPostRender()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRendering.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []templates.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRendering.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = templates.Post{
			Title:       "Good Morning Everyone",
			Body:        "This is the original post on my new blog",
			Description: "This is a description of my first post",
			Tags:        []string{"go", "tdd", "yay"},
		}
	)

	postRendering, err := templates.NewPostRender()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		err := postRendering.Render(io.Discard, aPost)
		if err != nil {
			b.Fatal(err)
		}
	}
}
