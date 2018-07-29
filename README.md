# Simple RSS Reader

A very simple RSS reader.

```
    const url = "https://polisen.se/aktuellt/rss/skane/handelser-rss---skane/"

    res, err := http.Get(url)
    if err != nil {
        return
    }
    defer res.Body.Close()

    feed := rss.ReadAll(res.Body)

    fmt.Println(feed.Channel.Title)

    for _, itm := range feed.Channel.Items {
        fmt.Println(itm)
    }
```