package internal

import (
	"fmt"
	"os"
)

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
