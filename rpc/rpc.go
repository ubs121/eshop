package rpc

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	HttpMethods string = "POST, GET, OPTIONS"
)

type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

/// хүсэлт унших
func ParseRequest(r *http.Request, v interface{}) error {
	// TODO: хүсэлт зөвхөн зөвшөөрөгдсөн сайтаас үүссэн болохыг шалгах

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)

	// DEBUG ONLY
	s, _ := json.Marshal(v)
	log.Printf("rpc << %s: %s\n", r.URL.String(), string(s))

	return err
}

/// хариу бичих
func WriteResponse(r *http.Request, w http.ResponseWriter, result interface{}, e error) {
	// copy Access-Control-Request-Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", HttpMethods)
	w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	doc := map[string]interface{}{}
	if e != nil {
		doc["Error"] = e.Error()
	} else {
		doc["Result"] = result
	}

	// debug
	s, _ := json.Marshal(doc)
	log.Printf("rpc >> %s\n", string(s))

	encoder := json.NewEncoder(w)
	encoder.Encode(doc)

}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func MakeGzipHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// gzip
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()

			handler.ServeHTTP(GzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
		} else {
			handler.ServeHTTP(w, r)
		}

	})
}
