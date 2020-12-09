package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return newServer(ctx, ":8088")
	})
	g.Go(func() error {
		return hookSignal(ctx)
	})
	if err := g.Wait(); err != nil {
		fmt.Println("error group err:", err.Error())
	}
}

func newServer(ctx context.Context, addr string) error {
	s := http.Server{
		Addr: addr,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("server ctx done")
		s.Shutdown(context.Background())
	}(ctx)
	fmt.Println("start server")
	return s.ListenAndServe()
}

func hookSignal(ctx context.Context) error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
	for {
		select {
		case s := <-c:
			return fmt.Errorf("get signal %v", s)
		case <-ctx.Done():
			return fmt.Errorf("signal ctx done")
		}
	}
}
