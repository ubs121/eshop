package main

import (
	"eshop/tags"
	"fmt"
	"log"
	"sort"

	db "github.com/ubs121/db/mongo"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	log.Println("db connect...")

	db.Open("127.0.0.1", "eshop")
	defer db.Close()

	q := db.C(db.TABLE_TAGS).Find(bson.M{})

	var ts []db.TAG
	err := q.All(&ts)
	if err != nil {
		panic(err)
	}

	// TODO: барааны rank-г tag rank-д оруулж тооцох

	graph := tags.New()
	for _, t := range ts {
		//fmt.Printf("%v\n", t["parent"])
		if t.Name != "" && t.Parent != "" {
			graph.Link(t.Name, t.Parent)
		}
	}

	probability_of_following_a_link := 0.85 // The bigger the number, less probability we have to teleport to some random link
	tolerance := 0.0001                     // the smaller the number, the more exact the result will be but more CPU cycles will be needed

	sortedTags := []db.TAG{}
	graph.Rank(probability_of_following_a_link, tolerance, func(t string, rank float64) {
		sortedTags = append(sortedTags, db.TAG{Name: t, Rank: rank})
	})

	sort.Sort(tags.ByRank(sortedTags))

	for _, t := range sortedTags {
		fmt.Println(t)
	}
}
