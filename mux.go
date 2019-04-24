package rss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

//MuxSources muxes
func MuxSources(sources ...string) chan Item {

	items := make(chan Item)

	go func() {

		var wg sync.WaitGroup

		for _, lnk := range sources {

			wg.Add(1)
			go func(u string) {

				defer wg.Done()

				res, err := http.Get(u)
				if err != nil {
					return
				}
				defer res.Body.Close()

				feed, err := NewFeed(res.Body)
				if err != nil {
					return
				}
				for _, itm := range feed.Channel.Items {
					items <- itm
				}

			}(lnk)
		}
		wg.Wait()
		close(items)
	}()

	return items
}

//FilterItems filters items through fn
func FilterItems(items <-chan Item, fn func(Item) bool) chan Item {
	out := make(chan Item)
	go func() {
		for itm := range items {
			if fn(itm) {
				out <- itm
			}
		}
		close(out)
	}()
	return out
}

//WriteAsJSON writes Items as JSON array to w
func WriteAsJSON(w io.Writer, items <-chan Item) (err error) {

	type JSItem struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"url"`
		PubDate     string `json:"pubdate"`
		Source      string `json:"source"`
		Image       string `json:"image"`
	}

	var result []JSItem
	for itm := range items {

		d, _ := itm.ParsePubDate()
		result = append(result, JSItem{
			Title:       itm.Title,
			Description: itm.Description,
			Link:        itm.Link,
			PubDate:     d.Format("2006-01-02"), //JS friendly date format
			Source:      itm.Source.Name,
			Image:       itm.Enclosure.URL,
		})
	}

	data, err := json.Marshal(result)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%s", data)

	return
}
