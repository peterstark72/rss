package rss_test

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/peterstark72/rss"
)

var sources = []string{
	fmt.Sprintf("https://tygelsjo-20030801.appspot.com?q=%s", url.QueryEscape("Tygelsjö")),
	"http://malmo.lokaltidningen.se/rss/senestenytrss/",
	"https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/",
}

func TestMuxSources(t *testing.T) {

	items1 := rss.MuxSources(sources...)

	items2 := rss.FilterItems(items1, func(itm rss.Item) bool {
		return strings.Contains(itm.Title, "Tygelsjö") || strings.Contains(itm.Description, "Tygelsjö")
	})

	<-items2
}

func TestJSON(t *testing.T) {

	items1 := rss.MuxSources(sources...)

	items2 := rss.FilterItems(items1, func(itm rss.Item) bool {
		return strings.Contains(itm.Title, "Tygelsjö") || strings.Contains(itm.Description, "Tygelsjö")
	})

	b := new(bytes.Buffer)
	rss.WriteAsJSON(os.Stdout, items2)

	if len(b.String()) == 0 {
		t.Error("Could not write JSON")
	}
}
