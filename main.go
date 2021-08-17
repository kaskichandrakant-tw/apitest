package main

import (
	"apitest/routers"
	"log"
)

func main() {
	r := routers.Setup()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
