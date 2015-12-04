package rpc

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "github.com/ubs121/db/mongo"
	jrpc "github.com/ubs121/rpc/json"
)

const (
	tableTags = "tags"
	tableP    = "p"
)

type (
	suggestRequest struct {
		Query bson.M
		Text  string
	}

	dataRequest struct {
		ID         string
		Collection string
		Query      bson.M
		Sort       []string
		Select     []string
		Rank       bool
		Skip       int
		Limit      int
		Data       bson.M
	}

	findReply struct {
		Data  []bson.M
		Count int
	}

	// A P represents a product.
	p struct {
		ID      string    `bson:"_id"`
		Name    string    `bson:"name"`
		About   string    `bson:"about"`
		Price   float64   `bson:"price"`
		Photo   string    `bson:"photo"`
		Views   int       `bson:"views"`
		Buynum  int       `bson:"buynum"`
		Rank    float64   `bson:"rank"`
		AddTime time.Time `bson:"addtime"`
		Tags    []string  `bson:"_tags"`
	}

	// TAG represents a tag.
	TAG struct {
		ID     string  `bson:"_id"`
		Name   string  `bson:"tag"`
		Parent string  `bson:"parent"`
		Rank   float64 `bson:"rank"`
	}
)

var (
	tagProjection = []string{"_id", "tag", "parent", "rank"}
	pProjection   = []string{"_id", "name", "about", "price", "photo", "addtime", "rank", "views", "buynum", "_tags", "_events", "_attachments"}
)

// TODO: suggest хийх
func suggest(w http.ResponseWriter, r *http.Request) {
	args := suggestRequest{}
	jrpc.ParseRequest(r, &args)

	proj := []string{"_id", "name"}
	sort := []string{"rank-"}
	resp, err := db.Find(tableP, args.Query, proj, sort, 0, 5) // эхний 5

	jrpc.WriteResponse(r, w, resp, err)
}

// хандалтын (tags) трэнд, эхний 25
func tags(w http.ResponseWriter, r *http.Request) {
	proj := []string{"_id", "name", "rank"}
	sort := []string{"rank-"}
	resp, err := db.Find(tableTags, bson.M{}, proj, sort, 0, 25)

	jrpc.WriteResponse(r, w, resp, err)
}

func view(w http.ResponseWriter, r *http.Request) {
	args := dataRequest{}
	jrpc.ParseRequest(r, &args)

	obj, err := db.FindOne(tableP, bson.M{"_id": args.ID}, pProjection)

	if err == nil {
		// TODO: үзсэн тоолуур нэмэх
		// TODO: rank шинэчилэх
	}

	jrpc.WriteResponse(r, w, obj, err)
}

func like(w http.ResponseWriter, r *http.Request) {
	// TODO: зөвхөн like хийсэн бараанд notification явуулах уу?
	// TODO: Like: барааны rank шинэчилэх
}

func comment(w http.ResponseWriter, r *http.Request) {
	// TODO: comment үлдээх
}

// бараа захиалах, худалдан авах
func order(w http.ResponseWriter, r *http.Request) {
	// TODO: Order: барааны rank шинэчилэх
}

// нэг мөр хайх, энэ хэрэгтэй юу?
func findOne(w http.ResponseWriter, r *http.Request) {
	args := dataRequest{}
	jrpc.ParseRequest(r, &args)
	resp, err := db.FindOne(args.Collection, args.Query, args.Select)
	jrpc.WriteResponse(r, w, resp, err)
}

// олон мөр хайх
func find(w http.ResponseWriter, r *http.Request) {
	args := dataRequest{}
	jrpc.ParseRequest(r, &args)

	var resp findReply

	// TODO: injection шалгалт - onFind(args.Query, w, r)

	n, err := db.Count(args.Collection, args.Query)
	resp.Count = n

	// Select байхгүй бол
	if args.Select == nil {
		args.Select = pProjection
	}

	args.Sort = append(args.Sort, "rank-")

	// барааг 25 мөрөөр хязгаарлах
	if args.Collection == "p" && args.Limit > 25 {
		args.Limit = 25
	}

	// TODO: _tag хайлт бол тухайн tag-н rank-г нэмэх

	resp.Data, err = db.Find(args.Collection, args.Query, args.Select, args.Sort, args.Skip, args.Limit)

	jrpc.WriteResponse(r, w, resp, err)
}

// RegisterShopService adds this service into RPC registry
func RegisterShopService(mux *http.ServeMux) {

	mux.HandleFunc("/s/find", find)
	mux.HandleFunc("/s/findOne", findOne)
	mux.HandleFunc("/s/suggest", suggest)
	mux.HandleFunc("/tags", tags)
	mux.HandleFunc("/p/{id}", view)
	mux.HandleFunc("/p/like", like)
	mux.HandleFunc("/p/comment", comment)
	mux.HandleFunc("/p/order", order)
}
