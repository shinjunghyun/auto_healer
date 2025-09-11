package ocr_helper

// func ReadString(img image.Image, x, y, width, height int) (uint64, error) {
// 	t, err := ocr.New(ocr.TesseractPath(`C:\Program Files\Tesseract-OCR\tesseract.exe`))
// 	if err != nil {
// 		return 0, err
// 	}
// 	extractedText, err := t.TextFromImageFile(`C:\dev\go\baram_classic_keyboard_hooker\auto_healer\internal\auto\image_helper\2baram_capture.png`)
// 	if err != nil {
// 		return 0, err
// 	}
// 	log.Trace().Msgf("Extracted Text: %s", extractedText)
// 	return 0, nil
// }

// func ReadString(img image.Image, x, y, width, height int) (uint64, error) {
// 	// --- 1. 지정된 영역 크롭 ---
// 	subImg := img.(interface {
// 		SubImage(r image.Rectangle) image.Image
// 	}).SubImage(image.Rect(x, y, x+width, y+height))

// 	// --- 2. image.Image → PNG → Mat 변환 ---
// 	tmpIn, err := os.CreateTemp("", "ocr-in-*.png")
// 	if err != nil {
// 		return 0, err
// 	}
// 	if err := png.Encode(tmpIn, subImg); err != nil {
// 		return 0, err
// 	}
// 	tmpIn.Close()
// 	defer os.Remove(tmpIn.Name())

// 	src := gocv.IMRead(tmpIn.Name(), gocv.IMReadColor)
// 	if src.Empty() {
// 		return 0, fmt.Errorf("failed to read cropped image as Mat")
// 	}
// 	defer src.Close()

// 	// --- 3. 전처리 (Grayscale → Threshold → Morphology) ---
// 	gray := gocv.NewMat()
// 	defer gray.Close()
// 	gocv.CvtColor(src, &gray, gocv.ColorBGRToGray)

// 	binary := gocv.NewMat()
// 	defer binary.Close()
// 	gocv.Threshold(gray, &binary, 0, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)

// 	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(2, 2))
// 	defer kernel.Close()
// 	gocv.Dilate(binary, &binary, kernel)
// 	gocv.Erode(binary, &binary, kernel)

// 	// --- 4. 전처리 결과 PNG 저장 (디버깅용) ---
// 	tmpOut, err := os.CreateTemp("", "ocr-out-*.png")
// 	if err != nil {
// 		return 0, err
// 	}
// 	gocv.IMWrite(tmpOut.Name(), binary)
// 	tmpOut.Close()
// 	defer os.Remove(tmpOut.Name())

// 	// --- 5. Tesseract OCR ---
// 	client := gosseract.NewClient()
// 	defer client.Close()

// 	client.SetLanguage("eng")
// 	client.SetImage(tmpOut.Name())
// 	client.SetVariable("tessedit_char_whitelist", "0123456789")
// 	client.SetVariable("classify_bln_numeric_mode", "1") // 숫자 전용 모드
// 	client.SetVariable("tessedit_pageseg_mode", "7")     // 한 줄 텍스트
// 	client.SetVariable("tessedit_ocr_engine_mode", "3")  // default engine

// 	text, err := client.Text()
// 	if err != nil {
// 		return 0, err
// 	}

// 	// --- 6. 후처리 (숫자만 남기기) ---
// 	text = strings.TrimSpace(text)
// 	text = strings.Map(func(r rune) rune {
// 		if r >= '0' && r <= '9' {
// 			return r
// 		}
// 		return -1
// 	}, text)

// 	if text == "" {
// 		return 0, nil
// 	}
// 	return strconv.ParseUint(text, 10, 64)
// }
