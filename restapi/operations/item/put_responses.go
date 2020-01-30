// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ASinha24/supermartswaggerAPI/models"
)

// PutOKCode is the HTTP code returned for type PutOK
const PutOKCode int = 200

/*PutOK item updated

swagger:response putOK
*/
type PutOK struct {

	/*
	  In: Body
	*/
	Payload *models.ItemResponse `json:"body,omitempty"`
}

// NewPutOK creates PutOK with default headers values
func NewPutOK() *PutOK {

	return &PutOK{}
}

// WithPayload adds the payload to the put o k response
func (o *PutOK) WithPayload(payload *models.ItemResponse) *PutOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put o k response
func (o *PutOK) SetPayload(payload *models.ItemResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}