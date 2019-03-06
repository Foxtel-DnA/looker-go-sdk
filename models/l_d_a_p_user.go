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

// LDAPUser l d a p user
// swagger:model LDAPUser
type LDAPUser struct {

	// Array of user's email addresses and aliases for use in migration
	// Read Only: true
	AllEmails []string `json:"all_emails"`

	// Dictionary of user's attrributes (name/value)
	// Read Only: true
	Attributes map[string]string `json:"attributes,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Primary email address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// First name
	// Read Only: true
	FirstName string `json:"first_name,omitempty"`

	// Array of user's groups (group names only)
	// Read Only: true
	Groups []string `json:"groups"`

	// Last Name
	// Read Only: true
	LastName string `json:"last_name,omitempty"`

	// LDAP's distinguished name for the user record
	// Read Only: true
	LdapDn string `json:"ldap_dn,omitempty"`

	// LDAP's Unique ID for the user
	// Read Only: true
	LdapID string `json:"ldap_id,omitempty"`

	// Array of user's roles (role names only)
	// Read Only: true
	Roles []string `json:"roles"`

	// Link to ldap config
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this l d a p user
func (m *LDAPUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPUser) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPUser) UnmarshalBinary(b []byte) error {
	var res LDAPUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
