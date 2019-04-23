package rss

//FilterFnc returns true if the items is okey
type FilterFnc func(itm Item) bool

//FilterItems filters out items that pass fn
func FilterItems(fn FilterFnc, in <-chan Item) <-chan Item {
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
