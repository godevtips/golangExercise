package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

const project_path = "./golang/noisy_image/english"

func main() {

	resultRemovalError := removeFile("result.png")
	if resultRemovalError != nil {
		fmt.Println(resultRemovalError)
	}

	originalImage, err := getJpegImage("image")

	if err != nil {
		log.Fatal(err)
	}

	noisyOverlayImage, secondError := createNoisyOverLay(originalImage, 50)

	if secondError != nil {
		log.Fatal(secondError)
	}

	resultingImageError := mergeAndCreateNoisyImage(originalImage, noisyOverlayImage)

	if resultingImageError != nil {
		log.Fatal(resultingImageError)
	}

}

func removeFile(file string) error {

	// remove file if exist using Remove() function
	resultRemovalError := os.Remove(fmt.Sprintf("%s/%s", project_path, file))
	if resultRemovalError != nil {
		return resultRemovalError
	}

	return nil
}

func getJpegImage(filename string) (image.Image, error) {

	imageFile, err := os.Open(fmt.Sprintf("%s/%s.jpg", project_path, filename))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open original file: %s", err))
	}

	decodedImage, err := jpeg.Decode(imageFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to decode original image: %s", err))
	}

	defer func(image_file *os.File) {
		_ = image_file.Close()
	}(imageFile)

	return decodedImage, nil
}

func createNoisyOverLay(backgroundImage image.Image, opacity int) (image.Image, error) {

	var requested_opacity = uint8(0)
	if opacity > 0 && opacity < 256 {
		requested_opacity = uint8(math.Round(float64(opacity) * 2.56))
	}

	imageBound := backgroundImage.Bounds()
	imageWidth := imageBound.Dx()
	imageHeight := imageBound.Dy()
	myImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for p := 0; p < imageWidth*imageHeight; p++ {
		pixelOffset := 4 * p
		myImage.Pix[0+pixelOffset] = uint8(rand.Intn(256))
		myImage.Pix[1+pixelOffset] = uint8(rand.Intn(256))
		myImage.Pix[2+pixelOffset] = uint8(rand.Intn(256))
		myImage.Pix[3+pixelOffset] = requested_opacity
	}

	tempNoisyImageOutputFile, png_err := os.Create(fmt.Sprintf("%s/%s", project_path, "temp_noisy_image.png"))
	if png_err != nil {
		log.Fatal(png_err)
	}
	encodingError := png.Encode(tempNoisyImageOutputFile, myImage)
	if encodingError != nil {
		return nil, encodingError
	}

	tempNoisyImage, openImageError := os.Open(fmt.Sprintf("%s/%s", project_path, "temp_noisy_image.png"))
	if openImageError != nil {
		return nil, openImageError
	}

	result, pngDecodingError := png.Decode(tempNoisyImage)
	if pngDecodingError != nil {
		return nil, pngDecodingError
	}

	defer func(tempNoisyImage *os.File) {
		_ = tempNoisyImage.Close()
	}(tempNoisyImage)

	// remove temp_noisy_image.png using Remove() function
	removalError := removeFile("temp_noisy_image.png")
	if removalError != nil {
		return nil, removalError
	}

	return result, nil
}

func mergeAndCreateNoisyImage(originalImage image.Image, noiseOverlayImage image.Image) error {

	offset := image.Pt(0, 0)
	originalImageBounds := originalImage.Bounds()
	noisyImage := image.NewRGBA(originalImageBounds)
	draw.Draw(noisyImage, originalImageBounds, originalImage, image.Point{}, draw.Src)
	draw.Draw(noisyImage, noiseOverlayImage.Bounds().Add(offset), noiseOverlayImage, image.Point{}, draw.Over)

	resultImage, err := os.Create(fmt.Sprintf("%s/%s", project_path, "result.png"))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create result image: %s", err))
	}

	pngEncodingError := png.Encode(resultImage, noisyImage)
	if pngEncodingError != nil {
		return pngEncodingError
	}

	defer func(third *os.File) {
		_ = third.Close()
	}(resultImage)

	return nil
}
