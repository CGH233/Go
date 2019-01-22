package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0xbe, 0x13, 0x13, 0xff}, color.RGBA{0xe9, 0xc0, 0x2b, 0xff,}, color.RGBA{0x28,0x1d,0xf7,0xff}}

const (
	blackIndex = 0
	redIndex = 1
	yellowIndex = 2
	blueIndex = 3
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request)	{
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0;i < nframes; i++ {
		rect := image.Rect(0, 0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			var n int = i%4
			switch n {
			case 0:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			case 1:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
			case 2:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), yellowIndex)
			case 3:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}