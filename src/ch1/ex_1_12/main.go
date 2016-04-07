/*
Exercise 1.12

$ go run src/ch1/ex_1_12/main.go

> http://localhost:8888
> http://localhost:8888/?cycles=9
> http://localhost:8888/asGif
*/
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

func main() {
    lissajousHandler := func (w http.ResponseWriter, r *http.Request) {
        lissajous(w, r)
    }

    // Register the request handler for the root route
    http.HandleFunc("/", lissajousHandler)

    http.HandleFunc("/asGif", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<html> <body> <img src=\"/lissajous\" /> </body> </html>")
    })

    log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func lissajous(out io.Writer, r *http.Request) {

    const (
        cyclesDefault  = 5     // number of complete x oscillator revolutions
        res            = 0.001 // angular revolution
        size           = 100   // image canvas covers [-size..+size]
        nframes        = 64    // number of animation frames
        delay          = 8     // delay between frames in 10ms units
    )

    cycles, err := strconv.Atoi(r.FormValue("cycles"))

    if err != nil || cycles <= 0 {
        cycles = cyclesDefault
        fmt.Printf("Using default cycles of %d. Err if present is: %v\n", cycles, err)
    } else {
        fmt.Printf("Using cycles from form values: %d.\n", cycles)
    }

    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)

        for t := 0.0; t < float64(cycles) * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)

            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
        }

        phase += 0.1

        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
