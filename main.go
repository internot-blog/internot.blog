package main

import (
	"fmt"

	"github.com/internot-blog/internot.blog.git/internal"
)

func main() {
	cfg := internal.GetConfig("./proompt.json")

	fmt.Println(internal.GenTextPrompt(cfg))
}
