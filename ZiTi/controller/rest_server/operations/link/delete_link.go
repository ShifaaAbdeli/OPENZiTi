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

// DeleteLinkHandlerFunc turns a function with the right signature into a delete link handler
type DeleteLinkHandlerFunc func(DeleteLinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteLinkHandlerFunc) Handle(params DeleteLinkParams) middleware.Responder {
	return fn(params)
}

// DeleteLinkHandler interface for that can handle valid delete link params
type DeleteLinkHandler interface {
	Handle(DeleteLinkParams) middleware.Responder
}

// NewDeleteLink creates a new http.Handler for the delete link operation
func NewDeleteLink(ctx *middleware.Context, handler DeleteLinkHandler) *DeleteLink {
	return &DeleteLink{Context: ctx, Handler: handler}
}

/* DeleteLink swagger:route DELETE /links/{id} Link deleteLink

Delete a link

Delete a link by id. Requires admin access.

*/
type DeleteLink struct {
	Context *middleware.Context
	Handler DeleteLinkHandler
}

func (o *DeleteLink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteLinkParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
