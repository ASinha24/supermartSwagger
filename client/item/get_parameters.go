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
)

// NewGetParams creates a new GetParams object
// with the default values initialized.
func NewGetParams() *GetParams {
	var ()
	return &GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParamsWithTimeout creates a new GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParamsWithTimeout(timeout time.Duration) *GetParams {
	var ()
	return &GetParams{

		timeout: timeout,
	}
}

// NewGetParamsWithContext creates a new GetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParamsWithContext(ctx context.Context) *GetParams {
	var ()
	return &GetParams{

		Context: ctx,
	}
}

// NewGetParamsWithHTTPClient creates a new GetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParamsWithHTTPClient(client *http.Client) *GetParams {
	var ()
	return &GetParams{
		HTTPClient: client,
	}
}

/*GetParams contains all the parameters to send to the API endpoint
for the get operation typically these are written to a http.Request
*/
type GetParams struct {

	/*ItemID*/
	ItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get params
func (o *GetParams) WithTimeout(timeout time.Duration) *GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get params
func (o *GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get params
func (o *GetParams) WithContext(ctx context.Context) *GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get params
func (o *GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get params
func (o *GetParams) WithHTTPClient(client *http.Client) *GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get params
func (o *GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the get params
func (o *GetParams) WithItemID(itemID string) *GetParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get params
func (o *GetParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
