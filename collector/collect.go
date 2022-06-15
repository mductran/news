package collector

import (
	"context"
	"fmt"
	"log"
	"newsonthego/sites"
	"time"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Collect(c *colly.Collector) {
	url := "https://vnexpress.net/rss"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}

func ExtractArticle(e *colly.HTMLElement) bson.D {
	return bson.D{
		{Key: "title", Value: e.ChildAttr("title", "")},
		{Key: "description", Value: e.ChildAttr("description", "")},
		{Key: "pubDate", Value: e.ChildAttr("pubDate", "")},
		{Key: "link", Value: e.ChildAttr("link", "")},
	}
}

func ScrapeArticle(url string) {
	c := colly.NewCollector(colly.AllowedDomains(sites.GetAllowedDomains()...))

	c.OnHTML("article.fck_detail", func(e *colly.HTMLElement) {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}

		context, cancel := context.WithTimeout(context.Background(), 10*time.Second) // set timeout deadline to 10s or operation deadline, whichever is shorter
		err = client.Connect(context)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(context)
		defer cancel()

		database := client.Database("news")
		articleCollection := database.Collection("articles")

		articleWrite, err := articleCollection.InsertOne(context, bson.D{
			{Key: url, Value: e.Text},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted article id ", articleWrite.InsertedID)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
