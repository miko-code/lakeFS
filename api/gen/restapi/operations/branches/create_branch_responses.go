// Code generated by go-swagger; DO NOT EDIT.

package branches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/treeverse/lakefs/api/gen/models"
)

// CreateBranchCreatedCode is the HTTP code returned for type CreateBranchCreated
const CreateBranchCreatedCode int = 201

/*CreateBranchCreated branch

swagger:response createBranchCreated
*/
type CreateBranchCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Ref `json:"body,omitempty"`
}

// NewCreateBranchCreated creates CreateBranchCreated with default headers values
func NewCreateBranchCreated() *CreateBranchCreated {

	return &CreateBranchCreated{}
}

// WithPayload adds the payload to the create branch created response
func (o *CreateBranchCreated) WithPayload(payload *models.Ref) *CreateBranchCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create branch created response
func (o *CreateBranchCreated) SetPayload(payload *models.Ref) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBranchCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBranchBadRequestCode is the HTTP code returned for type CreateBranchBadRequest
const CreateBranchBadRequestCode int = 400

/*CreateBranchBadRequest validation error

swagger:response createBranchBadRequest
*/
type CreateBranchBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBranchBadRequest creates CreateBranchBadRequest with default headers values
func NewCreateBranchBadRequest() *CreateBranchBadRequest {

	return &CreateBranchBadRequest{}
}

// WithPayload adds the payload to the create branch bad request response
func (o *CreateBranchBadRequest) WithPayload(payload *models.Error) *CreateBranchBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create branch bad request response
func (o *CreateBranchBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBranchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBranchUnauthorizedCode is the HTTP code returned for type CreateBranchUnauthorized
const CreateBranchUnauthorizedCode int = 401

/*CreateBranchUnauthorized Unauthorized

swagger:response createBranchUnauthorized
*/
type CreateBranchUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBranchUnauthorized creates CreateBranchUnauthorized with default headers values
func NewCreateBranchUnauthorized() *CreateBranchUnauthorized {

	return &CreateBranchUnauthorized{}
}

// WithPayload adds the payload to the create branch unauthorized response
func (o *CreateBranchUnauthorized) WithPayload(payload *models.Error) *CreateBranchUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create branch unauthorized response
func (o *CreateBranchUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBranchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateBranchDefault generic error response

swagger:response createBranchDefault
*/
type CreateBranchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBranchDefault creates CreateBranchDefault with default headers values
func NewCreateBranchDefault(code int) *CreateBranchDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBranchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create branch default response
func (o *CreateBranchDefault) WithStatusCode(code int) *CreateBranchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create branch default response
func (o *CreateBranchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create branch default response
func (o *CreateBranchDefault) WithPayload(payload *models.Error) *CreateBranchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create branch default response
func (o *CreateBranchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBranchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
