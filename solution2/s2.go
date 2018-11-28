package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Edges []color.Color

type Pixel struct {
	Edges
	IsSame bool
}

//Given two image detemine is one image is a subsect of the other.
//Both jpg, resoluion is the same and bit depth
func main() {
	img, f := openImage("img/trump.jpeg")
	p := getOriginalImagePixel(img)
	img2, f2 := openImage("img/trump2.jpeg")
	p2 := getImagePixels(img2)

	result := p.areImagesEqual(p2, f, f2)
	if result {
		fmt.Println(result)
		os.Exit(0)
	}

	fmt.Println(result)
}

func openImage(filename string) (image.Image, *os.File) {
	var img image.Image
	file1, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	if img, err = jpeg.Decode(file1); err != nil {
		panic(err)
	}

	return img, file1
}

func (img Pixel) areImagesEqual(img2 Pixel, file1, file2 *os.File) bool {
	fmt.Println(len(img.Edges))
	fmt.Println(len(img2.Edges))
	var equal = false
	fileInfo1, err := file1.Stat()
	if err != nil {
		panic(err)
	}
	fileInfo2, err := file2.Stat()
	if err != nil {
		panic(err)
	}
	file1.Close()
	file2.Close()
	if os.SameFile(fileInfo1, fileInfo2) {
		equal = true
	}
	encounters := 0
	for i := range img.Edges {
		for j := range img2.Edges {
			if img2.Edges[j] == img.Edges[i] {
				encounters++
				if encounters > 5000 {
					equal = true
					break
				}
			}
		}
	}
	return equal
}

func getImagePixels(img1 image.Image) Pixel {
	rows, columns := img1.Bounds().Dy(), img1.Bounds().Dx()
	pixelArray := Pixel{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 {
				color := img1.At(i, j)
				pixelArray.Edges = append(pixelArray.Edges, color)
			} else if i == rows-1 {
				color := img1.At(i, j)
				pixelArray.Edges = append(pixelArray.Edges, color)
			} else if i > 0 && i < rows-1 {
				if j == 0 || j == columns-1 {
					color := img1.At(i, j)
					pixelArray.Edges = append(pixelArray.Edges, color)
				}
			} else {
				continue
			}
		}
	}
	return pixelArray
}

func getOriginalImagePixel(img1 image.Image) Pixel {
	rows, columns := img1.Bounds().Dy(), img1.Bounds().Dx()
	pixelArray := Pixel{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			color := img1.At(i, j)
			pixelArray.Edges = append(pixelArray.Edges, color)

		}
	}
	return pixelArray
}

func (pi *Pixel) getRGBA(position int) (uint32, uint32, uint32, uint32) {
	r, g, b, a := pi.Edges[position].RGBA()
	return r, g, b, a
}
