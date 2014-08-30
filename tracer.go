package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img.Set(5, 5, color.RGBA{255, 0, 0, 255})

	outfile, err := os.Create("test.png")
	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}
