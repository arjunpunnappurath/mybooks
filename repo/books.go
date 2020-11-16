package repo

import (
	"fmt"
	"log"
)

type Repo struct{}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (r *Repo) ViewBooks() string {
	fmt.Println("Hit the view books function...")
	return "Hit view books in repo"
}
