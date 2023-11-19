package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/peterstark72/rss"
)

func main() {

	if len(os.Args) < 2 {
		panic("Usage: rss <url>")
	}

	res, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	defer res.Body.Close()

	f, err := rss.NewFeed(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, itm := range f.Channel.Items {

		if itm.PubDate.T.Format(time.DateOnly) == time.Now().Format(time.DateOnly) {
			fmt.Printf("%s - %s, %s\n", itm.PubDate.T.Format("2006-01-02"), itm.Title, itm.Description)
		}
	}
}
