package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var green color.Color = color.RGBA{0, 0xff, 0, 0xff}
var yellow color.Color = color.RGBA{0xff, 0xff, 0, 0xff}
var orange color.Color = color.RGBA{0xff, 0x88, 0, 0xff}
var red color.Color = color.RGBA{0xff, 0, 0, 0xff}

func main() {
	x, y := 8, 7
	r := image.Rect(0, 0, x, y)
	img := image.NewRGBA(r)

	cl := green
	fillX(0, img, cl)
	fillX(x-1, img, cl)
	fillY(0, img, cl)
	fillY(y-1, img, cl)

	img.Set(3, 1, cl)
	img.Set(4, 1, cl)
	img.Set(2, 2, cl)
	img.Set(5, 2, cl)
	img.Set(2, 4, cl)
	img.Set(5, 4, cl)
	img.Set(3, 5, cl)
	img.Set(4, 5, cl)

	Fill(r, image.Pt(1, 3), makeFiller(img))
	outputImage(img, "test")
}

func fillX(x int, img *image.RGBA, cl color.Color) {
	b := img.Bounds();
	for p := image.Pt(x, b.Min.Y); p.Y < b.Max.Y; p.Y++ {
		img.Set(p.X, p.Y, cl)
	}
}

func fillY(y int, img *image.RGBA, cl color.Color) {
	b := img.Bounds();
	for p := image.Pt(b.Min.X, y); p.X < b.Max.X; p.X++ {
		img.Set(p.X, p.Y, cl)
	}
}

func makeFiller(img *image.RGBA) func(x, y int) bool {
	return func(x, y int) bool {
		fmt.Println(x, y);
		r, g, b, a := img.At(x, y).RGBA()
		r |= r >> 8
		g |= g >> 8
		b |= b >> 8
		cl := color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
		switch(cl) {
		case green:
			img.Set(x, y, yellow)
		case yellow:
			img.Set(x, y, orange)
		case orange:
			img.Set(x, y, red)
		}

		if a == 0 {
			img.Set(x, y, green)
			return true
		}
		
		return false
	}
}

func outputImage(img *image.RGBA, name string) error {
	outFile, err := os.Create(name + ".png")
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}
