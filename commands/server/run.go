package server

import (
	"net/http"
	"web/gopkg/gins"
	"web/gopkg/graceful"
	"web/handler/api"

	"github.com/urfave/cli/v2"
)

func Run(*cli.Context) error {
	go func() {
		_ = http.ListenAndServe(":8999", nil)
	}()

	server := gins.NewHttpServer(":8080")
	server.RegisterHandler(
		api.NewHandler,
	)
	graceful.Start(server)
	graceful.Wait()
	return nil
}
