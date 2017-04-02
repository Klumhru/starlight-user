package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/klumhru/starlight-user/models"
)

// UpdateUserAcceptedCode is the HTTP code returned for type UpdateUserAccepted
const UpdateUserAcceptedCode int = 202

/*UpdateUserAccepted Updated

swagger:response updateUserAccepted
*/
type UpdateUserAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUpdateUserAccepted creates UpdateUserAccepted with default headers values
func NewUpdateUserAccepted() *UpdateUserAccepted {
	return &UpdateUserAccepted{}
}

// WithPayload adds the payload to the update user accepted response
func (o *UpdateUserAccepted) WithPayload(payload *models.User) *UpdateUserAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user accepted response
func (o *UpdateUserAccepted) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateUserDefault Error

swagger:response updateUserDefault
*/
type UpdateUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserDefault creates UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user default response
func (o *UpdateUserDefault) WithStatusCode(code int) *UpdateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user default response
func (o *UpdateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user default response
func (o *UpdateUserDefault) WithPayload(payload *models.Error) *UpdateUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user default response
func (o *UpdateUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}