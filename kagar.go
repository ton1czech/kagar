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
        "https://www.theparking.eu/used-cars/toyota-soarer.html#!/used-cars/toyota-soarer.html%3Fid_boite%3D2%26tri%3Dprix_croissant",
    }

    var cars []Car

    c.OnError(func(_ *colly.Response, err error) {
        panic(err)
    })

    c.OnHTML("div#lists ul#resultats li section.clearfix.complete-holder", func(h *colly.HTMLElement) {
        if h.ChildText("sponsor") != "sponsored" {
            car := Car{
                Manufacturer: strings.TrimSpace(h.ChildText("span.title-block.brand")),
                Model: strings.TrimSpace(h.ChildText("span.sub-title.title-block")),
                Detail: strings.TrimSpace(h.ChildText("span.nowrap")),
                Price: strings.TrimSpace(h.ChildText("p.prix")),
            }

            cars = append(cars, car)

            fmt.Println(cars)
        }

    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println(r.URL.String())
    })

    c.Visit(links[1])

    content, err := json.MarshalIndent(cars, "", "  ")

    if err != nil {
        fmt.Println(err.Error())
    }

    os.WriteFile("cars.json", content, 0644)
}