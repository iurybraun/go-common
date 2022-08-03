/*
 * Copyright © 2016-2021 Iury Braun
 * Copyright © 2017-2021 Neo7i
 */

package utils

import (
    "fmt"
    "os"
    "net/http"
    "io"
)

func DownloadFile(filepath string, url string) (err error) {
    // Create the file
    out, err := os.Create(filepath)
    if err != nil  {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check server response
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    // Writer the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil  {
        return err
    }

    return nil
}
