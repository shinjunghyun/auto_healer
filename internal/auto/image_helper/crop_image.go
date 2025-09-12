package image_helper

import "image"

func CropImage(img image.Image, rect image.Rectangle) (image.Image, error) {
	cropped := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			cropped.Set(x, y, img.At(x, y))
		}
	}
	return cropped, nil
}
