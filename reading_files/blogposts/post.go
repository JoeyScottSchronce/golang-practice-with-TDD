package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title:"
	descriptionSeparator = "Description:"
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := strings.TrimPrefix(readLine(), titleSeparator)
	description := strings.TrimPrefix(readLine(), descriptionSeparator)

	return Post{Title: title, Description: description}, nil
}
