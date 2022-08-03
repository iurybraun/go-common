package utils

import (
    "strconv"
)

func FloatToStr(n float64, precision int) string {
    return strconv.FormatFloat(n, 'f', precision, 64)
}
