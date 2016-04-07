/*
Tempconv performs Celsius and Fahrenheit temperature conversions

Adds String() method implementations for Celsius and Fahrenheit types
*/

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func FtoC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func CtoF(c Celsius) Fahrenheit {
    return Fahrenheit (c * 9 / 5 + 32)
}

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}
