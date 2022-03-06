package bots

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type News struct {
	Headline string
	Link     string
}

//THESE TYPES WILL LIKELY NEED TO CHANGE
type Stock struct {
	Name     string //ticker
	Price    string
	Change   string //% change on the day
	TopStory News
}

func StockScrap(conn *gin.Context) {
	URL := conn.Request.URL.Query().Get("url")
	if URL == "" {
		log.Fatal("request does not have a url")
	}
	log.Println("visiting", URL)

	c := colly.NewCollector()

	var stockUpdate []Stock

	c.OnHTML(".e1AOyf", func(e *colly.HTMLElement) {
		var stock Stock
		var stockNews News

		stock.Name = e.ChildText(".zzDege")
		stock.Price = e.ChildText(".YMlKec")
		stock.Change = e.ChildText(".JwB6zf")[0:5]

		stockNews.Headline = e.ChildText(".Yfwt5")
		stockNews.Link = e.ChildAttr("a[href]", "href")
		stock.TopStory = stockNews

		if stock.Name != "" {
			stockUpdate = append(stockUpdate, stock)
		}
	})

	c.Visit(URL)

	info, err := json.Marshal(stockUpdate)
	if err != nil {
		log.Fatal(err)
	}

	conn.JSON(200, string(info))
}
