package main

import (
	"context"
	"flag"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server/middleware"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/pkg/configs"
	"github.com/xiaozefeng/go-gateway/internal/pkg/logs"
	"golang.org/x/sync/errgroup"
)

func main() {
	var cfg string
	flag.StringVar(&cfg, "c", "", "cofing file")
	flag.Parse()

	err := configs.Initiliaze(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = logs.Initiliaze(viper.GetString("log.path"))
	if err != nil {
		log.Fatal(err)
	}
	cleanup, err := initDI()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewHTTPServer(viper.GetString("addr"), middleware.Login,
		middleware.NoCache,
		middleware.Options,
		middleware.Secure)

	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("stooping http s")
		return s.Shutdown(errCtx)
	})

	g.Go(func() error {
		log.Infof("starting http s at address: %s", viper.GetString("addr"))
		return s.ListenAndServe()
	})

	g.Go(func() error {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-done:
			log.Println("sig:", sig)
			// clean resources
			cleanup()
			cancel()
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Infof("group err: %v", err)
	}
}

func initDI() (func(), error) {
	routerSvc, cleanup, err := InitRouterService(eureka.ServerURL(viper.GetString("eureka_url")), "")
	if err != nil {
		return cleanup, err
	}
	server.SetRouterService(routerSvc)
	middleware.SetRouterService(routerSvc)
	return cleanup, nil
}
