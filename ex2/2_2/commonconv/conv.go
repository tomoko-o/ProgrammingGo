package commonconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(5/9*(f-32) + 273.15) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit(5/9*(k-32) + 273.15) }

func FtToM(ft Feet) Meter { return Meter(ft / 3.2808) }
func MToFt(m Meter) Feet  { return Feet(m * 3.2808) }

func KgToP(kg Kilogram) Pound { return Pound(kg * 0.454) }
func PToKg(p Pound) Kilogram  { return Kilogram(p / 0.454) }
