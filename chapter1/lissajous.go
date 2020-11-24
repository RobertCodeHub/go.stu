package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

const (
	blackIndex =1
)

var palette =[]color.Color{color.White, color.Black}

func lissajous(){

	const (
		cycle = 100
		res = 0.001
		size = 500
		frames = 1000
		delay = 8
	)

	phase := 0.00;
	freq := rand.Float64() * 3.00;
	anim := gif.GIF{LoopCount: frames}

	for i := 0.00; i < frames; i++{
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		lissajiousImage := image.NewPaletted(rect, palette)

		for t := 0.00; t < 2*cycle*math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			lissajiousImage.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 1.0;
		anim.Delay = append(anim.Delay, delay)
		anim. Image = append(anim. Image, lissajiousImage)
	}

	//创建图片文件
	mandebrot, err := os.Create("E:\\tmandelbrot2.png")

	//将内容输出到图片
	err = gif.EncodeAll(mandebrot, &anim)

	if(err != nil){
		fmt.Println(err)
	}
}

func main()  {
	lissajous()
}