package blogposts

import (
	"io"
	"strings"
)

type Post struct {
	Title string
}

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)

	if err != nil {
		return Post{}, err
	}

	title := strings.TrimPrefix(string(postData), "Title:")
	post := Post{Title: title}
	return post, nil
}
