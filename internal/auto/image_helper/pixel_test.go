package image_helper

import (
	"image"
	"image/color"
	"testing"
)

func TestGetPixelColor(t *testing.T) {
	tests := []struct {
		name      string
		img       image.Image
		x, y      int
		tolerance float32
		wantColor uint32
		wantErr   bool
	}{
		{
			name: "Valid pixel within bounds",
			img: func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 10, 10))
				img.Set(5, 5, color.RGBA{R: 255, G: 0, B: 0, A: 255})
				return img
			}(),
			x:         5,
			y:         5,
			tolerance: 0,
			wantColor: 0xFF0000FF,
			wantErr:   false,
		},
		{
			name: "Pixel out of bounds",
			img: func() image.Image {
				return image.NewRGBA(image.Rect(0, 0, 10, 10))
			}(),
			x:         15,
			y:         15,
			tolerance: 0,
			wantColor: 0,
			wantErr:   true,
		},
		{
			name: "Transparent pixel",
			img: func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 10, 10))
				img.Set(2, 2, color.RGBA{R: 0, G: 0, B: 0, A: 0})
				return img
			}(),
			x:         2,
			y:         2,
			tolerance: 0,
			wantColor: 0x00000000,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotColor, err := GetPixelColor(tt.img, tt.x, tt.y, tt.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPixelColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && *gotColor != tt.wantColor {
				t.Errorf("GetPixelColor() gotColor = %v, want %v", *gotColor, tt.wantColor)
			}
		})
	}
}
