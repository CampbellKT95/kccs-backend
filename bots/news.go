package bots

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Response struct {
	Image    string
	Headline string
	Link     string
}

func NewsScrap(conn *gin.Context) {
	// verify the param url exists from the request. grabs param variable 'url'
	URL := conn.Request.URL.Query().Get("url")
	if URL == "" {
		log.Fatal("request does not have a url")
	}
	log.Println("visiting", URL)

	//create a new collector
	//can designate allowed domains with colly.alloweddomains as a param
	c := colly.NewCollector()

	//stores the data we want
	var topStories []Response

	//onHTML allows collector to fire callback function when the specific html tag is reached (in this case, an <a> with an href)
	//the callback fires an anon function which gets the info from the href & appends it to response
	c.OnHTML(".module--promo .module__content .media-list .media-list__item", func(e *colly.HTMLElement) {
		var story Response

		story.Image = e.ChildAttr("img[src]", "src")
		story.Headline = e.ChildText("a[href]")
		story.Link = e.ChildAttr("a[href]", "href")

		if story.Headline != "" {
			topStories = append(topStories, story)
		}
	})

	//function to visit the site & begin scrapping
	c.Visit(URL)

	//grabs just the top 5 results & turns them into json to be returned
	// info, err := json.Marshal(topStories[1:6])
	info, err := json.MarshalIndent(topStories, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	//return response to user
	conn.JSON(200, string(info))
}
