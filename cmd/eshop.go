package main

import (
	"eshop/rpc"
	"log"
	"net/http"

	db "github.com/ubs121/db/mongo"
)

func main() {
	log.Println("db connect...")

	db.Open("127.0.0.1", "eshop")
	defer db.Close()

	mux := http.NewServeMux()
	rpc.RegisterShopService(mux)
	//rpc.RegisterAdminService(mux)

	log.Println("eshop is started...")
	http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", mux)

	log.Println("eshop is stopped.")
}
