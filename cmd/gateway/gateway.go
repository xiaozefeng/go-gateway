package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api"
	"github.com/xiaozefeng/go-gateway/internal/pkg/configs"
	"github.com/xiaozefeng/go-gateway/internal/pkg/logs"
	"github.com/xiaozefeng/go-gateway/internal/pkg/middleware"
	"github.com/xiaozefeng/go-gateway/internal/pkg/wire"
	"golang.org/x/sync/errgroup"
)

var cfg string

func init() {
	flag.StringVar(&cfg, "c", "", "cofnig file")
	flag.Parse()
}

func main() {
	err := InitDepend()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	engine := gin.New()
	var handlers []gin.HandlerFunc
	handlers = append(handlers, gin.Recovery())
	handlers = append(handlers, middleware.NoCache)
	handlers = append(handlers, middleware.Options)
	handlers = append(handlers, middleware.Secure)
	handlers = append(handlers, middleware.Login)
	handlers = append(handlers, gin.Logger())

	api.InitRouter(engine, handlers...)

	server := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: engine,
	}

	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("stooping http server")
		return server.Shutdown(errCtx)
	})

	g.Go(func() error {
		log.Infof("starting http server at address: %s", viper.GetString("addr"))
		return server.ListenAndServe()
	})

	g.Go(func() error {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-done:
			log.Println("sig:", sig)
			// clean resources
			wire.GetDB().Close()
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

func InitDepend() error {
	err := configs.Init(cfg)
	if err != nil {
		return err
	}

	err = logs.Init(viper.GetString("log.path"))
	if err != nil {
		return err
	}
	err = wire.InitDI()
	if err != nil {
		return err
	}
	return nil
}
