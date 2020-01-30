// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ASinha24/supermartswaggerAPI/models"
)

// NewPutParams creates a new PutParams object
// with the default values initialized.
func NewPutParams() *PutParams {
	var ()
	return &PutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutParamsWithTimeout creates a new PutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutParamsWithTimeout(timeout time.Duration) *PutParams {
	var ()
	return &PutParams{

		timeout: timeout,
	}
}

// NewPutParamsWithContext creates a new PutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutParamsWithContext(ctx context.Context) *PutParams {
	var ()
	return &PutParams{

		Context: ctx,
	}
}

// NewPutParamsWithHTTPClient creates a new PutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutParamsWithHTTPClient(client *http.Client) *PutParams {
	var ()
	return &PutParams{
		HTTPClient: client,
	}
}

/*PutParams contains all the parameters to send to the API endpoint
for the put operation typically these are written to a http.Request
*/
type PutParams struct {

	/*Item*/
	Item *models.Item
	/*ItemID*/
	ItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put params
func (o *PutParams) WithTimeout(timeout time.Duration) *PutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put params
func (o *PutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put params
func (o *PutParams) WithContext(ctx context.Context) *PutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put params
func (o *PutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put params
func (o *PutParams) WithHTTPClient(client *http.Client) *PutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put params
func (o *PutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItem adds the item to the put params
func (o *PutParams) WithItem(item *models.Item) *PutParams {
	o.SetItem(item)
	return o
}

// SetItem adds the item to the put params
func (o *PutParams) SetItem(item *models.Item) {
	o.Item = item
}

// WithItemID adds the itemID to the put params
func (o *PutParams) WithItemID(itemID string) *PutParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the put params
func (o *PutParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *PutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Item != nil {
		if err := r.SetBodyParam(o.Item); err != nil {
			return err
		}
	}

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
