package main

// Based on:
//  - Scratch container: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
//  - Non-privileged user: https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "log"
)

type ContentLengthResponse struct {
    Google int `json:"google"`
}

func ContentLengthHandler(res http.ResponseWriter, req *http.Request) {
    resp, err := http.Get("https://google.com")
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Printf("Length of content received from google.com: %d\n", len(body))
    fmt.Printf("Running as user ID: %d\n", os.Getuid())
    contentLengthResponse := &ContentLengthResponse{
        Google : len(body),
    }
    data, _ := json.MarshalIndent(contentLengthResponse, "", "  ")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {

    port := "7000"
    if p := os.Getenv("PORT"); p != "" {
        port = p
    }

    log.Printf("Listening on this host: http://localhost:%v\n", port)

    http.HandleFunc("/google", ContentLengthHandler)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("Unable to listen on :7000: ", err)
    }
}
