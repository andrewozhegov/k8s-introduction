package main

import (
    "net/http"
    "os"

    "github.com/Sirupsen/logrus"
    common_handlers "github.com/k8s-community/handlers/info" // use $ glide get to add
    "github.com/andrewozhegov/k8s-introduction/version"
    "github.com/k8s-community/utils/shutdown"
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

	// Readiness and liveness probes for Kubernetes
    r.GET("/info", common_handlers.Handler(version.RELEASE, version.REPO, version.COMMIT))
	r.GET("/healthz", func(c *router.Control) {
		c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
	})

    go r.Listen(":" + port) // or ("0.0.0.0:" + port)

    logger := log.WithField("event", "shutdown")
    sdHandler := shutdown.NewHandler(logger)
	sdHandler.RegisterShutdown(sd)
}

// sd does graceful shutdown of the service
func sd() (string, error) {
	// if service has to finish some tasks before shutting down, these tasks must be finished her	
    return "Ok", nil
}
