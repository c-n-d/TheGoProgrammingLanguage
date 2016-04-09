/*
Mandelbrot emits a PNG image of the Mandelbrot fractal.

$ go run src/ch3/mandelbrot/main.go > src/ch3/mandelbrot/mandelbrot.png

For color:
$ go run src/ch3/mandelbrot/main.go -c > src/ch3/mandelbrot/mandelbrot.png
*/

package main

import (
    "flag"
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

// Exercise 3.5 - Add color to the Mandelbrot fractal. Otherwise black and white
var useColor = flag.Bool("c", false, "Use color when displaying")

func init () {
   flag.Parse()
}

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))

    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin

        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)

            // Image point (px, py) represents complex value z
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast   = 15

    var v complex128

    for n := uint8(0); n < iterations; n++ {
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            if *useColor {
                return color.RGBA{255 - contrast * n, contrast * n,  255 - contrast * n, 255}
            }

            return color.Gray{255 - contrast * n}
        }
    }

    return color.Black
}
