package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Database struct {
    NewResults []Car `json:"newResults"`
    OldResults []Car `json:"oldResults"`
    Links      []Link   `json:"links"`
}

type Car struct {
    Model    string `json:"model"`
    Price    string `json:"price"`
    Location string `json:"location"`
    Url      string `json:"url"`
}

type Link struct {
    Name string `json:"name"`
    Url string `json:"link"`
}

func main() {
    cars := make([]Car, 0)

    c := colly.NewCollector(
        colly.AllowedDomains("www.theparking.eu"),
    )

    c.OnError(func(_ *colly.Response, err error) {
        panic(err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println(r.Request.URL)
    })

    c.Visit("")

    fmt.Println(cars)
    links := getUrls()

    for i := range links {
        fmt.Println(links[i].Url)
    }
}

func getUrls() (links []Link) {
    file, err := ioutil.ReadFile("./db.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(file, &links)
    if err != nil {
        panic(err)
    }

    return links
}
