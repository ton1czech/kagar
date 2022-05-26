package main

import (
	"encoding/json"
	"fmt"
	"os"
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
    c := colly.NewCollector(
        colly.AllowedDomains("www.theparking.eu", "theparking.eu"),
    )

    links := []string{
        "https://www.theparking.eu/used-cars/toyota-corolla-e80-ae86.html#!/used-cars/toyota-corolla-e80-ae86.html%3Fid_pays%3D1%7C5%26tri%3Dprix_croissant",
        "https://www.theparking.eu/used-cars/toyota-soarer.html#!/used-cars/toyota-soarer.html%3Ftri%3Dprix_croissant",
    }

    var cars []Car

    c.OnError(func(_ *colly.Response, err error) {
        panic(err)
    })

    c.OnHTML("div#lists ul#resultats li section.clearfix.complete-holder", func(h *colly.HTMLElement) {
        car := Car{
            Manufacturer: strings.TrimSpace(h.ChildText("span.title-block.brand")),
            // Model: strings.Trim(doc.Find("span.sub-title.title-block").Text(), " "),
            // Detail: strings.Trim(doc.Find("span.nowrap").Text(), " "),
            // Price: strings.Trim(doc.Find("p.prix").Text(), " "),
            // Location: strings.Trim(doc.Find("div.location > span.upper").Text(), " "),
        }
        fmt.Println(car.Manufacturer)

        cars = append(cars, car)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println(r.URL.String())
    })

    c.Visit(links[0])

    content, err := json.MarshalIndent(cars, "", "  ")

    if err != nil {
        fmt.Println(err.Error())
    }

    os.WriteFile("cars.json", content, 0644)
}