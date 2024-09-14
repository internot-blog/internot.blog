package internal

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func GenUniqueId() string {
	randomNumber := rand.Uint32()

	// Format the number as an 8-digit hexadecimal
	postId := fmt.Sprintf("%08x", randomNumber)

	return postId
}

func SaveText(text string, textFilename string) {
	// Create or open the file for writing
	file, err := os.Create(textFilename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the text to the file
	_, err = io.WriteString(file, text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func SaveImage(imageData []byte, imageFilename string) {
	file, err := os.Create(imageFilename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(imageData)
	if err != nil {
		fmt.Println("Error writing image data to file:", err)
	}
}
