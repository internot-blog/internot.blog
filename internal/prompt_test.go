package internal_test

import (
	"fmt"
	"testing"

	"github.com/internot-blog/internot.blog.git/internal"
)

var cfg internal.Config

func init() {
	cfg = internal.GetConfig("../proompt.json")
}

func TestGenTextPrompt(t *testing.T) {
	fmt.Println(internal.GenTextPrompt(cfg))
}

func TestGenImagePrompt(t *testing.T) {
	fmt.Println(internal.GenImagePrompt(cfg))
}
