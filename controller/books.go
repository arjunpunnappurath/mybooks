package controllers

import (
	"arjun/mybooks/repo"
	"fmt"
	"log"
)

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (c *Controller) ViewBooks() string {
	repo := repo.Repo{}
	result := repo.ViewBooks()
	fmt.Println("Hit view books in controllers")
	return result
}
