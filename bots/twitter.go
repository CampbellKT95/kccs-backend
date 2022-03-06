package bots

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RetrieveTweets() {

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.twitter.com/2/users/by/username/ryanmdahl", nil)
	if err != nil {
		log.Fatal(err)
	}

	bearerToken := "Bearer" + " " + os.Getenv("BEARER_TOKEN")
	request.Header.Add("Authorization", bearerToken)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
