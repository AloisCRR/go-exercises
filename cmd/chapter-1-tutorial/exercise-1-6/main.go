/*
	Modify the Lissajous program to produce images in multiple colors by adding
	more values to palette and then displaying them by changing the third argument
	of SetColorIndex in some interesting way.
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 0xFF, A: 0xff},
	color.RGBA{G: 0xFF, A: 0xff},
	color.RGBA{B: 0xFF, A: 0xff},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions
		res    = 0.001 // angular resolution
		size   = 100   // image canvas covers [-size..+size]
		frames = 64    // number of animation frames
		delay  = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	anim := gif.GIF{LoopCount: frames}

	phase := 0.0 // phase difference

	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {

			// colorIndex := uint8(int(t) % 4)

			colorIndex := rand.Intn(3) + 1

			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}

		phase += 0.1

		anim.Delay = append(anim.Delay, delay)

		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
