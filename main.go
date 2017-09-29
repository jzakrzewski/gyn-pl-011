package main

import (
	"github.com/samuel/go-pcx/pcx"
	"image"
	"image/color"
	"io/ioutil"
	"os"
)

const WIDTH = 320
const HEIGHT = 240

func main() {
	if len(os.Args) < 2 {
		panic("And where are the files?")
	}

	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		panic(err)
	}

	result := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for _, f := range files {
		imgFile, err := os.Open(os.Args[1] + "/" + f.Name())
		if err != nil {
			panic(err)
		}

		img, err := pcx.Decode(imgFile)
		if err != nil {
			panic(err)
		}
		imgFile.Close()

	Loop:
		for i := 0; i < HEIGHT; i++ {
			for j := 0; j < WIDTH; j++ {
				c := img.At(j, i).(color.RGBA)
				if c.R > 0 || c.G > 0 || c.B > 0 {
					px := 4 * (i*WIDTH + j)
					result.Pix[px] = c.R
					result.Pix[px+1] = c.G
					result.Pix[px+2] = c.B //assume no alpha
					break Loop             // assume one relevant pixel per image
				}
			}
		}
	}

	out, err := os.Create("/tmp/out.pcx")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	pcx.Encode(out, result)
}
