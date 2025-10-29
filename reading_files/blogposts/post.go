package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title:"
	descriptionSeparator = "Description:"
	tagsSeparator        = "Tags:"
)

func newPost(postFile io.Reader) (Post, error) {

	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readTagLine := func(scanner *bufio.Scanner, tagName string) []string {
		scanner.Scan()
		rawTags := strings.Split(strings.TrimPrefix(scanner.Text(), tagName), ",")
		tags := make([]string, 0, len(rawTags))
		for _, tag := range rawTags {
			tags = append(tags, strings.TrimSpace(tag))
		}
		return tags
	}

	readBodyLine := func(scanner *bufio.Scanner) string {
		scanner.Scan() // ignore the "---" line

		buf := bytes.Buffer{}

		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        readTagLine(scanner, tagsSeparator),
		Body:        readBodyLine(scanner),
	}, nil
}
