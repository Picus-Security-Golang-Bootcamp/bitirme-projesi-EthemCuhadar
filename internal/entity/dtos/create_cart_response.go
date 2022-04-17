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

// CreateCartResponse create cart response
//
// swagger:model CreateCartResponse
type CreateCartResponse struct {

	// cancel time
	// Example: 2022-04-17T07:00:27.677Z
	// Required: true
	// Format: date
	CancelTime *strfmt.Date `json:"cancelTime"`

	// id
	// Example: 62995601-15f3-40bc-80c8-900f62116000
	// Required: true
	ID *string `json:"id"`

	// is ordered
	// Example: true
	// Required: true
	IsOrdered *bool `json:"isOrdered"`

	// item
	// Example: [item1, item2, item3]
	// Required: true
	Item []*Item `json:"item"`

	// order time
	// Example: 2022-04-17T07:00:27.677Z
	// Required: true
	// Format: date
	OrderTime *strfmt.Date `json:"orderTime"`

	// price
	// Example: 99.99
	// Required: true
	Price *float64 `json:"price"`

	// user Id
	// Example: 62995601-15f3-40bc-80c8-900f62116000
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this create cart response
func (m *CreateCartResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancelTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsOrdered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCartResponse) validateCancelTime(formats strfmt.Registry) error {

	if err := validate.Required("cancelTime", "body", m.CancelTime); err != nil {
		return err
	}

	if err := validate.FormatOf("cancelTime", "body", "date", m.CancelTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateCartResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CreateCartResponse) validateIsOrdered(formats strfmt.Registry) error {

	if err := validate.Required("isOrdered", "body", m.IsOrdered); err != nil {
		return err
	}

	return nil
}

func (m *CreateCartResponse) validateItem(formats strfmt.Registry) error {

	if err := validate.Required("item", "body", m.Item); err != nil {
		return err
	}

	for i := 0; i < len(m.Item); i++ {
		if swag.IsZero(m.Item[i]) { // not required
			continue
		}

		if m.Item[i] != nil {
			if err := m.Item[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("item" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateCartResponse) validateOrderTime(formats strfmt.Registry) error {

	if err := validate.Required("orderTime", "body", m.OrderTime); err != nil {
		return err
	}

	if err := validate.FormatOf("orderTime", "body", "date", m.OrderTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateCartResponse) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *CreateCartResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create cart response based on the context it is used
func (m *CreateCartResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCartResponse) contextValidateItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Item); i++ {

		if m.Item[i] != nil {
			if err := m.Item[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("item" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCartResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCartResponse) UnmarshalBinary(b []byte) error {
	var res CreateCartResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
