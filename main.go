package main

import (
	"flag"
	"image"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
)

// create the squares by cropping and saving to the output dir
func createsq(img image.Image, size int, i int, j int, m int) {
	filepath := filepath.Join(".", "output", strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(m)+".png")

	// create the rect and crop
	x0 := j * size
	x1 := x0 + size
	y0 := m * size
	y1 := y0 + size
	rect := image.Rect(x0, y0, x1, y1)
	img = imaging.Crop(img, rect)

	// save
	log.Printf("Saving %s", filepath)
	err := imaging.Save(img, filepath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

// loop through all the squares
func loop(src image.Image, si int, sj int, sm int) {
	// count from 0 to 7
	for i := si; i <= 7; i++ {
		// calculate the number of layers
		limit := int(math.Exp2(float64(i))) - 1

		// resize at this scale
		size := 256
		scale := int(math.Exp2(float64(i))) * size
		img := imaging.Resize(src, scale, 0, imaging.Lanczos)

		// count from 0 to limit
		for j := sj; j <= limit; j++ {
			// create the dir
			dirpath := filepath.Join(".", "output", strconv.Itoa(i), strconv.Itoa(j))
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				log.Fatalf("failed to create directory: %v", err)
			}

			// count the same limit for the images
			for m := sm; m <= limit; m++ {
				createsq(img, size, i, j, m)
			}
		}
	}
}

// main
func main() {
	// set flags
	mappath := flag.String("map", "map.png", "location of the map image")
	si := flag.Int("i", 0, "starting i value")
	sj := flag.Int("j", 0, "starting j value")
	sm := flag.Int("m", 0, "starting m value")
	flag.Parse()

	// Open the map
	src, err := imaging.Open(*mappath)
	if err != nil {
		log.Fatalf("failed to open map: %v", err)
	}

	loop(src, *si, *sj, *sm)
}
