package main

import (
	"io/ioutil"

	"fmt"
	"log"
	"os"
	"strings"

	db "github.com/ubs121/db/mongo"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: data_import <csv or json file> <table name>")
		return
	}

	log.Println("db connect...")

	db.Open("127.0.0.1", "eshop")
	defer db.Close()

	log.Println("import...")

	fileName := os.Args[1]
	colName := os.Args[2]

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// импортлох
	if strings.HasSuffix(fileName, ".json") {
		db.ImportJSON(colName, data)
	} else {
		db.ImportCSV(colName, data)
	}
}
