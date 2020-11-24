package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	blackIndexTwo = 6
)

var paletteTwo =[]color.Color{color.White, color.Black, color.RGBA{150,150,150,150}, color.RGBA{100,100,100,100}, color.RGBA{200,200,200,200}}

func lissajousTwo(w http.ResponseWriter, r *http.Request){

	parameter  := r.URL.Query().Get("cycle")
	if parameter  == ""{
		fmt.Fprintln(w, "cycle is nli")
		return
	}

	//将获取到的字符串参数转为float64
	cycle, _ := strconv.ParseFloat(parameter , 64)

	const (
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

		lissajiousImage := image.NewPaletted(rect, paletteTwo)

		for t := 0.00; t < 2*cycle*math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			lissajiousImage.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndexTwo)
		}

		phase += 1.0;
		anim.Delay = append(anim.Delay, delay)
		anim. Image = append(anim. Image, lissajiousImage)
	}

	//将内容输出到浏览器
	gif.EncodeAll(w, &anim)
}

//练习1.12
func main() {
	http.HandleFunc("/lissajous", lissajousTwo)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}