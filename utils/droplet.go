/*
 * Copyright © 2016-2018 Iury Braun
 * Copyright © 2017-2018 Weyboo
 */

package utils

import (
    //"log"
    //"strconv"
    //"strings"
    
    "net/http"
)

func getDroplet(r *http.Request) string {
    var droplet string
    
    // map req.Header
    map_reqheader := make(map[string]interface{})
    
    for key, values := range r.Header {  // range over map
        for _, value := range values {   // range over []string
            map_reqheader[key] = value
        }
    }
    
    SourceData := map_reqheader["X-Req-Source-Data"]
    if SourceData != nil {
        droplet = map_reqheader["X-Req-Source-Data"].(string)
    }
    
    return droplet
}
