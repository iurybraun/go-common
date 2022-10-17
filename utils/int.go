package utils

import (
    "strconv"
    "math/rand"
)

func IntToStr(n int32) string {
    //return strconv.Itoa(n)
    return strconv.FormatInt(int64(n), 10)
}

// USE: PercentageChange(100, 130)
func PercentageChange(old, new int32) (delta float64) {
    diff := float64(new - old)
    delta = (diff / float64(old)) * 100
    return
}

/* !ERROR
func DiscountPerc(value int32, percent float64) (delta int32) {
    intVal := float64(value)
    discountValue := int32(intVal/100 * percent)
    
    delta = value - discountValue
	return
}*/

func DiscountPerc(value int32, percent float64) (delta int32) {
	intVal := float64(value)
	discountValue := intVal / 100.0 * percent

	delta = int32(intVal - discountValue)
	return
}

/*func CalcValueFromPerc(value int, percent float64) (delta int32) {
    delta = int32((float64(value / 100)) * percent)
    return
}*/

func CalcValueFromPerc(value int, percent float64) (delta int32) {
    delta = int32(float64(value)/100 * percent)
    return
}

func RandInt(min int32, max int32) int32 {
    return min + rand.Int31n(max-min)
}

func Diff(a, b int32) int32 {
   if a < b {
      return b - a
   }
   return a - b
}
