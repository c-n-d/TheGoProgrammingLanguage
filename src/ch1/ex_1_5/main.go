/*
Exercise 1.5: Lissajous generates GIF animations of random Lissajous figures

The palette for the animation is {black, green}

https://en.wikipedia.org/wiki/Lissajous_curve

$ go run src/ch1/ex_1_5/main.go > 1.5.out.gif
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
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xFF},
                            color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
    blackIndex = 0 // first color in the palette
    greenIndex = 1 // second color in the palette
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {

    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular revolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 2.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)

        for t := 0.0; t < cycles * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)

            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex)
        }

        phase += 0.1

        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
