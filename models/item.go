// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item item
// swagger:model item
type Item struct {

	// itemname
	// Required: true
	Itemname *string `json:"itemname"`

	// price
	Price int64 `json:"price,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateItemname(formats strfmt.Registry) error {

	if err := validate.Required("itemname", "body", m.Itemname); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
