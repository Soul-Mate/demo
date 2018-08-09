package main

import (
	"github.com/Soul-Mate/demo/soa_demo/controllers"
	"net/http"
	"log"
)

func main()  {
	orderC := new(controllers.OrderController)
	http.HandleFunc("/orders/create", orderC.Create())
	log.Fatal(http.ListenAndServe(":8080", nil))
}