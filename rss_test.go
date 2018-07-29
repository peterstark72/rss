package rss_test

import (
	"net/http"
	"testing"

	"github.com/peterstark72/rss"
)

func TestPolisen(t *testing.T) {

	const url = "https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/"

	res, err := http.Get(url)
	if err != nil {
		t.Error("Could not load")
		return
	}
	defer res.Body.Close()

	feed := rss.ReadAll(res.Body)

	if len(feed.Channel.Items) == 0 {
		t.Error("Missing items")
	}
}

func TestMyNewsDesk(t *testing.T) {

	const url = "http://www.mynewsdesk.com/se/search/rss?page=1&query=tygelsj%C3%B6&sites%5B%5D=se&type_of_medias=&utf8=%E2%9C%93"

	res, err := http.Get(url)
	if err != nil {
		t.Error("Could not load")
		return
	}
	defer res.Body.Close()

	feed := rss.ReadAll(res.Body)

	if len(feed.Channel.Items) == 0 {
		t.Error("Missing items")
	}
}
