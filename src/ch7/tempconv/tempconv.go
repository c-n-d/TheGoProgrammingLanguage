/*
Package tempconv provides Celsisus, Fahrenheit and Kelvin temp conversions
*/

package tempconv7

import (
    "flag"
    "fmt"

    "ch2/tempconv"
)

// *celsiusFlag statisfies the flag.Value interface.
type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
    var unit string
    var value float64
    fmt.Sscanf(s, "%f%s", &value, &unit) // no error checking necessary

    switch unit {
        case "C", "°C":
            f.Celsius = tempconv.Celsius(value)
            return nil
        case "F", "°F":
            f.Celsius = tempconv.FtoC(tempconv.Fahrenheit(value))
            return nil
        // Exercise 7.6 - Add support for Kelvin
        case "K", "°K":
            f.Celsius = tempconv.KtoC(tempconv.Kelvin(value))
            return nil
    }
    return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
    f := celsiusFlag{value}
    flag.CommandLine.Var(&f, name, usage)
    return &f.Celsius
}
