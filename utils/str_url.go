/*
 * Copyright © 2016-2019 Iury Braun
 * Copyright © 2017-2019 Weyboo
 */

package utils

import (
    "log"
    "strings"
    "net/url"
)

func StrExtractDomain(url_str string) string {
    u, err := url.Parse(url_str)
    if err != nil {
        log.Fatal(err)
    }
    
    parts := strings.Split(u.Hostname(), ".")
    return parts[len(parts)-2] + "." + parts[len(parts)-1]
}
