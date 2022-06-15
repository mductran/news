package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// fmt.Println("crawling")

	// collector.ScrapeArticle("https://vnexpress.net/khau-trang-deo-tiep-hay-thoi-4473790.html")

	c := colly.NewCollector(colly.AllowedDomains("vnexpress.net"))
	url := "https://vnexpress.net/rss/tin-moi-nhat.rss"

	c.OnXML("//channel/item", func(e *colly.XMLElement) {
		fmt.Println(e.ChildAttr("//title", ""))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
