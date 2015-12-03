package rpc

import (
	"eshop/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type (
	SuggestRequest struct {
		Query bson.M
		Text  string
	}

	DataRequest struct {
		Id         string
		Collection string
		Query      bson.M
		Sort       []string
		Select     []string
		Rank       bool
		Skip       int
		Limit      int
		Data       bson.M
	}

	FindReply struct {
		Data  []bson.M
		Count int
	}
)

var (
	TAG_PROJ = []string{"_id", "tag", "parent", "rank"}
	P_PROJ   = []string{"_id", "name", "about", "price", "photo", "addtime", "rank", "views", "buynum", "_tags", "_events", "_attachments"}
)

// TODO: suggest хийх
func Suggest(w http.ResponseWriter, r *http.Request) {
	args := SuggestRequest{}
	ParseRequest(r, &args)

	proj := []string{"_id", "name"}
	sort := []string{"rank-"}
	resp, err := db.Find(db.TABLE_P, args.Query, proj, sort, 0, 5) // эхний 5

	WriteResponse(r, w, resp, err)
}

// хандалтын (tags) трэнд, эхний 25
func Tags(w http.ResponseWriter, r *http.Request) {
	proj := []string{"_id", "name", "rank"}
	sort := []string{"rank-"}
	resp, err := db.Find(db.TABLE_TAGS, bson.M{}, proj, sort, 0, 25)

	WriteResponse(r, w, resp, err)
}

func View(w http.ResponseWriter, r *http.Request) {
	args := DataRequest{}
	ParseRequest(r, &args)

	obj, err := db.FindOne(db.TABLE_P, bson.M{"_id": args.Id}, P_PROJ)

	if err == nil {
		// TODO: үзсэн тоолуур нэмэх
		// TODO: rank шинэчилэх
	}

	WriteResponse(r, w, obj, err)
}

func Like(w http.ResponseWriter, r *http.Request) {
	// TODO: зөвхөн like хийсэн бараанд notification явуулах уу?
	// TODO: Like: барааны rank шинэчилэх
}

func Comment(w http.ResponseWriter, r *http.Request) {
	// TODO: comment үлдээх
}

func Related(w http.ResponseWriter, r *http.Request) {

}

// бараа захиалах, худалдан авах
func Order(w http.ResponseWriter, r *http.Request) {
	// TODO: Order: барааны rank шинэчилэх
}

// нэг мөр хайх, энэ хэрэгтэй юу?
func FindOne(w http.ResponseWriter, r *http.Request) {
	args := DataRequest{}
	ParseRequest(r, &args)
	resp, err := db.FindOne(args.Collection, args.Query, args.Select)
	WriteResponse(r, w, resp, err)
}

// олон мөр хайх
func Find(w http.ResponseWriter, r *http.Request) {
	args := DataRequest{}
	ParseRequest(r, &args)

	var resp FindReply

	// TODO: injection шалгалт - onFind(args.Query, w, r)

	n, err := db.Count(args.Collection, args.Query)
	resp.Count = n

	// Select байхгүй бол
	if args.Select == nil {
		args.Select = P_PROJ
	}

	args.Sort = append(args.Sort, "rank-")

	// барааг 25 мөрөөр хязгаарлах
	if args.Collection == "p" && args.Limit > 25 {
		args.Limit = 25
	}

	// TODO: _tag хайлт бол тухайн tag-н rank-г нэмэх

	resp.Data, err = db.Find(args.Collection, args.Query, args.Select, args.Sort, args.Skip, args.Limit)

	WriteResponse(r, w, resp, err)
}

func RegisterShopService(mux *http.ServeMux) {

	mux.HandleFunc("/s/find", Find)
	mux.HandleFunc("/s/findOne", FindOne)
	mux.HandleFunc("/s/suggest", Suggest)
	mux.HandleFunc("/tags", Tags)
	mux.HandleFunc("/p/{id}", View)
	mux.HandleFunc("/p/like", Like)
	mux.HandleFunc("/p/comment", Comment)
	mux.HandleFunc("/p/order", Comment)
}
