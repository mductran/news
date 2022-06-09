package main

import (
	"fmt"
	"newsonthego/mongo"
)

func main() {

	mongo.Init()
	fmt.Println("crawling")
}
