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

// SpaceBase space base
// swagger:model SpaceBase
type SpaceBase struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Children Count
	// Read Only: true
	ChildCount int64 `json:"child_count,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// User Id of Creator
	// Read Only: true
	CreatorID int64 `json:"creator_id,omitempty"`

	// Embedder's Id if this space was autogenerated as an embedding shared space via 'external_group_id' in an SSO embed login
	// Read Only: true
	ExternalID string `json:"external_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Space is an embed space
	// Read Only: true
	IsEmbed *bool `json:"is_embed,omitempty"`

	// Space is the root embed shared space
	// Read Only: true
	IsEmbedSharedRoot *bool `json:"is_embed_shared_root,omitempty"`

	// Space is the root embed users space
	// Read Only: true
	IsEmbedUsersRoot *bool `json:"is_embed_users_root,omitempty"`

	// Space is a user's personal space
	// Read Only: true
	IsPersonal *bool `json:"is_personal,omitempty"`

	// Space is descendant of a user's personal space
	// Read Only: true
	IsPersonalDescendant *bool `json:"is_personal_descendant,omitempty"`

	// (DEPRECATED) Space is the root shared space (alias of is_shared_root)
	// Read Only: true
	IsRoot *bool `json:"is_root,omitempty"`

	// Space is the root shared space
	// Read Only: true
	IsSharedRoot *bool `json:"is_shared_root,omitempty"`

	// (DEPRECATED) Space is the root user space (alias of is_users_root
	// Read Only: true
	IsUserRoot *bool `json:"is_user_root,omitempty"`

	// Space is the root user space
	// Read Only: true
	IsUsersRoot *bool `json:"is_users_root,omitempty"`

	// Unique Name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Id of Parent
	// Required: true
	ParentID *int64 `json:"parent_id"`
}

// Validate validates this space base
func (m *SpaceBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpaceBase) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpaceBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpaceBase) UnmarshalBinary(b []byte) error {
	var res SpaceBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
