package templates_test

import (
	"bytes"
	"testing"

	"example.com/templates"
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

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := templates.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Good Morning Everyone</h1>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
