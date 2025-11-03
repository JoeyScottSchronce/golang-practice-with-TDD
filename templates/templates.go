package templates

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostRendering struct {
	templ    *template.Template
	mdParser *parser.Parser
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

var (
	//go:embed blog_templates/*.gohtml
	postTemplates embed.FS
)

func NewPostRender() (*PostRendering, error) {
	templ, err := template.ParseFS(postTemplates, "blog_templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRendering{templ: templ, mdParser: parser}, nil
}

func (r *PostRendering) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func (r *PostRendering) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func newPostVM(p Post, r *PostRendering) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))

	return vm
}

// TODO: continue attempt to render the new Body from posts/post.md.
