package main

import (
	"encoding/xml"
	_ "encoding/xml"
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}
type Channel struct {
	Title     string `xml:"title"`
	ItemsList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	HeadLine     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nBelow are all the Google Search Trends For Today!")
	fmt.Println("======================================================")

	for i := range r.Channel.ItemsList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("search term", r.Channel.ItemsList[i].Title)
		fmt.Println("link to the Trend:", r.Channel.ItemsList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemsList[i].NewsItems[0].HeadLine)
		fmt.Println("Link to article:", r.Channel.ItemsList[i].NewsItems[0].HeadlineLink)
		fmt.Println("======================================================")
	}
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=CA")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}
