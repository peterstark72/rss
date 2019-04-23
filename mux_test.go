package rss_test

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/peterstark72/rss"
)

var q = "Tygelsjö"

var urls = []string{
	"https://tygelsjo-20030801.appspot.com?q=Tygelsjö",
	"http://malmo.lokaltidningen.se/rss/senestenytrss/",
	"https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/",
}

func titles(itm rss.Item) bool {
	return strings.Contains(itm.Title, q) || strings.Contains(itm.Description, q)
}

func TestMux(t *testing.T) {

	feeds := make(chan rss.Feed)

	m := rss.NewMux(feeds)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {

			defer wg.Done()

			res, err := http.Get(u)
			if err != nil {
				return
			}
			defer res.Body.Close()

			f, err := rss.NewFeed(res.Body)
			if err != nil {
				return
			}
			feeds <- *f
		}(u)
	}
	var items []rss.Item
	go func() {
		for itm := range rss.FilterItems(titles, m.Items()) {
			items = append(items, itm)
		}
	}()
	wg.Wait()

}
