/*
Exercise 8.5 - Take a CPU bound sequential program and execute its main loop in parallel,
               how much faster does it run on a multiprocessor machine? Optimal number of goroutines to use?

$ go build ch8/ex_8_5
$ go build ch7/surface
$ ex_8_5 &
$ surface &

$ open "http://localhost:8001/plot?expr=sin(-x)*pow(1.5,-r)"
$ open "http://localhost:8001/plot?expr=pow(2, sin(y))*pow(2,sin(x))/12"
$ open "http://localhost:8001/plot?expr=sin(x*y/10)/10"

$ open "http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)"
$ open "http://localhost:8000/plot?expr=pow(2, sin(y))*pow(2,sin(x))/12"
$ open "http://localhost:8000/plot?expr=sin(x*y/10)/10"
*/

package main

import (
    "fmt"
    "io"
    "log"
    "math"
    "net/http"
    "sync"

    "ch7/eval"
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

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
    // Find point (x, y) at corner of cell (i, j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute the surface height z.
    z := f(x, y)

    // Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
    sx := width / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return sx, sy
}

func surface(writer io.Writer, f func(x, y float64) float64) {
    svgChan := make(chan string, cells*cells)
    fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height)

    var wg sync.WaitGroup
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            wg.Add(1)
            go func(i, j int) {
                defer wg.Done()
                ax, ay := corner(f, i+1, j)
                bx, by := corner(f, i, j)
                cx, cy := corner(f, i, j+1)
                dx, dy := corner(f, i+1, j+1)
                svgChan <- fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
                    ax, ay, bx, by, cx, cy, dx, dy)
            }(i, j)
        }
    }

    go func() {
        wg.Wait()
        close(svgChan)
    }()

    for svg := range svgChan {
        fmt.Fprintf(writer, svg);
    }
    fmt.Fprintf(writer, "</svg>");
}

func parseAndCheck(s string) (eval.Expr, error) {
    if s == "" {
        return nil, fmt.Errorf("empty expression")
    }
    expr, err := eval.Parse(s)
    if err != nil {
        return nil, err
    }
    vars := make(map[eval.Var]bool)
    if err := expr.Check(vars); err != nil {
        return nil, err
    }
    for v := range vars {
        if v != "x" && v != "y" && v != "r" {
            return nil, fmt.Errorf("undefined variable: %s", v)
        }
    }
    return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    expr, err := parseAndCheck(r.Form.Get("expr"))
    if err != nil {
        http.Error(w, "bad expr: " + err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "image/svg+xml")
    surface(w , func(x, y float64) float64 {
            r := math.Hypot(x, y) // distance from (0, 0)
            return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
        })
}

func main() {
    http.HandleFunc("/plot", plot)
    log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
