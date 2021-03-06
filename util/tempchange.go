package util

//自定义类型Celsius
type Celsius float64

//自定义类型Fahrenheit
type Fahrenheit float64

const(
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF (c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

func FToC (f Fahrenheit) Celsius {
	return Celsius((f - 32)*5/9)
}
