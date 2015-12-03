package main

import (
	"eshop/db"
	"eshop/rpc"
	//"fmt"
	"net/http"
	"log"
)

func main() {
	log.Println("db connect...")

	db.Open("127.0.0.1")
	defer db.Close()

	mux := http.NewServeMux()
	rpc.RegisterShopService(mux)
	//rpc.RegisterAdminService(mux)

	log.Println("eshop is started...")
	http.ListenAndServe(":3000", mux)
}
