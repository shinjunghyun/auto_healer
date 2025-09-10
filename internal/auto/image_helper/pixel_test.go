package image_helper

import (
	"image"
	"image/color"
	"testing"
)

func TestGetPixelColor(t *testing.T) {
	// Create a test image
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	img.Set(5, 5, color.RGBA{R: 255, G: 128, B: 64, A: 255}) // Set a specific pixel color

	tests := []struct {
		name      string
		x, y      int
		wantColor uint32
		wantErr   bool
	}{
		{
			name:      "Valid pixel",
			x:         5,
			y:         5,
			wantColor: uint32((255 << 24) | (128 << 16) | (64 << 8) | 255),
			wantErr:   false,
		},
		{
			name:      "Out of bounds (negative coordinates)",
			x:         -1,
			y:         -1,
			wantColor: 0xFFFFFFFF, // Use 0xFFFFFFFF as a sentinel value for invalid color
			wantErr:   true,
		},
		{
			name:      "Out of bounds (exceeds image size)",
			x:         10,
			y:         10,
			wantColor: 0xFFFFFFFF, // Use 0xFFFFFFFF as a sentinel value for invalid color
			wantErr:   true,
		},
		{
			name:      "Another valid pixel",
			x:         0,
			y:         0,
			wantColor: 0, // Default color is black with full alpha
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotColor, err := GetPixelColor(img, tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPixelColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if uint32(gotColor) != tt.wantColor {
				t.Errorf("GetPixelColor() = %v, want %v", gotColor, tt.wantColor)
			}
		})
	}
}
