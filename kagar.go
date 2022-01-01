package main

import (
	"fmt"
	"strings"

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
        "https://www.theparking.eu/used-cars/toyota-corolla-e80-ae86.html#!/used-cars/toyota-corolla-e80-ae86.html%3Fid_pays%3D1%7C5%26tri%3Dprix_croissant",
        "https://www.theparking.eu/used-cars/toyota-soarer.html#!/used-cars/toyota-soarer.html%3Ftri%3Dprix_croissant",
    }

    scrapeWeb(links)
}

func scrapeWeb(links []string) {
    cars := []Car{}

    c := colly.NewCollector(
        colly.AllowedDomains("www.theparking.eu", "theparking.eu"),
    )

    c.OnError(func(_ *colly.Response, err error) {
        panic(err)
    })

    c.OnHTML("div#lists ul#resultats li", func(h *colly.HTMLElement) {
        doc := h.DOM

        car := Car{}
        car.Manufacturer = strings.Trim(doc.Find("span.title-block.brand").Text(), " ")
        cars = append(cars, car)
    })

    c.Visit(links[0])

    fmt.Println(cars)
}