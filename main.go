package main

import (
	"fmt"
	"newsonthego/collector"
)

func main() {
	fmt.Println("crawling")

	collector.ScrapeArticle("https://vnexpress.net/khau-trang-deo-tiep-hay-thoi-4473790.html")
}
