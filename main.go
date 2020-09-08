package main

import (
	"aggregator_info/handler"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", "9011")
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/handler_1inch", handler.OneInch_handler)
	e.POST("/handler_bancor", handler.HandlerBancor)
	e.POST("/handler_paraswap", handler.HandlerParaswap)
	e.POST("/handler_kyberswap", handler.HandlerKyberswap)
	e.POST("/handler_zeroX", handler.ZeroX_handler)
	e.POST("/handler_mooniswap", handler.Mooniswap_handler)
	e.POST("/handler_dforce", handler.Dforce_handler)
	e.POST("/handler_uniswap_v2", handler.Uniswap_v2_handler)
	e.POST("/handler_sushiswap", handler.Sushiswap_handler) // TODO: ERROR
	e.POST("/handler_curve", handler.Curve_handler)

	// dYdX
	// uniswap v1
	// 0x
	// balancer
	// DDEX
	// Loopring
	// DoDo
	// Oasis
	// IDEX
	// DEX.AG
	// Tokenlon

	go func() {
		if err := e.Start(viper.GetString("port")); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// go getData()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
