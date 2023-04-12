package main

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"log"
)

func main() {
	fmt.Println("[Info] Face Detection using gocv")
	file := flag.String("image", "", "use -image and the path of the image to give the input")

	// Parsing the commandline arguments
	flag.Parse()

	if !flag.Parsed() {
		log.Fatal("Argument parsing failed")
	}

	if *file == "" {
		log.Fatal("Invalid input please specify the image path using -image flag")
	}

	classifier := gocv.NewCascadeClassifier()
	fmt.Println("[Info] Loading classifier model..")
	if !classifier.Load("models/cascade_classifier.xml") {
		log.Fatal("Failed to load the classifier file")
	}
	defer classifier.Close()

	image := gocv.IMRead(*file, gocv.IMReadAnyColor)
	gray := gocv.NewMat()
	defer gray.Close()

	// Changing the color mode into gray for ease of classification.
	gocv.CvtColor(image, &gray, gocv.ColorBGRToGray)

	fmt.Println("[Info] Detecting faces...")
	faces := classifier.DetectMultiScale(gray)

	for _, rect := range faces {
		// Dimensions
		x := rect.Min.X
		y := rect.Min.Y
		w := rect.Dx()
		h := rect.Dy()
		rectangle := image.Rect(x, y, x+w, y+h)
		red := color.RGBA{255, 0, 0, 1.0}
		gocv.Rectangle(&image, rectangle, red, 1)

		// Writing image into a file.
		if !gocv.IMWrite("output.jpg", image) {
			log.Fatal("Failed to write the output image")
		}

		window := gocv.NewWindow("Face Detection")
		window.IMShow(image)
		window.WaitKey(10000)
	}
}
