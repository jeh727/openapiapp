package app

import (
	"fmt"
	"net/http"
	"time"

	middleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/jeh727/openapiapp/internal/app/appapi"
)

func RunServer(address string, readHeaderTimeout time.Duration) error {
	swagger, err := appapi.GetSwagger()
	if err != nil {
		return fmt.Errorf("error loading swagger spec\n%w", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStore := &PetStore{}

	mux := http.NewServeMux()

	// We now register our petStore above as the handler for the interface
	appapi.HandlerFromMux(petStore, mux)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	h := middleware.OapiRequestValidator(swagger)(mux)

	server := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: readHeaderTimeout,
		Handler:           h,
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server failed\n%w", err)
	}

	return nil
}
