/*
Mandelbrot emits a PNG image of the Mandelbrot fractal.

Determines the color of the pixel based of the supersample of four pixels

$ go run src/ch3/ex_3_6/main.go > src/ch3/ex_3_6/mandelbrot_3_6.png
*/

package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

func main() {
    const (
        xmin, ymin, xmax, ymax  = -2, -2, +2, +2
        width, height           = 1024, 1024
        // Doubling the width, height quadruples the number of pixels
        sampleSize              = 2
        superWidth, superHeight = sampleSize * width, sampleSize * height
    )

    supersampleImg := make([][]color.Color, superHeight)

    for i := range supersampleImg {
        supersampleImg[i] = make([]color.Color, superWidth)
    }

    img := image.NewRGBA(image.Rect(0, 0, width, height))

    for py := 0; py < superHeight; py++ {
        y := float64(py) / superHeight * (ymax - ymin) + ymin

        for px := 0; px < superWidth; px++ {
            x := float64(px) / superWidth * (xmax - xmin) + xmin
            z := complex(x, y)

           supersampleImg[px][py] = mandelbrot(z)
        }
    }

    for py := 0; py < height; py++ {
        realY:= py*sampleSize

        for px := 0; px < width; px++ {
            realX := px*sampleSize

            // Average the colors from the four supersampled pixels
            avgColor := averageColor(supersampleImg[realX][realY], supersampleImg[realX+1][realY],
                                     supersampleImg[realX][realY+1], supersampleImg[realX+1][realY+1])

            img.Set(px, py, avgColor)
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
            return color.RGBA{255 - contrast * n, contrast * n,  255 - contrast * n, 255}
        }
    }

    return color.Black
}

// Averages the RGBA elements of a list of Colors
func averageColor(colors...color.Color) color.Color {
    var red, green, blue, alpha uint32
    size := uint32(len(colors))

    for _, c := range colors {
        cr, cg, cb, ca := c.RGBA()

        red += cr
        green += cg
        blue += cb
        alpha += ca
    }

    return color.RGBA{uint8(red / size), uint8(green / size), uint8(blue / size), uint8(alpha / size)}
}
