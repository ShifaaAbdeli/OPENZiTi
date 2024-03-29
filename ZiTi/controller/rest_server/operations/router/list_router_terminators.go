// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRouterTerminatorsHandlerFunc turns a function with the right signature into a list router terminators handler
type ListRouterTerminatorsHandlerFunc func(ListRouterTerminatorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRouterTerminatorsHandlerFunc) Handle(params ListRouterTerminatorsParams) middleware.Responder {
	return fn(params)
}

// ListRouterTerminatorsHandler interface for that can handle valid list router terminators params
type ListRouterTerminatorsHandler interface {
	Handle(ListRouterTerminatorsParams) middleware.Responder
}

// NewListRouterTerminators creates a new http.Handler for the list router terminators operation
func NewListRouterTerminators(ctx *middleware.Context, handler ListRouterTerminatorsHandler) *ListRouterTerminators {
	return &ListRouterTerminators{Context: ctx, Handler: handler}
}

/* ListRouterTerminators swagger:route GET /routers/{id}/terminators Router listRouterTerminators

List of terminators assigned to a router

Retrieves a list of terminator resources that are assigned specific router; supports filtering, sorting, and pagination.


*/
type ListRouterTerminators struct {
	Context *middleware.Context
	Handler ListRouterTerminatorsHandler
}

func (o *ListRouterTerminators) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListRouterTerminatorsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
