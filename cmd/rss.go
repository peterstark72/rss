package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/peterstark72/rss"
)

func main() {

	if len(os.Args) == 0 {
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
		fmt.Printf("%s - %s\n", itm.PubDate.T.Format("2006-01-02"), itm.Title)
	}
}
