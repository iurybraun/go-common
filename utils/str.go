/*
 * Copyright © 2016-2018 Iury Braun
 * Copyright © 2017-2018 Weyboo
 */

package utils

import (
    "log"
    "strconv"
    "strings"
)

func StrToBool(str string) bool {
    if str == "on" || str == "true" {
        return true
    }
    return false
}

func StrToFloat(str string) float64 {
    v1 := strings.Replace(str, ".", "", -1)
    v2 := strings.Replace(v1, ",", ".", -1)
    
    if v2 != "" {
        // "var float float32" up here somewhere
        value, err := strconv.ParseFloat(v2, 64)
        if err != nil {
            log.Println(err)
            return 0.0
        }
        return value
    } else {
        return 0.0
    }
}

func StrToInt(str string) int {
    //value, err := strconv.ParseInt(str, 10, 64)
    value, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    return value
}




func StrBetween(value string, a string, b string) string {
    // Get substring between two strings.
    posFirst := strings.Index(value, a)
    if posFirst == -1 {
        return ""
    }
    posLast := strings.Index(value, b)
    if posLast == -1 {
        return ""
    }
    posFirstAdjusted := posFirst + len(a)
    if posFirstAdjusted >= posLast {
        return ""
    }
    return value[posFirstAdjusted:posLast]
}

func StrBefore(value string, a string) string {
    // Get substring before a string.
    pos := strings.Index(value, a)
    if pos == -1 {
        return ""
    }
    return value[0:pos]
}

func StrAfter(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

func StrJustNumber(s string) string {
    return strings.Map(
        func(r rune) rune {
            if r == '0' || r == '1' || r == '2' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9' {
                return r
            }
            return -1
        },
        s,
    )
}

/* regexp
func removeDupliciedSpace(input string) string {
	
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
    output := re_leadclose_whtsp.ReplaceAllString(input, "")
    return re_inside_whtsp.ReplaceAllString(output, " ")
	
}*/

func StrStandardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func RandString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
        bytes[i] = byte(RandInt(48, 57))
    }
    return string(bytes)
}
