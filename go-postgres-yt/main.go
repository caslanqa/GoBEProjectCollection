package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caslanqa/go-postgres-yt/router"
)

func main()  {
	r := router.Router()
	fmt.Println("Server starting on port:8080")

	log.Fatal(http.ListenAndServe(":8080",r))
}