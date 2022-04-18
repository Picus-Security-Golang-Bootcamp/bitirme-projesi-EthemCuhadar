// Code generated by go-swagger; DO NOT EDIT.

package dtos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestCategoryDto request category dto
//
// swagger:model RequestCategoryDto
type RequestCategoryDto struct {

	// name
	// Example: Men Jeans
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this request category dto
func (m *RequestCategoryDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestCategoryDto) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this request category dto based on context it is used
func (m *RequestCategoryDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestCategoryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestCategoryDto) UnmarshalBinary(b []byte) error {
	var res RequestCategoryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
