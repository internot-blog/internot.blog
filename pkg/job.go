package pkg

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/internot-blog/internot.blog.git/internal"
)

func MakePost(cfg internal.Config) (string, error) {
	postId := internal.GenUniqueId()
	postDir := "posts/" + postId

	err := os.MkdirAll(postDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	prompt := internal.GenTextPrompt(cfg)
	err = MakeText(cfg, postDir+"/post.md", prompt)
	if err != nil {
		return "", err
	}

	seed := rand.Int63()
	err = MakeImage(cfg, fmt.Sprintf("%s/image.%d.png", postDir, seed), prompt, seed)
	if err != nil {
		return "", err
	}

	fmt.Printf("\n---------------\nCreated post in %s\n", postDir)

	return postDir, nil
}

func MakeText(cfg internal.Config, textFilename string, prompt string) error {
	fmt.Println("Prompt:", prompt)

	response, err := internal.GenText(cfg, prompt)
	if err != nil {
		return err
	}

	formattedResponse := internal.FormtTextResponse(response)

	internal.SaveText(formattedResponse, textFilename)
	fmt.Println(formattedResponse)

	return nil
}

func MakeImage(cfg internal.Config, imageFilename string, prompt string, seed int64) error {
	width := 512
	height := 512
	steps := 10
	guidance := 7.5

	imageData, err := internal.GenImage(cfg, prompt, seed, width, height, steps, guidance)
	if err != nil {
		return err
	}

	internal.SaveImage(imageData, imageFilename)
	fmt.Println("Image saved.")

	return nil
}

func MakeLogos(cfg internal.Config, prompt string) error {
	err := os.MkdirAll("logos", os.ModePerm)
	if err != nil {
		return err
	}

	for i := 0; i < 100; i++ {
		seed := rand.Int63()
		err = MakeImage(cfg, fmt.Sprintf("logos/%d.png", seed), prompt, seed)
		if err != nil {
			return err
		}
	}

	return nil
}
