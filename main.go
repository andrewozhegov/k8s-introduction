package main

import (
    "github.com/takama/router" // use $ go get to install it
    //"github.com/julienschmidt/httprouter" // good solution too
)

// Run server: go build && k8s-introduction
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	r := router.New()
	r.GET("/", home)
	r.Listen(":8000")
}
