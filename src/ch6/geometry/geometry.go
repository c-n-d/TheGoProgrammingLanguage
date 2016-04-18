/*
Package geometry defines simple types for plane geometry.
*/

package geometry

import "math"

type Point struct {
    X, Y float64
}

// traditional functions
func Distance(p, q Point) float64 {
    return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (path Path) Distance() float64{
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum =+ path[i-1].Distance(path[i])
        }
    }
    return sum
}
