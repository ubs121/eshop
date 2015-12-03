package tags

import "eshop/db"

type ByRank []db.TAG

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[i].Rank > a[j].Rank }


// product ranking
func ProductRank() {
	// TODO: Барааг үзсэн тоо, хайлтын тоо (түлхүүр үгс), like тоо, захиалсан тоо, худалдаж авсан тоогоор оноолох,
	// TODO: үнийг мөн тусгаж болно. 100'000 тутамд 1 оноо
}
