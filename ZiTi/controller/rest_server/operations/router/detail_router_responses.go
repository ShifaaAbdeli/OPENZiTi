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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/ziti/controller/rest_model"
)

// DetailRouterOKCode is the HTTP code returned for type DetailRouterOK
const DetailRouterOKCode int = 200

/*DetailRouterOK A single router

swagger:response detailRouterOK
*/
type DetailRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailRouterEnvelope `json:"body,omitempty"`
}

// NewDetailRouterOK creates DetailRouterOK with default headers values
func NewDetailRouterOK() *DetailRouterOK {

	return &DetailRouterOK{}
}

// WithPayload adds the payload to the detail router o k response
func (o *DetailRouterOK) WithPayload(payload *rest_model.DetailRouterEnvelope) *DetailRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail router o k response
func (o *DetailRouterOK) SetPayload(payload *rest_model.DetailRouterEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailRouterUnauthorizedCode is the HTTP code returned for type DetailRouterUnauthorized
const DetailRouterUnauthorizedCode int = 401

/*DetailRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailRouterUnauthorized
*/
type DetailRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailRouterUnauthorized creates DetailRouterUnauthorized with default headers values
func NewDetailRouterUnauthorized() *DetailRouterUnauthorized {

	return &DetailRouterUnauthorized{}
}

// WithPayload adds the payload to the detail router unauthorized response
func (o *DetailRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail router unauthorized response
func (o *DetailRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailRouterNotFoundCode is the HTTP code returned for type DetailRouterNotFound
const DetailRouterNotFoundCode int = 404

/*DetailRouterNotFound The requested resource does not exist

swagger:response detailRouterNotFound
*/
type DetailRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailRouterNotFound creates DetailRouterNotFound with default headers values
func NewDetailRouterNotFound() *DetailRouterNotFound {

	return &DetailRouterNotFound{}
}

// WithPayload adds the payload to the detail router not found response
func (o *DetailRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail router not found response
func (o *DetailRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailRouterTooManyRequestsCode is the HTTP code returned for type DetailRouterTooManyRequests
const DetailRouterTooManyRequestsCode int = 429

/*DetailRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailRouterTooManyRequests
*/
type DetailRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailRouterTooManyRequests creates DetailRouterTooManyRequests with default headers values
func NewDetailRouterTooManyRequests() *DetailRouterTooManyRequests {

	return &DetailRouterTooManyRequests{}
}

// WithPayload adds the payload to the detail router too many requests response
func (o *DetailRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail router too many requests response
func (o *DetailRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
