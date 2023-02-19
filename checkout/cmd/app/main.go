package main

import (
	"fmt"
	"net/http"
	"route256/libs/jsonreqwrap"
)

// const port = ":8080"

const HACKER_NEWS_URL = "https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty"

type HackerNewsReq struct{}
type HackerNewsRes struct {
	By          string  `json:"by"`
	Descendants int64   `json:"descendants"`
	ID          int64   `json:"id"`
	Kids        []int64 `json:"kids"`
	Score       int64   `json:"score"`
	Time        int64   `json:"time"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	URL         string  `json:"url"`
}

func main() {
	client := jsonreqwrap.NewClient[HackerNewsReq, HackerNewsRes](HACKER_NEWS_URL, http.MethodGet)

	res, err := client.Run(HackerNewsReq{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", res)
}
