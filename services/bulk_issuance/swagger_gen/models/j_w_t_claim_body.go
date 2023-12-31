// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JWTClaimBody j w t claim body
//
// swagger:model JWTClaimBody
type JWTClaimBody struct {

	// preferred username
	PreferredUsername string `json:"preferred_username,omitempty"`

	// realm access
	RealmAccess *JWTClaimBodyRealmAccess `json:"realm_access,omitempty"`

	// resource access
	ResourceAccess map[string]Group `json:"resource_access,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this j w t claim body
func (m *JWTClaimBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRealmAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JWTClaimBody) validateRealmAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.RealmAccess) { // not required
		return nil
	}

	if m.RealmAccess != nil {
		if err := m.RealmAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realm_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realm_access")
			}
			return err
		}
	}

	return nil
}

func (m *JWTClaimBody) validateResourceAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceAccess) { // not required
		return nil
	}

	for k := range m.ResourceAccess {

		if err := validate.Required("resource_access"+"."+k, "body", m.ResourceAccess[k]); err != nil {
			return err
		}
		if val, ok := m.ResourceAccess[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource_access" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resource_access" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this j w t claim body based on the context it is used
func (m *JWTClaimBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRealmAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JWTClaimBody) contextValidateRealmAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.RealmAccess != nil {
		if err := m.RealmAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realm_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realm_access")
			}
			return err
		}
	}

	return nil
}

func (m *JWTClaimBody) contextValidateResourceAccess(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ResourceAccess {

		if val, ok := m.ResourceAccess[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JWTClaimBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JWTClaimBody) UnmarshalBinary(b []byte) error {
	var res JWTClaimBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JWTClaimBodyRealmAccess j w t claim body realm access
//
// swagger:model JWTClaimBodyRealmAccess
type JWTClaimBodyRealmAccess struct {

	// roles
	Roles []string `json:"roles"`
}

// Validate validates this j w t claim body realm access
func (m *JWTClaimBodyRealmAccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this j w t claim body realm access based on context it is used
func (m *JWTClaimBodyRealmAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JWTClaimBodyRealmAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JWTClaimBodyRealmAccess) UnmarshalBinary(b []byte) error {
	var res JWTClaimBodyRealmAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
