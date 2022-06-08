package collector

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Collect(c *colly.Collector) {
	url := "https://vnexpress.net/rss"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
