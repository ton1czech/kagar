package main

import (
    "fmt"

    // "github.com/gocolly/colly"
)

type Car struct {
    Model    string `json:"model"`
    Price    string `json:"price"`
    Location string `json:"location"`
    Url      string `json:"url"`
}

func main() {
    fmt.Println("Hello")
}