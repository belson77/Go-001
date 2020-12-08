package main

import (
	"fmt"
	"log"
	"reflect"
	"net/http"
	"github.com/belson7/Go-000/Week02/dao"
	"github.com/belson7/Go-000/Week02/service"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovr ", r)
		}
	}()

	errCh := make(chan error)

	go func() {
		for {
			err := <-errCh
			fmt.Printf("%+v, %v\n", err, reflect.TypeOf(err))
		}
	}()

	d := dao.New()
	s := service.New(d, errCh)

	http.HandleFunc("/user/get", s.GetUserHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}