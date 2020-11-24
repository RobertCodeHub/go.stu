package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// creating image at *2 resolution.
	img2 := image.NewRGBA(image.Rect(0, 0, width * 2, height * 2))
	for py := 0; py < height * 2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < width * 2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin
			z := complex(x, y)
			img2.SetRGBA(px, py, mandelbrotTwo(z))
		}
	}
	// decreasing image size.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py = py + 2 {
		for px := 0; px < width*2; px = px + 2 {
			var c = []color.RGBA{
				img2.RGBAAt(px, py),
				img2.RGBAAt(px+1, py),
				img2.RGBAAt(px, py+1),
				img2.RGBAAt(px+1, py+1),
			}
			var pr, pg, pb, pa int
			for n := 0; n < 4; n++ {
				pr += int(c[n].R)
				pg += int(c[n].G)
				pb += int(c[n].B)
				pa += int(c[n].A)
			}
			img.SetRGBA(px/2, py/2, color.RGBA{uint8(pr / 4), uint8(pg / 4), uint8(pb / 4), uint8(pa / 4)})
		}
	}

	//创建图片文件
	mandebrot, err := os.Create("E:\\tmandelbrot.png")

	//将内容输出到图片
	err = png.Encode(mandebrot, img)

	if(err != nil){
		fmt.Println(err)
	}
}

func mandelbrotTwo(z complex128) color.RGBA {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast * (n + 2), contrast * (n + 1), contrast * n, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}