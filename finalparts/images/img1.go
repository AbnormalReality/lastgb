package main

import (
	"github.com/mattn/go-gtk/gtk"
	"gocv.io/x/gocv"
	"image/color"
)

func main() {
	// Load face detection model
	classifier := gocv.NewCascadeClassifier()
	classifier.Load("haarcascade_frontalface_default.xml")
	defer classifier.Close()

	// Read input image
	img := gocv.IMRead("input.jpg", gocv.IMReadColor)
	defer img.Close()

	// Perform face detection
	rects := classifier.DetectMultiScale(img)

	// Draw bounding boxes around detected faces
	for _, r := range rects {
		gocv.Rectangle(img, r, color.RGBA{255, 0, 0, 0}, 3)
	}

	// Display the output
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Face Detection")
	window.Connect("destroy", gtk.MainQuit)

	image := gtk.NewImageFromPixbuf(gocv.ToPixbuf(img))
	window.Add(image)
	window.SetDefaultSize(img.Cols(), img.Rows())
	window.ShowAll()

	gtk.Main()
}
