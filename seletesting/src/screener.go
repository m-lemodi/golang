package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	gim "github.com/ozankasikci/go-image-merge"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func screenAll() {
	//Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		save(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	// Capture all desktop region into an image.
	fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	save(img, "all.png")

}
func screenThis(screen int) {

	//Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= screen {
		panic("Bad screen")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	bounds := screenshot.GetDisplayBounds(screen)
	all = bounds.Union(all)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", screen, bounds.Dx(), bounds.Dy())
	save(img, fileName)

	fmt.Printf("#%d : %v \"%s\"\n", screen, bounds, fileName)

}

func ScreenHandler(source, target []byte) error {
	file1, err := os.Create("image1.png")
	if err != nil {
		return err
	}
	defer file1.Close()
	file1.Write(source)

	file2, err := os.Create("image2.png")
	if err != nil {
		return err
	}
	defer file2.Close()
	file2.Write(target)

	grids := []*gim.Grid{
		{ImageFilePath: "image1.png"},
		{ImageFilePath: "image2.png"},
	}

	// merge the images into a 2x1 grid
	rgba, err := gim.New(grids, 2, 1).Merge()
	if err != nil {
		return err
	}
	// save the output to jpg or png
	file, err := os.Create("output.png")
	if err != nil {
		return err
	}
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		return err
	}
	return nil

}
