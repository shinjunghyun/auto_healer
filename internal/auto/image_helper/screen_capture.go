package image_helper

import (
	"image"
	"unsafe"

	"golang.org/x/sys/windows"
)

func CaptureScreen(hwnd uintptr) (image.Image, error) {
	user32 := windows.NewLazySystemDLL("user32.dll")
	gdi32 := windows.NewLazySystemDLL("gdi32.dll")

	getDC := user32.NewProc("GetDC")
	releaseDC := user32.NewProc("ReleaseDC")
	createCompatibleDC := gdi32.NewProc("CreateCompatibleDC")
	createCompatibleBitmap := gdi32.NewProc("CreateCompatibleBitmap")
	selectObject := gdi32.NewProc("SelectObject")
	bitBlt := gdi32.NewProc("BitBlt")
	deleteObject := gdi32.NewProc("DeleteObject")
	deleteDC := gdi32.NewProc("DeleteDC")
	getClientRect := user32.NewProc("GetClientRect")
	getDIBits := gdi32.NewProc("GetDIBits")

	// Get the client rect of the window
	var rect struct {
		Left, Top, Right, Bottom int32
	}
	_, _, _ = getClientRect.Call(hwnd, uintptr(unsafe.Pointer(&rect)))

	width := int(rect.Right - rect.Left)
	height := int(rect.Bottom - rect.Top)

	// Get the device context of the window
	hdc, _, _ := getDC.Call(hwnd)
	defer releaseDC.Call(hwnd, hdc)

	// Create a compatible DC and bitmap
	memDC, _, _ := createCompatibleDC.Call(hdc)
	defer deleteDC.Call(memDC)

	hBitmap, _, _ := createCompatibleBitmap.Call(hdc, uintptr(width), uintptr(height))
	defer deleteObject.Call(hBitmap)

	selectObject.Call(memDC, hBitmap)

	// Copy the screen content to the memory DC
	bitBlt.Call(memDC, 0, 0, uintptr(width), uintptr(height), hdc, 0, 0, 0x00CC0020) // SRCCOPY

	// Prepare to read the bitmap data
	var bmpInfo struct {
		Size          uint32
		Width         int32
		Height        int32
		Planes        uint16
		BitCount      uint16
		Compression   uint32
		SizeImage     uint32
		XPelsPerMeter int32
		YPelsPerMeter int32
		ClrUsed       uint32
		ClrImportant  uint32
	}
	bmpInfo.Size = uint32(unsafe.Sizeof(bmpInfo))
	bmpInfo.Width = int32(width)
	bmpInfo.Height = -int32(height) // Negative to indicate a top-down DIB
	bmpInfo.Planes = 1
	bmpInfo.BitCount = 32   // 32 bits per pixel
	bmpInfo.Compression = 0 // BI_RGB

	// Allocate memory for the bitmap data
	bitmapData := make([]byte, width*height*4)

	// Get the bitmap data
	getDIBits.Call(
		hdc,
		hBitmap,
		0,
		uintptr(height),
		uintptr(unsafe.Pointer(&bitmapData[0])),
		uintptr(unsafe.Pointer(&bmpInfo)),
		0,
	)

	// Convert the bitmap data to an image.RGBA
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	pix := img.Pix
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := (y*width + x) * 4
			j := (y * img.Stride) + (x * 4)
			pix[j] = bitmapData[i+2]   // R
			pix[j+1] = bitmapData[i+1] // G
			pix[j+2] = bitmapData[i]   // B
			pix[j+3] = bitmapData[i+3] // A
		}
	}

	return img, nil
}
