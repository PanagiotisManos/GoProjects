package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type customWalker struct {
	outputFile *os.File
}

func (w *customWalker) Walk(name exif.FieldName, tag *tiff.Tag) error {
	_, err := fmt.Fprintf(w.outputFile, "%s: %s\n", name, tag)
	return err
}

func main() {
	// Path to the image file
	imagePath := "D:/Εικόνες και Βίντεο/Me & Marnie/DSC_1223.jpg"

	// Open the image file
	f, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Decode the EXIF data
	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new text file to store the EXIF data
	outputFile, err := os.Create("exif_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Create a customWalker instance with the output file
	walker := &customWalker{outputFile: outputFile}

	// Iterate over all EXIF fields using the customWalker
	if err := x.Walk(walker); err != nil {
		log.Fatal(err)
	}

	fmt.Println("EXIF data extracted and saved to exif_data.txt")
}