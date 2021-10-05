package main

import (
	"context"
	"crypto/rand"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "product-test/config"
	"product-test/services"

	h "product-test/helpers"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
)

func main() {
	//init context
	ctx, err := h.NewRepositoryContext(rand.Reader, &http.Transport{})
	if err != nil {
		log.Fatal("can't init service context :", err)
	}

	//gin setup
	gin.SetMode(gin.ReleaseMode)

	r := Routing(ctx)
	//http server
	srv := &http.Server{Addr: ":" + ctx.Config.App.Port, Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ctx.Log.Fatal("can't run service", zap.Error(err))
		}
	}()
	ctx.Log.Info(ctx.Config.App.Name + " initiated at port " + ctx.Config.App.Port)

	// gracefully shutdown
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx.Log.Info("Shutdown " + ctx.Config.App.Name + " repository")

	cts, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(cts); err != nil {
		ctx.Log.Warn("can't shutdown "+ctx.Config.App.Name+" repository", zap.Error(err))
	}

	ctx.Log.Info(ctx.Config.App.Name + " repository exiting")
}

func Routing(ctx cfg.RepositoryContext) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	pprof.Register(r)

	p := ginprometheus.NewPrometheus("gin")

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL
		url.RawQuery = ""
		return url.String()
	}
	p.Use(r)

	//services
	function := r.Group("/services")

	{
		function.POST("/add-product", services.AddProduct(ctx))
		function.GET("/list-product/:sort", services.ProductList(ctx))
		//function.POST("/get-va", bri.GetBriva(ctx))
	}

	return r
}
