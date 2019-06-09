package main

import (
	"io/ioutil"
	"path"
)

type markdownRepository interface {
	getAll() []markdownPost
	getByTitle(string) markdownPost
}

type fileSystemMarkdownRepository struct {
	sourceDirectory string
}

func newFileSystemMarkdownRepository() markdownRepository {
	return fileSystemMarkdownRepository{
		sourceDirectory: "./static",
	}
}

func (s fileSystemMarkdownRepository) getAll() []markdownPost {
	files, err := ioutil.ReadDir(s.sourceDirectory)

	if err != nil {
		panic(err)
	}

	var posts []markdownPost

	for _, file := range files {
		if post, ok := newMarkdownPost(path.Join(s.sourceDirectory, file.Name())); ok {
			posts = append(posts, post)
		}
	}

	return posts
}

func (s fileSystemMarkdownRepository) getByTitle(title string) markdownPost {
	post, _ := newMarkdownPost(path.Join(s.sourceDirectory, title+".md"))
	return post
}
