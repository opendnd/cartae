package main

import (
	"image"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
)

func createsq(src image.Image, i int, j int, m int) {
	filepath := filepath.Join(".", "output", strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(m)+".png")
	size := 256
	total := int(math.Exp2(float64(i))) * size
	src = imaging.Resize(src, total, 0, imaging.Lanczos)
	x0 := j * size
	x1 := x0 + size
	y0 := m * size
	y1 := y0 + size
	rect := image.Rect(x0, y0, x1, y1)
	src = imaging.Crop(src, rect)

	log.Printf("Saving %s", filepath)
	err := imaging.Save(src, filepath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func main() {
	// Open a test image.
	src, err := imaging.Open("map.png")
	if err != nil {
		log.Fatalf("failed to open map: %v", err)
	}

	// count from 0 to 7
	for i := 0; i <= 7; i++ {
		// calculate the number of layers
		limit := int(math.Exp2(float64(i))) - 1

		// count from 0 to limit
		for j := 0; j <= limit; j++ {
			// create the dir
			dirpath := filepath.Join(".", "output", strconv.Itoa(i), strconv.Itoa(j))
			err = os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				log.Fatalf("failed to create directory: %v", err)
			}

			// count the same limit for the images
			for m := 0; m <= limit; m++ {
				createsq(src, i, j, m)
			}
		}
	}
}
