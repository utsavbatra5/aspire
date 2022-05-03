package api

import (
	"aspire/config"
	"aspire/logger"
	"aspire/middleware"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitAPI() error {
	var (
		err error
	)

	httpport := fmt.Sprintf(":%d", config.Conf.Server.Port)

	println("Server is listening to ", httpport)

	router := gin.Default()
	srv := &http.Server{
		Addr:    httpport,
		Handler: router,
	}

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	go func() {
		select {
		case <-interrupt:
			time.Sleep(1 * time.Millisecond)
			srv.Shutdown(context.TODO())
			break
		}
	}()

	router.Use(gin.Recovery())

	// all requests are routed through authentication middleware
	loanRouter := router.Group("/loan", middleware.AuthenticationMiddleware())
	loanRouter.GET("/getLoan", GetUserLoans)
	loanRouter.POST("/create", CreateUserLoan)
	loanRouter.PUT("/approve", ApproveLoan)
	loanRouter.PUT("/payEmi", PayEMI)

	if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.ZeroLog().Error().Err(err).Msg("Failed to start server" + error.Error(err))
	}
	return nil
}
