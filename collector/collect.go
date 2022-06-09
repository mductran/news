package collector

import (
	"fmt"
	"newsonthego/sites"

	"github.com/gocolly/colly"
)

func Collect(c *colly.Collector) {
	url := "https://vnexpress.net/rss"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}

func ScrapeArticle(url string) {
	c := colly.NewCollector(colly.AllowedDomains(sites.GetAllowedDomains()...))

	c.OnHTML("div[class='fck-detail']", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
