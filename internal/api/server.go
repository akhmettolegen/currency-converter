package api

import (
	"context"
	"fmt"
	"github.com/akhmettolegen/currency-converter/pkg/application"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error){

	app, err := application.Get()
	if err != nil {
		return nil, err
	}

	router, err := New(*app)

	srv := http.Server{
		Addr: fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port),
		Handler: router,
	}
	return &Server{&srv}, nil
}

func (srv *Server) Start() {
	fmt.Println("Starting server..")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Println("Listening on", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Println("Shutting down server.. Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Server gracefully stopped")
}