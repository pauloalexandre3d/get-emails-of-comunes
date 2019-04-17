package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type provinces []province

type province struct {
	name string
	link string
}

func getProvinces(r region) provinces {

	c := colly.NewCollector()

	provinces := provinces{}

	c.OnHTML("#jk > table.ut", func(e *colly.HTMLElement) {
		e.ForEach("td > a", func(_ int, e *colly.HTMLElement) {
			province := province{e.Text, e.Attr("href")}
			provinces = append(provinces, province)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.tuttitalia.it" + r.link)

	return provinces

}

func (p provinces) showMenu() {
	fmt.Println("Choose a Province:")
	for i, province := range p {
		fmt.Println(i, " - ", province.name)
	}
}
