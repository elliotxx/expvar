package main

import (
	"log"

	"github.com/elliotxx/expvar"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/debug/vars", expvar.Handler())

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
