/*
Exercise 5.6 - Modify the corner function in surface to use named results and bare return statements

$ go run src/ch5/ex_5_6/main.go > src/ch5/ex_5_6/surface.html
$ open src/ch5/ex_5_6/surface.html

To run this example as a web server:

$ go run src/ch5/ex_5_6/main.go -s
$ open http://localhost:8888/
*/

package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "math"
    "net/http"
    "os"
)

const (
    width, height = 600, 300            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis range (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angel of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var asServer = flag.Bool("s", false, "Starts a server on :8888 to serve the SVG")

func init() {
    flag.Parse()
    http.HandleFunc("/", handler)
}

func main() {
    if *asServer {
        log.Fatal(http.ListenAndServe("localhost:8888", nil))
    }

    generateSurface(os.Stdout)
}

// Exercise 3.4 - Web server to construct the surface and write the SVG to the client
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")

    generateSurface(w)
}

func generateSurface(writer io.Writer) {
    // Don't set stroke for the entire SVG element
    fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height)

    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)

            // Exercise 3.1 - Skip invalid polygons
            if !validPolygon(ax, ay, bx, by, cx, cy, dx, dy) {
                   continue
               }

            // Exercise 3.3 - Determine red / blue for the polygon. Only using the (i, j) corner for now
            var color = determineColor(i, j)

            fmt.Fprintf(writer, "<polygon style='stroke:%s' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
              color, ax, ay, bx, by, cx, cy, dx, dy)
        }
    }

    fmt.Fprintf(writer, "</svg>");
}

func corner(i, j int) (sx, sy float64) {
    // Find point (x, y) at corner of cell (i, j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute the surface height z.
    z := f(x, y)

    // Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
    sx = width / 2 + (x - y) * cos30 * xyscale
    sy = height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0, 0)
    return math.Sin(r) / r
}

// An invalid polygon is defined to be a polygon that
// consists of a non-finite coordinate
func validPolygon(fs ...float64) bool {
    var valid = true

    for _, f := range fs {
        valid = valid && isFinite(f)
    }

    return valid
}

// A finite number must be non-infinite and must be an otherwise valid number
func isFinite(f float64) bool {
    return !( math.IsInf(f, 0) || math.IsNaN(f) )
}

// Returns red/blue based on whether the Z dimension in 3d space is
// above or below the Z axis (Z=0)
func determineColor(i, j int) string {
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    z := f(x, y)

    if z > 0 {
        return "#ff0000"
    }

    return "#0000ff"
}
