package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
)

const (
	//画布大小
	width, height = 1920, 1080
	cells         = 100
	xyrange       = 30.0
	//x,y轴单位长度像素
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

//sin(30) cos(30)
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x,y float64) float64 {
	//到(0,0)距离
	r := math.Hypot(x, y)
	return math.Sin(r)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w)
}

func corner(i,j int) (float64, float64){

	//求出网格单元级i,j 订单坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度z
	z := f(x, y)

	sx := width  / 2 + (x-y) * cos30 * xyscale
	sy := height / 2 + (x-y) * sin30 * xyscale - z * zscale

	return sx, sy
}

func draw(w io.Writer){
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)

	for i := 0 ; i < cells; i ++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w,"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w,"</svg>")
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}