package main

import (
	apis "arjun/mybooks/api"
	driver "arjun/mybooks/drivers"
	"database/sql"
	"fmt"
	"log"

	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {

	db = driver.ConnectDB()
	fmt.Println("hello books...")
	books := apis.BookApi{}
	result := books.ViewBooks()
	fmt.Println(result)
}
