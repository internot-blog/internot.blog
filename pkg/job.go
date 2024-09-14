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
	width := 512
	height := 512
	steps := 10

	err = MakeImage(cfg, fmt.Sprintf("%s/image.%d.png", postDir, seed), prompt, seed, width, height, steps)
	if err != nil {
		return "", err
	}

	fmt.Printf("\n---------------\nCreated post in %s\n", postDir)

	return postDir, nil
}

func MakeAd(cfg internal.Config) error {
	return nil
}

func MakeText(cfg internal.Config, textFilename string, prompt string) error {
	fmt.Println("Text prompt:", prompt)

	response, err := internal.GenText(cfg, prompt)
	if err != nil {
		return err
	}

	formattedResponse := internal.FormtTextResponse(response)

	internal.SaveText(formattedResponse, textFilename)
	fmt.Println(formattedResponse)

	return nil
}

func MakeImage(cfg internal.Config, imageFilename string, prompt string, seed int64, width int, height int, steps int) error {
	guidance := 7.5

	fmt.Printf("Image prompt: %s\n", prompt)
	fmt.Printf("Seed: %d\n", seed)
	fmt.Printf("Width: %d\n", width)
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Steps: %d\n", steps)

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

	width := 512
	height := 512
	steps := 10

	for i := 0; i < 100; i++ {
		seed := rand.Int63()
		err = MakeImage(cfg, fmt.Sprintf("logos/%d.png", seed), prompt, seed, width, height, steps)
		if err != nil {
			return err
		}
	}

	return nil
}
