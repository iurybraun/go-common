package utils

import (
    "bytes"
    "io"
    "io/ioutil"
    "mime"
    "mime/multipart"
    "net/http"
    "strings"
)

// Get form values without invalidating the request body in case the data is multiform
func GetFormValues(request *http.Request, keys []string) []string {
    
    var values []string
    mediaType, params, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
    if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
        for i := range keys {
            values = append(values, request.FormValue(keys[i]))
        }
    } else { // multi form
        buf, _ := ioutil.ReadAll(request.Body)
        origBody := ioutil.NopCloser(bytes.NewBuffer(buf))
        var rdr = multipart.NewReader(bytes.NewBuffer(buf), params["boundary"])
        for len(values) < len(keys) {
            part, err_part := rdr.NextPart()
            if err_part == io.EOF {
                break
            }
            for i := range keys {
                if part.FormName() == keys[i] {
                    buf := new(bytes.Buffer)
                    buf.ReadFrom(part)
                    values = append(values, buf.String())
                }
            }
        }
        request.Body = origBody
    }
    if len(values) == len(keys) {
        return values
    } else {
        return nil
    }
    
}

// Get form value without invalidating the request body in case the data is multiform
func GetFormValue(request *http.Request, key string) string {
    
    if result := GetFormValues(request, []string{key}); len(result) == 1 {
        return result[0]
    } else {
        return ""
    }
    
}
