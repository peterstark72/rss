package rss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
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
func FilterItems(in <-chan Item, fn func(Item) bool) chan Item {
	out := make(chan Item)
	go func() {
		for itm := range in {
			if fn(itm) {
				out <- itm
			}
		}
		close(out)
	}()
	return out
}

//WriteAsJSON writes Items as JSON array to w
func WriteAsJSON(w io.Writer, in <-chan Item) (err error) {

	var items []Item
	for itm := range in {
		items = append(items, itm)
	}

	sort.Sort(DateSorter(items))

	data, err := json.Marshal(items)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%s", data)

	return
}
