package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/internot-blog/internot.blog.git/internal"
	"github.com/internot-blog/internot.blog.git/pkg"
)

func main() {
	// Load config
	cfg := internal.GetConfig("./proompt.json")

	// Define flags
	helpFlag := flag.Bool("help", false, "Print usage information")
	logoFlag := flag.Bool("logo", false, "Generate a logo")

	// Define flags for --gen-image
	genImageFlag := flag.Bool("gen-image", false, "Generate an image with prompt, seed, and steps")
	genImagePrompt := flag.String("prompt", "", "Prompt for generating the image")
	genImageSeed := flag.Int64("seed", rand.Int63(), "Seed for image generation")
	genImageWidth := flag.Int("width", 512, "Width of generated image")
	genImageHeight := flag.Int("height", 512, "Height of generated image")
	genImageSteps := flag.Int("steps", 50, "Number of steps for image generation")

	// Parse the flags
	flag.Parse()

	// Execute based on the flags
	switch {
	case *helpFlag:
		fmt.Println("Usage:")
		fmt.Println("--logo : Generate multiple site logos")
		fmt.Println("--gen-image --prompt <prompt> --seed <seed> --steps <steps> : Generate an image with specified parameters")
	case *logoFlag:
		// Generate potential site logos
		prompt := "dead internet theory blog website logo named 'internot.blog'"
		err := pkg.MakeLogos(cfg, prompt)
		if err != nil {
			panic(err)
		}
	case *genImageFlag:
		err := pkg.MakeImage(
			cfg,
			fmt.Sprintf("image.%d.%d-steps.png", *genImageSeed, *genImageSteps),
			*genImagePrompt,
			*genImageSeed,
			*genImageWidth,
			*genImageHeight,
			*genImageSteps,
		)
		if err != nil {
			panic(err)
		}
	default:
		// Create a post (directory with markdown and image files)
		for {
			_, err := pkg.MakePost(cfg)
			if err != nil {
				panic(err)
			}
		}
	}
}
