package main

import (
    "os"
    "github.com/Sirupsen/logrus"
    "github.com/takama/router" // use $ go get to install it
    //"github.com/julienschmidt/httprouter" // good solution too
)

var log = logrus.New()

// Run server: go build && k8s-introduction
// Try requests: curl http://127.0.0.1:8000/test
func main() {
    port := os.Getenv("SERVICE_PORT") // specify port when run: $ env SERVICE_PORT=8000 ./k8s-introduction
	if len(port) == 0 {
        log.Fatal("Service port in not specified")
    }
 
    r := router.New()
    r.Logger = logger
	r.GET("/", home)
    r.Listen(":" + port) // or ("0.0.0.0:" + port)
}
