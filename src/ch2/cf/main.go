/*
CF reads numbers from the command line args or from stdin and converts them between
the types in the tempconv and lenghtconv packages.
 
$ go run src/ch2/cf/main.go 123 456
123°F = 50.55555555555556°C, 123°F = 323.7055555555556°K, 123°C = 253.4°F, 123°C = 396.15°K, 123°K = -238.27°F, 123°K = -150.14999999999998°C
123 feet = 37.4904 meter, 123 meter = 403.54332 feet

456°F = 235.55555555555554°C, 456°F = 508.7055555555556°K, 456°C = 852.8°F, 456°C = 729.15°K, 456°K = 361.12999999999994°F, 456°K = 182.85000000000002°C
456 feet = 138.9888 meter, 456 meter = 1496.06304 feet
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"

    "ch2/tempconv"
    "ch2/ex_2_2"
)

func main() {
    if len(os.Args[1:]) > 0 {
        for _, arg := range os.Args[1:] {
            reportConversions(arg)
        } 
    } else {
        var scanner = bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
            reportConversions(scanner.Text())
        }
    }
}

func reportConversions(arg string) {
    t, err := strconv.ParseFloat(arg, 64)

    if err != nil {
        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        os.Exit(1)
    }

    f := tempconv.Fahrenheit(t)
    c := tempconv.Celsius(t)
    k := tempconv.Kelvin(t)

    ft := lengthconv.Feet(t)
    m := lengthconv.Meter(t)

    fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
        f, tempconv.FtoC(f),
        f, tempconv.FtoK(f),
        c, tempconv.CtoF(c),
        c, tempconv.CtoK(c),
        k, tempconv.KtoF(k),
        k, tempconv.KtoC(k))

        fmt.Printf("%s = %s, %s = %s\n\n",
        ft, lengthconv.FeetToMeter(ft),
        m, lengthconv.MeterToFeet(m))
}
