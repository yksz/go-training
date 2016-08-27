package equal

import (
	"math"
	"reflect"
)

func equal(x, y reflect.Value) bool {
	xf := toFloat64(x)
	yf := toFloat64(y)
	diff := math.Abs(xf - yf)
	return diff < 1.0e-10
}

func toFloat64(x reflect.Value) float64 {
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return float64(x.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return float64(x.Uint())

	case reflect.Float32, reflect.Float64:
		return x.Float()
	}
	panic("not a number")
}

func Equal(x, y interface{}) bool {
	return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}
