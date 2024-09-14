package internal

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

var categories = []string{
	"technology", "music", "politics", "pop-culture", "movies", "video-games",
	"journalism", "drama", "cars", "planes", "engineering", "history",
	"the", "of", "and", "but", "why", "today", "business",
}

func genFrontMatter(resp string) string {
	tmpl, err := template.ParseFiles("templates/frontmatter.md")
	// tmpl, err := template.ParseFiles("../templates/frontmatter.md")
	if err != nil {
		panic(fmt.Sprintln("Error reading templates/frontmatter.md: ", err))
	}

	var title string
	var desc string
	lines := strings.Split(resp, "\n")
	for _, line := range lines {
		trimmedLine := strings.Trim(line, ` #"*'`)
		strings.ReplaceAll(line, "\"", "\\\"")
		if len(trimmedLine) == 0 {
			continue
		}

		parts := strings.FieldsFunc(trimmedLine, func(r rune) bool {
			// TEST: make sure these are the best chars to split on (;,?)
			return strings.ContainsRune(".:", r)
		})

		title = parts[0]
		if len(parts) > 1 {
			desc = strings.TrimSpace(parts[1])
		}

		break
	}

	keywords := FindNGrams(resp)

	frontMatter := struct {
		Title       string
		Description string
		Date        string
		Tags        []string
		Categories  string
	}{
		Title:       title,
		Description: desc,
		Date:        time.Now().Format("2006-01-02"),
		Tags:        keywords,
		Categories:  categories[rand.Intn(len(categories))],
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, frontMatter)
	if err != nil {
		panic(fmt.Sprintln("Error executing frontMatter template: ", err))
	}

	return result.String()
}

// TODO: do any other required post-processing here
func FormtTextResponse(resp string) string {
	frontMatter := genFrontMatter(resp)

	// TODO: remove the first line (its used as the header)

	return fmt.Sprintf("%s\n%s", frontMatter, resp)
}
