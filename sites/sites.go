package sites

import (
	// "encoding/json"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Collection struct {
	Collection []Language `json:"collection"`
}

type Language struct {
	Lang  string `json:"language"`
	Sites []Site `json:"sites"`
}

type Site struct {
	Name       string     `json:"name"`
	Url        string     `json:"url"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func ReadSites() *Collection {
	file, err := ioutil.ReadFile("sites/sites.json")
	if err != nil {
		panic(err)
	}

	var result Collection
	err = json.Unmarshal(file, &result)
	if err != nil {
		panic(err)
	}
	return &result
}

func GetAllowedDomains() []string {
	var domains []string
	collection := ReadSites()
	for i := 0; i < len(collection.Collection); i++ {
		sites := collection.Collection[i].Sites
		for j := 0; j < len(sites); j++ {
			url := sites[j].Url
			domains = append(domains, url)
			domains = append(domains, strings.TrimPrefix(url, "https://"))
		}
	}
	return domains
}
