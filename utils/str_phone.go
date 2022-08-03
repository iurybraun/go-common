/*
 * Copyright © 2016-2019 Iury Braun
 * Copyright © 2017-2019 Weyboo
 */

package utils

import (
    //"strconv"
    //"strings"
)

func DecodeMobilePhone(s string) (string, string){
    strPhone := StrJustNumber(s)
    
    if len(strPhone) > 0 {
        s1 := strPhone[0:1]
        s2 := strPhone[1:2]
        s3 := strPhone[2:3]
        
        number  := strPhone[len(strPhone)-9:]
        
        area := ""
        if s1 == "0" {
            area = s1 + s2 + s3
        } else {
            area = s1 + s2
        }
        
        return area, number
    } else {
        return "", ""
    }
}
