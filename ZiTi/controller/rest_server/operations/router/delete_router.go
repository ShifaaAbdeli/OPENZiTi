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

// DeleteRouterHandlerFunc turns a function with the right signature into a delete router handler
type DeleteRouterHandlerFunc func(DeleteRouterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRouterHandlerFunc) Handle(params DeleteRouterParams) middleware.Responder {
	return fn(params)
}

// DeleteRouterHandler interface for that can handle valid delete router params
type DeleteRouterHandler interface {
	Handle(DeleteRouterParams) middleware.Responder
}

// NewDeleteRouter creates a new http.Handler for the delete router operation
func NewDeleteRouter(ctx *middleware.Context, handler DeleteRouterHandler) *DeleteRouter {
	return &DeleteRouter{Context: ctx, Handler: handler}
}

/* DeleteRouter swagger:route DELETE /routers/{id} Router deleteRouter

Delete a router

Delete a router by id. Requires admin access.

*/
type DeleteRouter struct {
	Context *middleware.Context
	Handler DeleteRouterHandler
}

func (o *DeleteRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteRouterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
