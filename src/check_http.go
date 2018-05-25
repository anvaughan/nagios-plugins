package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func main() {
    var rc int = 0
    var url string = os.Args[1]
    start := time.Now()
    resp, err := http.Get(url)
    end := time.Now()
    elapsed := end.Sub(start)
    if err != nil {
        fmt.Printf("??? - %s - took %s\n", err, elapsed)
        os.Exit(2)
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        rc = 0
    } else if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
        rc = 1
    } else if resp.StatusCode >= 500 {
        rc = 2
    } else {
        rc = 2
    }
    fmt.Printf("%d - %s - took %s\n", resp.StatusCode, http.StatusText(resp.StatusCode), elapsed)
    os.Exit(rc)
}
