package main

import (
	"github.com/internot-blog/internot.blog.git/internal"
	"github.com/internot-blog/internot.blog.git/pkg"
)

func main() {
	cfg := internal.GetConfig("./proompt.json")

	// create a post (directory with markdown and image files)
	_, err := pkg.MakePost(cfg)
	if err != nil {
		panic(err)
	}

	// Site logo generator
	// prompt := "dead internet theory blog website logo named 'internot.blog'"
	// MakeLogos(cfg, prompt)
}
