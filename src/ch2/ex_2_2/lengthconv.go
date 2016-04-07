/*
Exercise 2.2 - Add additional conversions for feet & meters
*/

package lengthconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string {
    return fmt.Sprintf("%g feet", f)
}

func (m Meter) String() string {
    return fmt.Sprintf("%g meter", m)
}

func FeetToMeter(f Feet) Meter {
    return Meter(f * 0.3048)
}

func MeterToFeet(m Meter) Feet {
    return Feet(m * 3.28084)
}