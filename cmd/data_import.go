package main


import (
	"io/ioutil"

	"eshop/db"
	"os"
	"log"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: data_import <csv or json file> <table name>")
		return
	}

	log.Println("db connect...")

	db.Open("127.0.0.1")
	defer db.Close()

	log.Println("import...")

	fileName:=os.Args[1]
	colName:=os.Args[2]

	data, err := ioutil.ReadFile(fileName)

	if err!= nil {
		panic(err)
	}

  // импортлох
	if strings.HasSuffix(fileName, ".json") {
		db.ImportJson(colName, data)
	} else {
		db.ImportCsv(colName, data)
	}
}
