package templates

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostRendering struct {
	templ *template.Template
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

	return &PostRendering{templ: templ}, nil
}

func (r *PostRendering) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}
