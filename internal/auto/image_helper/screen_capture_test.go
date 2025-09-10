package image_helper

import (
	"image/png"
	"os"
	"testing"
)

func TestCaptureBaramScreen(t *testing.T) {
	// Attempt to capture the Baram screen
	img, err := CaptureBaramScreen()
	if err != nil {
		t.Errorf("Failed to capture Baram screen: %v\n", err)
		return
	}

	// Save the captured image to a file for verification
	file, err := os.Create("./baram_capture.png")
	if err != nil {
		t.Errorf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Errorf("Failed to save image: %v\n", err)
		return
	}

	t.Logf("Baram screen captured and saved to ./baram_capture.png")
}
