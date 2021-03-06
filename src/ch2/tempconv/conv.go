/*
Conv.go adds Fahrenheit/Celsius/Kelvin temperature conversions to the tempconv package
*/

package tempconv

func FtoC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func CtoF(c Celsius) Fahrenheit {
    return Fahrenheit (c * 9 / 5 + 32)
}

/*
Exercise 2.1 - Adding Kelvin conversions to tempconv package
*/
func KtoC(k Kelvin) Celsius {
    return Celsius(k - 273.15)
}

func KtoF(k Kelvin) Fahrenheit {
    return Fahrenheit(k * 9 / 5 - 459.67)
}

func FtoK(f Fahrenheit) Kelvin {
    return Kelvin((f +  459.67) * 5 / 9)
}

func CtoK(c Celsius) Kelvin {
    return Kelvin(c + 273.15)
}
