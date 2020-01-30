// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ASinha24/supermartswaggerAPI/restapi/operations"
	"github.com/ASinha24/supermartswaggerAPI/restapi/operations/item"
)

//go:generate swagger generate server --target ..\..\supermartswaggerAPI --name SupermartExample --spec ..\swagger.yaml

func configureFlags(api *operations.SupermartExampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SupermartExampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ItemCreateHandler == nil {
		api.ItemCreateHandler = item.CreateHandlerFunc(func(params item.CreateParams) middleware.Responder {
			return middleware.NotImplemented("operation item.Create has not yet been implemented")
		})
	}
	if api.ItemGetHandler == nil {
		api.ItemGetHandler = item.GetHandlerFunc(func(params item.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation item.Get has not yet been implemented")
		})
	}
	if api.ItemListHandler == nil {
		api.ItemListHandler = item.ListHandlerFunc(func(params item.ListParams) middleware.Responder {
			return middleware.NotImplemented("operation item.List has not yet been implemented")
		})
	}
	if api.ItemDeleteHandler == nil {
		api.ItemDeleteHandler = item.DeleteHandlerFunc(func(params item.DeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation item.Delete has not yet been implemented")
		})
	}
	if api.ItemPutHandler == nil {
		api.ItemPutHandler = item.PutHandlerFunc(func(params item.PutParams) middleware.Responder {
			return middleware.NotImplemented("operation item.Put has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
