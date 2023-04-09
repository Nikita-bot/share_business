// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"wash_bonus/internal/app"
	"wash_bonus/openapi/restapi/operations"
	"wash_bonus/openapi/restapi/operations/session"
	"wash_bonus/openapi/restapi/operations/standard"
	"wash_bonus/openapi/restapi/operations/user"
)

//go:generate swagger generate server --target ../../openapi --name WashBonus --spec ../swagger.yaml --principal wash_bonus/internal/app.Auth --exclude-main --strict-responders

func configureFlags(api *operations.WashBonusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WashBonusAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.AuthKeyAuth == nil {
		api.AuthKeyAuth = func(token string) (*app.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.SessionAssignUserToSessionHandler == nil {
		api.SessionAssignUserToSessionHandler = session.AssignUserToSessionHandlerFunc(func(params session.AssignUserToSessionParams, principal *app.Auth) session.AssignUserToSessionResponder {
			return session.AssignUserToSessionNotImplemented()
		})
	}
	if api.UserGetBalanceHandler == nil {
		api.UserGetBalanceHandler = user.GetBalanceHandlerFunc(func(params user.GetBalanceParams, principal *app.Auth) user.GetBalanceResponder {
			return user.GetBalanceNotImplemented()
		})
	}
	if api.UserGetProfileHandler == nil {
		api.UserGetProfileHandler = user.GetProfileHandlerFunc(func(params user.GetProfileParams, principal *app.Auth) user.GetProfileResponder {
			return user.GetProfileNotImplemented()
		})
	}
	if api.SessionGetSessionHandler == nil {
		api.SessionGetSessionHandler = session.GetSessionHandlerFunc(func(params session.GetSessionParams, principal *app.Auth) session.GetSessionResponder {
			return session.GetSessionNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *app.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.SessionPostSessionHandler == nil {
		api.SessionPostSessionHandler = session.PostSessionHandlerFunc(func(params session.PostSessionParams, principal *app.Auth) session.PostSessionResponder {
			return session.PostSessionNotImplemented()
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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
