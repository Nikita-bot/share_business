// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washBonus/internal/app"
)

// UpdateServerGroupHandlerFunc turns a function with the right signature into a update server group handler
type UpdateServerGroupHandlerFunc func(UpdateServerGroupParams, *app.Auth) UpdateServerGroupResponder

// Handle executing the request and returning a response
func (fn UpdateServerGroupHandlerFunc) Handle(params UpdateServerGroupParams, principal *app.Auth) UpdateServerGroupResponder {
	return fn(params, principal)
}

// UpdateServerGroupHandler interface for that can handle valid update server group params
type UpdateServerGroupHandler interface {
	Handle(UpdateServerGroupParams, *app.Auth) UpdateServerGroupResponder
}

// NewUpdateServerGroup creates a new http.Handler for the update server group operation
func NewUpdateServerGroup(ctx *middleware.Context, handler UpdateServerGroupHandler) *UpdateServerGroup {
	return &UpdateServerGroup{Context: ctx, Handler: handler}
}

/*
	UpdateServerGroup swagger:route PATCH /server-groups/{groupId} serverGroups updateServerGroup

UpdateServerGroup update server group API
*/
type UpdateServerGroup struct {
	Context *middleware.Context
	Handler UpdateServerGroupHandler
}

func (o *UpdateServerGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateServerGroupParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *app.Auth
	if uprinc != nil {
		principal = uprinc.(*app.Auth) // this is really a app.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
