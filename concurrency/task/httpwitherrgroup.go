package task

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServeManager struct {
	Stop       chan struct{}
	Serve      *http.Server
	PprofServe *http.Server
}

var (
	g, ctx  = errgroup.WithContext(context.Background())
	manager = &ServeManager{
		make(chan struct{}),
		&http.Server{Addr: ":8080"},
		&http.Server{Addr: ":8081"},
	}
)

func (m *ServeManager) shutdown() error {
	var shutdown = func(serve *http.Server) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return serve.Shutdown(ctx)
	}

	var err error
	if m.Serve != nil {
		err = shutdown(m.Serve)
	}
	if m.PprofServe != nil {
		if err1 := shutdown(m.PprofServe); err == nil {
			err = err1
		}
	}
	return err
}

func serveApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hello, serve App!")
	})
	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		manager.Stop <- struct{}{}
		_, _ = fmt.Fprint(writer, "Shutdown!")
	})
	manager.Serve.Handler = mux
	err := manager.Serve.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func servePprof() error {
	if manager.PprofServe == nil {
		return nil
	}
	err := manager.PprofServe.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func listenSignal() error {
	var c = make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		close(c)
		close(manager.Stop)
	}()
	select {
	case <-manager.Stop:
		fmt.Println("Server API shutdown...")
	case sig := <-c:
		fmt.Printf("OS signal shutdown:%v\n", sig)
	case <-ctx.Done():
		fmt.Println("App Exception")
	}

	err1 := manager.shutdown()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return err1
}

func HttpStart() error {
	g.Go(serveApp)
	g.Go(servePprof)
	g.Go(listenSignal)
	return g.Wait()
}
