/*
ColoredPoint demonstrates struct embedding

$ go run src/ch6/coloredpoint/main.go
-- Init 1 --
{1 2} {4 6}
5
func(main.Point, main.Point) float64
{2 4}
func(*main.Point, float64)

-- Init 2 --
5
{2 2} {2 2}

-- Main --
5
10
*/

package main

import (
    "fmt"
    "image/color"
    "math"
)

type Point struct {
    X, Y float64
}

type ColoredPoint struct {
    Point
    Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
    dX := p.X-q.X
    dY := p.Y-q.Y
    return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
    fmt.Println("-- Main --")
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    var p = ColoredPoint{Point{1, 1}, red}
    var q = ColoredPoint{Point{5, 4}, blue}

    fmt.Println(p.Distance(q.Point))
    p.ScaleBy(2)
    q.ScaleBy(2)
    fmt.Println(p.Distance(q.Point))
}

func init() {
    fmt.Println("-- Init 1 --")
    p := Point{1, 2}
    q := Point{4, 6}

    fmt.Println(p, q)

    distance := Point.Distance
    fmt.Println(distance(p, q))
    fmt.Printf("%T\n", distance)

    scale := (*Point).ScaleBy
    scale(&p, 2)
    fmt.Println(p)
    fmt.Printf("%T\n\n", scale)
}

func init() {
    fmt.Println("-- Init 2 --")
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}

    type ColoredPoint struct {
        *Point
        Color color.RGBA
    }
    p := ColoredPoint{&Point{1, 1}, red}
    q := ColoredPoint{&Point{5, 4}, blue}

    fmt.Println(p.Distance(*q.Point))
    q.Point = p.Point
    p.ScaleBy(2)
    fmt.Println(*p.Point, *q.Point)
    fmt.Println()
}
