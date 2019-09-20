package ch2

import "fmt"

//Celsius temperature
type Celsius float64

//Fahrenheit temperature
type Fahrenheit float64

/* Even though type Celsius and Fahrenheit have the same underlying type, float64,
   they are not the same type, so they cannot be compared or combined in arithmetic
   expressions;an explicit conversion like Celsius(t) or Fahrenheit(t) is required
   to convert from a float64.Celsius(t) or Fahrenheit(t) is conversion not function call.
*/
const (
	//AbsoluteZero temperature for Celsius
	AbsoluteZero Celsius = -273.15
	//FreezingC temperature for Celsius
	FreezingC Celsius = 0
	//BoilingC temperature for Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

/*
    fmt.Println(ch2.BoilingC - ch2.FreezingC)
    fmt.Println(ch2.CToF(ch2.BoilingC - ch2.FreezingC))
	mt.Println(ch2.CToF(ch2.Celsius(float64(float32(int(37))))))
*/
