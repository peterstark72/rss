package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/peterstark72/rss"
)

func main() {

	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	feed := rss.ReadAll(res.Body)

	for _, itm := range feed.Channel.Items {
		fmt.Println(itm.Title)
	}

}
