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

package link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListLinksHandlerFunc turns a function with the right signature into a list links handler
type ListLinksHandlerFunc func(ListLinksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListLinksHandlerFunc) Handle(params ListLinksParams) middleware.Responder {
	return fn(params)
}

// ListLinksHandler interface for that can handle valid list links params
type ListLinksHandler interface {
	Handle(ListLinksParams) middleware.Responder
}

// NewListLinks creates a new http.Handler for the list links operation
func NewListLinks(ctx *middleware.Context, handler ListLinksHandler) *ListLinks {
	return &ListLinks{Context: ctx, Handler: handler}
}

/* ListLinks swagger:route GET /links Link listLinks

List links

Retrieves a list of link resources; does not supports filtering, sorting, or pagination. Requires admin access.


*/
type ListLinks struct {
	Context *middleware.Context
	Handler ListLinksHandler
}

func (o *ListLinks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListLinksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
