package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/peterstark72/rss"
)

func main() {

	if len(os.Args) != 2 {
		panic("Please enter <RSS Feed URL>")
	}

	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		return
	}

	feed := rss.ReadAll(res.Body)

	for _, itm := range feed.Channel.Items {
		fmt.Println(itm.Title)
	}

}
