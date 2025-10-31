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
			Body:        "This is my first post on my new blog",
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
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = templates.Post{
			Title:       "Good Morning Everyone",
			Body:        "This is my first post on my new blog",
			Description: "This is a description of my first post",
			Tags:        []string{"go", "tdd", "yay"},
		}
	)

	postRendering, err := templates.NewPostRender()

	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRendering.Render(io.Discard, aPost)
	}
}
