package ocr_helper

// func TestReadString(t *testing.T) {
// 	// img, err := image_helper.CaptureBaramScreen()
// 	// if err != nil {
// 	// 	t.Fatalf("Failed to capture screen: %v", err)
// 	// 	return
// 	// }

// 	imagePath := "C:\\dev\\go\\baram_classic_keyboard_hooker\\auto_healer\\internal\\auto\\image_helper\\1baram_capture.png"

// 	file, err := os.Open(imagePath)
// 	if err != nil {
// 		t.Fatalf("Failed to open image file: %v", err)
// 		return
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		t.Fatalf("Failed to decode image: %v", err)
// 		return
// 	}

// 	x := 895
// 	y := 646
// 	width := 1003 - x
// 	height := 666 - y

// 	num, err := ReadString(nil, 0, 0, 0, 0)
// 	// num, err := ReadString(img, x, y, width, height)
// 	if err != nil {
// 		t.Fatalf("Failed to read string: %v", err)
// 		return
// 	}
// 	t.Logf("Read number: %d", num)
// }
