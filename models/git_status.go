// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GitStatus git status
// swagger:model GitStatus
type GitStatus struct {

	// Git action: add, delete, etc
	// Read Only: true
	Action string `json:"action,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// When true, changes to the local file conflict with the remote repository
	// Read Only: true
	Conflict *bool `json:"conflict,omitempty"`

	// When true, the file can be reverted to an earlier state
	// Read Only: true
	Revertable *bool `json:"revertable,omitempty"`

	// Git description of the action
	// Read Only: true
	Text string `json:"text,omitempty"`
}

// Validate validates this git status
func (m *GitStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitStatus) UnmarshalBinary(b []byte) error {
	var res GitStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}