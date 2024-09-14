package internal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/internot-blog/internot.blog.git/internal"
)

func TestFormatTextResponse(t *testing.T) {
	contents, err := os.ReadFile("./posts/13890e20/post.md")
	if err != nil {
		t.Fatalf("Failed to read post file.")
	}

	out := internal.FormtTextResponse(string(contents))
	if len(out) < 10 {
		t.Fatalf("Output too small.")
	}
	fmt.Println(out)
}
