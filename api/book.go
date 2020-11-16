package apis

import (
	controllers "arjun/mybooks/controller"
	"fmt"
)

type BookApi struct{}

func (b *BookApi) ViewBooks() string {
	controller := controllers.Controller{}
	fmt.Println("Hit the ViewBooks api")
	result := controller.ViewBooks()
	return result
}
