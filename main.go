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
	"github.com/go-gateway/api"
	"github.com/go-gateway/configs"
	"github.com/go-gateway/internal/data/db"
	"github.com/go-gateway/logs"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

var cfg string

func init() {
	flag.StringVar(&cfg, "c", "", "cofnig file")
	flag.Parse()
}

func main() {
	err := configs.InitializeConfig(cfg)
	if err != nil {
		log.Println(err)
		panic("load config failed")
	}

	err = logs.InitLog(viper.GetString("log.path"))
	if err != nil {
		panic("init logging failed")
	}

	err = db.Init()
	if err != nil {
		log.Println(err)
		panic("init db failed")
	}

	gin.SetMode(viper.GetString("runmode"))
	router := gin.New()
	var handlers []gin.HandlerFunc
	api.InitializeRouter(router, handlers...)

	srv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: router,
	}

	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("stooping http server")
		return srv.Shutdown(errCtx)
	})

	g.Go(func() error {
		log.Infof("starting http server at address: %s", viper.GetString("addr"))
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-done:
			log.Println("sig:", sig)
			// clean resources
			db.Close()
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
