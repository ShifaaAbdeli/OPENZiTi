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

// CreateRouterCreatedCode is the HTTP code returned for type CreateRouterCreated
const CreateRouterCreatedCode int = 201

/*CreateRouterCreated The create request was successful and the resource has been added at the following location

swagger:response createRouterCreated
*/
type CreateRouterCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateRouterCreated creates CreateRouterCreated with default headers values
func NewCreateRouterCreated() *CreateRouterCreated {

	return &CreateRouterCreated{}
}

// WithPayload adds the payload to the create router created response
func (o *CreateRouterCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateRouterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create router created response
func (o *CreateRouterCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRouterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRouterBadRequestCode is the HTTP code returned for type CreateRouterBadRequest
const CreateRouterBadRequestCode int = 400

/*CreateRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createRouterBadRequest
*/
type CreateRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateRouterBadRequest creates CreateRouterBadRequest with default headers values
func NewCreateRouterBadRequest() *CreateRouterBadRequest {

	return &CreateRouterBadRequest{}
}

// WithPayload adds the payload to the create router bad request response
func (o *CreateRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create router bad request response
func (o *CreateRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRouterUnauthorizedCode is the HTTP code returned for type CreateRouterUnauthorized
const CreateRouterUnauthorizedCode int = 401

/*CreateRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createRouterUnauthorized
*/
type CreateRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateRouterUnauthorized creates CreateRouterUnauthorized with default headers values
func NewCreateRouterUnauthorized() *CreateRouterUnauthorized {

	return &CreateRouterUnauthorized{}
}

// WithPayload adds the payload to the create router unauthorized response
func (o *CreateRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create router unauthorized response
func (o *CreateRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRouterTooManyRequestsCode is the HTTP code returned for type CreateRouterTooManyRequests
const CreateRouterTooManyRequestsCode int = 429

/*CreateRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createRouterTooManyRequests
*/
type CreateRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateRouterTooManyRequests creates CreateRouterTooManyRequests with default headers values
func NewCreateRouterTooManyRequests() *CreateRouterTooManyRequests {

	return &CreateRouterTooManyRequests{}
}

// WithPayload adds the payload to the create router too many requests response
func (o *CreateRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create router too many requests response
func (o *CreateRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
