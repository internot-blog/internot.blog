package internal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/internot-blog/internot.blog.git/internal"
)

var cfg internal.Config

func init() {
	cfg = internal.GetConfig("../proompt.json")
}

func TestGenTextPrompt(t *testing.T) {
	// HACK: this is bad practice, if tests fail, this may be why
	// solves template directory issue inside of prompt gen and frontmatter gen funcs
	if err := os.Chdir("../"); err != nil {
		panic("Failed to change directory")
	}

	fmt.Println(internal.GenTextPrompt(cfg))
}

func TestGenImagePrompt(t *testing.T) {
	fmt.Println(internal.GenImagePrompt(cfg))
}
