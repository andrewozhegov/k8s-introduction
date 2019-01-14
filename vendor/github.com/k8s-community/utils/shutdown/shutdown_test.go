package shutdown

import (
	"fmt"
	"os"
	"syscall"
	"testing"

	"bufio"
	"bytes"
	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"strings"
	"sync"
	"time"
)

// TestHandlerSuite test engine
func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

// HandlerSuite is a suite to test enigne
type HandlerSuite struct {
	suite.Suite
}

var wg sync.WaitGroup

// TestShutdownNoError check case for shutdown without error
/*func (es *HandlerSuite) TestShutdownNoError() {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = writer

	logger := log.WithField("test", "TestShutdownNoError")

	h := NewHandler(logger)
	h.AddShutdownSignal(syscall.SIGUSR1)

	wg.Add(1)

	h.RegisterShutdown(shutdownNoError)

	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	es.NoError(err)

	err = process.Signal(syscall.SIGUSR1)
	es.NoError(err)

	wg.Wait()

	time.Sleep(2 * time.Second)

	data := strings.Split(b.String(), "\n")
	fmt.Println(data)
	es.Equal(3, len(data))
}

// TestShutdownWithError check case for shutdown with error
func (es *HandlerSuite) TestShutdownWithError() {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = writer

	logger := log.WithField("test", "TestShutdownWithError")

	h := NewHandler(logger)
	h.AddShutdownSignal(syscall.SIGUSR1)

	wg.Add(1)
	h.RegisterShutdown(shutdownWithError)

	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	es.NoError(err)

	err = process.Signal(syscall.SIGUSR1)
	es.NoError(err)

	wg.Wait()

	err = writer.Flush()
	es.NoError(err)

	//	data := strings.Split(b.String(), "\n")
	//	es.Equal(3, len(data))
}*/

func shutdownNoError() (string, error) {
	defer wg.Done()
	return "OK", nil
}

func shutdownWithError() (string, error) {
	defer wg.Done()
	msg := "error"
	return msg, fmt.Errorf("Some error: %s", msg)
}
