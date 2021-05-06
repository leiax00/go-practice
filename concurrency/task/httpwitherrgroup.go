package task

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var g, _ = errgroup.WithContext(context.Background())
var c = make(chan os.Signal)

func serveApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hello, serve App!")
	})
	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		c <- syscall.SIGQUIT
		_, _ = fmt.Fprint(writer, "Shutdown!")
	})
	return http.ListenAndServe(":8080", mux)
}

func serveDebug() error {
	return http.ListenAndServe(":8081", http.DefaultServeMux)
}

func listenSignal() error {
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		close(c)
	}()
	sig := <-c
	fmt.Println("Do something for quit...")

	return errors.Errorf("SYSTEM INTERRUPT: %v", sig)
}

func HttpStart() error {
	g.Go(serveApp)
	g.Go(serveDebug)
	g.Go(listenSignal)
	return g.Wait()
}
