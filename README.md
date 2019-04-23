# Simple RSS package

A very simple RSS package for reading RSS Feeds.

```Go
    const url = "https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/"

    res, err := http.Get(url)
    if err != nil {
        return
    }
    defer res.Body.Close()

    feed, _ := rss.NewFeed(res.Body)

    fmt.Println(feed.Channel.Title)

    for _, itm := range feed.Channel.Items {
        fmt.Println(itm)
    }
```