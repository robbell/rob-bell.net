package processors

import (
	"bytes"
	"text/template"

	"github.com/robbell/hi/app/io"
	"github.com/robbell/hi/app/markdown"
)

// Post processor generates HTML pages for single posts
type Post struct{}

// Process generates an HTML page for a single post
func (p Post) Process(post markdown.Post) error {
	var postBuffer bytes.Buffer

	tmpl := template.Must(template.ParseFiles("./templates/post.html", "./templates/base.html"))
	if err := tmpl.Execute(&postBuffer, post); err != nil {
		return err
	}

	writePath := "./static/" + post.Permalink + ".html"
	if err := io.WriteToDisk(writePath, postBuffer.String()); err != nil {
		return err
	}

	return nil
}

// Finish is a NOP for a single post
func (p Post) Finish() error {
	return nil
}
