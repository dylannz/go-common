package extmath

import "math"

// Round rounds a float64 value n to the required significant figures sf.
// Stolen shamelessly from: https://gist.github.com/pelegm/c48cff315cd223f7cf7b
func Round(val float64, sf int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(sf))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(.5, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
