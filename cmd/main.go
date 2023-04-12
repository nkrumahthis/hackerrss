package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Title		string `xml:"title"`
	Description	string `xml:"description"`
	Link		string `xml:"link"`
}

type Channel struct {
	Title 	string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	Items []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}

func main() {
	fmt.Println("Hacker News")

	// parse the mini flag
	var mini bool
	flag.BoolVar(&mini, "mini", false, "show only title and link")
	flag.Parse()

	// Fetch RSS feed
	resp, err := http.Get("https://feeds.feedburner.com/TheHackersNews")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Parse RSS feed
	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		log.Fatal(err)
	}

	// Display news items in console
	fmt.Println("Latest news:")
	for _, item := range rss.Channel.Items {
		fmt.Println("Title:", item.Title)
		if !mini {
			fmt.Println("Description:", item.Description)
		}
		fmt.Println("Link:", item.Link)
		fmt.Println()
	}

}