package rss

//Mux is the essential structure
type Mux struct {
	Feeds <-chan Feed
}

//NewMux creates a new multiplexer
func NewMux(feeds <-chan Feed) (m *Mux) {
	m = &Mux{
		Feeds: feeds,
	}
	return
}

//Items returns channel with multiplexed items
func (m *Mux) Items() chan Item {

	items := make(chan Item)

	go func() {
		for feed := range m.Feeds {
			for _, itm := range feed.Channel.Items {
				items <- itm
			}
		}
		close(items)
	}()
	return items
}
