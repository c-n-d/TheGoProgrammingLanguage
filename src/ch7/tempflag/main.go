/*
Tempflag prints the value of -temp in Celsius

Exercise 7.6 - Add support for Kelvin - see ch7/tempconv/tempconv.go

Exercise 7.7 - The help message contains '°C' because the default is converted into a Celsius
               type. The String() for Celsius prints '°C'.

$ go run src/ch7/tempflag/main.go -temp 40F
4.444444444444445°C
$ go run src/ch7/tempflag/main.go -temp 30C
30°C
$ go run src/ch7/tempflag/main.go -temp 271.3K
-1.849999999999966°C
*/

package main

import (
    "flag"
    "fmt"

    "ch7/tempconv"
)

var temp = tempconv7.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
    flag.Parse()
    fmt.Println(*temp)
}
