package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type region struct {
	name string
	link string
}

type regions []region

func getRegions() regions {
	c := colly.NewCollector()

	regions := regions{}

	c.OnHTML("#ar > div > div.fa > dl", func(e *colly.HTMLElement) {
		e.ForEach("dt > a", func(_ int, e *colly.HTMLElement) {
			region := region{e.Text, e.Attr("href")}
			regions = append(regions, region)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.tuttitalia.it/")

	return regions
}

func (r regions) showMenu() {
	fmt.Println("Choose a Region:")
	for i, region := range r {
		fmt.Println(i, " - ", region.name)
	}
}
