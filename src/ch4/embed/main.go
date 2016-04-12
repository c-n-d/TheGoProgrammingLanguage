/*
embed demonstates embedded structs and how they can be referenced using
the shorthand notation w.X => w.Circle.Point.X

$ go run src/ch4/embed/main.go
main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:20}
*/

package main

import "fmt"

type Point struct {
    X, Y int
}

type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}

func main() {
    w := Wheel{Circle{Point{8, 8}, 5}, 20}

    w = Wheel{
        Circle: Circle{
            Point: Point{X: 8, Y: 8},
            Radius: 5,
        },
        Spokes: 20, // NOTE: Trailing comma is required here (and after Radius)
    }

    fmt.Printf("%#v\n", w)

    w.X = 42

    fmt.Printf("%#v\n", w)
}
