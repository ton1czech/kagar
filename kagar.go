package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Car struct {
    Manufacturer string
    Model        string
    Detail       string
    Price        string
    Location     string
    Fuel         string
    Year         string
    Transmission string
    Mileage      string
}

func main() {
    links := []string{
        "https://www.theparking.eu/used-cars/toyota-corolla-e80-ae86.html#!/used-cars/toyota-corolla-e80-ae86.html%3Ftri%3Dprix_croissant",
        "https://www.theparking.eu/used-cars/toyota-soarer.html#!/used-cars/toyota-soarer.html%3Ftri%3Dprix_croissant",
    }

    scrapeWeb(links)
}

func scrapeWeb(links []string) {
    c := colly.NewCollector(
        colly.AllowedDomains("www.theparking.eu", "theparking.eu"),
    )

    c.OnError(func(_ *colly.Response, err error) {
        panic(err)
    })

    c.OnHTML("ul#resultats", func(h *colly.HTMLElement) {
        fmt.Println(h)
    })

    for _, link := range links {
        c.Visit(link)
    }
}