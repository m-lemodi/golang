package main

import (
	gim "github.com/ozankasikci/go-image-merge"
	"image/jpeg"
	"os"
)

func mergeImages(sourcePath, targetPath string) error {
	grids := []*gim.Grid{
		{ImageFilePath: sourcePath},
		{ImageFilePath: targetPath},
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

func ScreenHandler(source, target []byte, domain string) error {
	sourcePath := domain + "-pre-migration.png"
	targetPath := domain + "-post-migration.png"

	file1, err := os.Create(sourcePath)
	if err != nil {
		return err
	}
	defer file1.Close()
	file1.Write(source)

	file2, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file2.Close()
	file2.Write(target)

	err = mergeImages(sourcePath, targetPath)
	if err != nil {
		return err
	}
	return nil

}
