package main

import(
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

func mandelbrot(z complex128) color.Color {

	const contrast   = 15
	const iterations = 200

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast * n}
		}
	}

	//随机数0-255之间
	var r uint8 = uint8(rand.Intn(255))
	var g uint8 = uint8(rand.Intn(255))
	var b uint8 = uint8(rand.Intn(255))
	var a uint8 = uint8(rand.Intn(255))

	//使用这个生成全彩图片
	return color.NRGBA{r, g, b, a}
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	//创建图片文件
	mandebrot, err := os.Create("E:\\tmandelbrot2.png")

	//将内容输出到图片
	err = png.Encode(mandebrot, img)

	if(err != nil){
		fmt.Println(err)
	}
}