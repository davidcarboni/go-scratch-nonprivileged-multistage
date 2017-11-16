package main

// Based on: 
//  - Scratch container: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
//  - Non-privileged user: https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get("https://google.com")
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Println(len(body))
    fmt.Println("Ohai! :D")

    fmt.Printf("Running as user ID: %d\n", os.Getuid())
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
