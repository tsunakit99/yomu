package main

import (
	"github.com/tsunakit99/yomu/interfaces"
)

func main() {
	router := interfaces.NewRouter()
	router.Start(":8080")
}
