//Package ch2 separate type declaration,type method from utility functions
package ch2

//CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
