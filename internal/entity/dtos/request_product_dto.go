// Code generated by go-swagger; DO NOT EDIT.

package dtos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestProductDto request product dto
//
// swagger:model RequestProductDto
type RequestProductDto struct {

	// brand
	// Example: Sample Brand
	// Required: true
	Brand *string `json:"brand"`

	// categories
	// Example: [book1, book2, book3]
	// Required: true
	Categories []*RequestCategoryDto `json:"categories"`

	// description
	// Example: This is a sample description
	// Required: true
	Description *string `json:"description"`

	// id
	// Example: 62995601-15f3-40bc-80c8-900f62116000
	// Required: true
	ID *string `json:"id"`

	// name
	// Example: Brand X Man Jean
	// Required: true
	Name *string `json:"name"`

	// price
	// Example: 19.99
	// Required: true
	Price *float64 `json:"price"`

	// stock
	// Example: 100
	// Required: true
	Stock *int64 `json:"stock"`
}

// Validate validates this request product dto
func (m *RequestProductDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestProductDto) validateBrand(formats strfmt.Registry) error {

	if err := validate.Required("brand", "body", m.Brand); err != nil {
		return err
	}

	return nil
}

func (m *RequestProductDto) validateCategories(formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequestProductDto) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *RequestProductDto) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RequestProductDto) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RequestProductDto) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *RequestProductDto) validateStock(formats strfmt.Registry) error {

	if err := validate.Required("stock", "body", m.Stock); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this request product dto based on the context it is used
func (m *RequestProductDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestProductDto) contextValidateCategories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Categories); i++ {

		if m.Categories[i] != nil {
			if err := m.Categories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestProductDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestProductDto) UnmarshalBinary(b []byte) error {
	var res RequestProductDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
