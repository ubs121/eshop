package main

// Crawler on ubshop.mn

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	// бааз руу холбох
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("eshop").C("p")

	offset := 1
	limit := 280

	for offset+limit < 5498 {

		resp, err := http.PostForm("https://ubshop.mn/rest/product/all",
			url.Values{"offset": {strconv.Itoa(offset)}, "limit": {strconv.Itoa(limit)}})

		defer resp.Body.Close()

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(resp.Body)

		for {
			var m map[string]map[string]interface{}

			if err := decoder.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			lst := m["data"]["list"].([]interface{})
			for i := 0; i < len(lst); i++ {
				p := lst[i].(map[string]interface{})
				p["_id"] = p["id"]

				fmt.Printf("%v\n", p["_id"])

				_, err = c.UpsertId(p["_id"], p)
				if err != nil {
					log.Fatal(err)
				}

			}

		}

		offset += limit
	}

}
