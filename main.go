package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
)

func Argparse() (string, image.Rectangle) {
	parser := argparse.NewParser("flags", "configs")

	img := parser.String("i", "img", &argparse.Options{Required: true, Help: "Path of Original Image"})

	xmin := parser.Int("l", "xmin", &argparse.Options{Required: true, Help: "XMIN of Crop Potision"})
	ymin := parser.Int("t", "ymin", &argparse.Options{Required: true, Help: "YMIN of Crop Potision"})
	xmax := parser.Int("r", "xmax", &argparse.Options{Required: true, Help: "XMAX of Crop Potision"})
	ymax := parser.Int("b", "ymax", &argparse.Options{Required: true, Help: "YMAX of Crop Potision"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	min := image.Point{*xmin, *ymin}
	max := image.Point{*xmax, *ymax}
	bounds := image.Rectangle{min, max}

	return *img, bounds
}

func main() {
	impath, cropBounds := Argparse()

	src, err := os.Open(impath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	imgDst := image.NewRGBA(cropBounds)
	for y := cropBounds.Min.Y; y < cropBounds.Max.Y; y++ {
		for x := cropBounds.Min.X; x < cropBounds.Max.X; x++ {
			imgDst.Set(x, y, img.At(x, y))
		}
	}
	png.Encode(os.Stdout, imgDst)
	return
}
