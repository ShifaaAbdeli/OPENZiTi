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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListServiceTerminatorsHandlerFunc turns a function with the right signature into a list service terminators handler
type ListServiceTerminatorsHandlerFunc func(ListServiceTerminatorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListServiceTerminatorsHandlerFunc) Handle(params ListServiceTerminatorsParams) middleware.Responder {
	return fn(params)
}

// ListServiceTerminatorsHandler interface for that can handle valid list service terminators params
type ListServiceTerminatorsHandler interface {
	Handle(ListServiceTerminatorsParams) middleware.Responder
}

// NewListServiceTerminators creates a new http.Handler for the list service terminators operation
func NewListServiceTerminators(ctx *middleware.Context, handler ListServiceTerminatorsHandler) *ListServiceTerminators {
	return &ListServiceTerminators{Context: ctx, Handler: handler}
}

/* ListServiceTerminators swagger:route GET /services/{id}/terminators Service listServiceTerminators

List of terminators assigned to a service

Retrieves a list of terminator resources that are assigned specific service; supports filtering, sorting, and pagination.


*/
type ListServiceTerminators struct {
	Context *middleware.Context
	Handler ListServiceTerminatorsHandler
}

func (o *ListServiceTerminators) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListServiceTerminatorsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
