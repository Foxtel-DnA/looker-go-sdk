// Code generated by go-swagger; DO NOT EDIT.

package models

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

// SamlConfig saml config
//
// swagger:model SamlConfig
type SamlConfig struct {

	// Allows roles to be directly assigned to SAML auth'd users.
	AllowDirectRoles bool `json:"allow_direct_roles,omitempty"`

	// Allow SAML auth'd users to be members of non-reflected Looker groups. If 'false', user will be removed from non-reflected groups on login.
	AllowNormalGroupMembership bool `json:"allow_normal_group_membership,omitempty"`

	// SAML auth'd users will inherit roles from non-reflected Looker groups.
	AllowRolesFromNormalGroups bool `json:"allow_roles_from_normal_groups,omitempty"`

	// Count of seconds of clock drift to allow when validating timestamps of assertions.
	AllowedClockDrift int64 `json:"allowed_clock_drift,omitempty"`

	// Allow alternate email-based login via '/login/email' for admins and for specified users with the 'login_special_email' permission. This option is useful as a fallback during ldap setup, if ldap config problems occur later, or if you need to support some users who are not in your ldap directory. Looker email/password logins are always disabled for regular users when ldap is enabled.
	AlternateEmailLoginAllowed bool `json:"alternate_email_login_allowed,omitempty"`

	// Users will not be allowed to login at all unless a role for them is found in Saml if set to true
	AuthRequiresRole bool `json:"auth_requires_role,omitempty"`

	// Bypass the login page when user authentication is required. Redirect to IdP immediately instead.
	BypassLoginPage bool `json:"bypass_login_page,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// (Write-Only) Array of ids of groups that will be applied to new users the first time they login via Saml
	DefaultNewUserGroupIds []int64 `json:"default_new_user_group_ids"`

	// (Read-only) Groups that will be applied to new users the first time they login via Saml
	// Read Only: true
	DefaultNewUserGroups []*Group `json:"default_new_user_groups"`

	// (Write-Only) Array of ids of roles that will be applied to new users the first time they login via Saml
	DefaultNewUserRoleIds []int64 `json:"default_new_user_role_ids"`

	// (Read-only) Roles that will be applied to new users the first time they login via Saml
	// Read Only: true
	DefaultNewUserRoles []*Role `json:"default_new_user_roles"`

	// Enable/Disable Saml authentication for the server
	Enabled bool `json:"enabled,omitempty"`

	// (Read-only) Array of mappings between Saml Groups and Looker Roles
	// Read Only: true
	Groups []*SamlGroupRead `json:"groups"`

	// Name of user record attributes used to indicate groups. Used when 'groups_finder_type' is set to 'grouped_attribute_values'
	GroupsAttribute string `json:"groups_attribute,omitempty"`

	// Identifier for a strategy for how Looker will find groups in the SAML response. One of ['grouped_attribute_values', 'individual_attributes']
	GroupsFinderType string `json:"groups_finder_type,omitempty"`

	// Value for group attribute used to indicate membership. Used when 'groups_finder_type' is set to 'individual_attributes'
	GroupsMemberValue string `json:"groups_member_value,omitempty"`

	// (Read/Write) Array of mappings between Saml Groups and arrays of Looker Role ids
	GroupsWithRoleIds []*SamlGroupWrite `json:"groups_with_role_ids"`

	// Identity Provider Audience (set in IdP config). Optional in Looker. Set this only if you want Looker to validate the audience value returned by the IdP.
	IdpAudience string `json:"idp_audience,omitempty"`

	// Identity Provider Certificate (provided by IdP)
	IdpCert string `json:"idp_cert,omitempty"`

	// Identity Provider Issuer (provided by IdP)
	IdpIssuer string `json:"idp_issuer,omitempty"`

	// Identity Provider Url (provided by IdP)
	IdpURL string `json:"idp_url,omitempty"`

	// When this config was last modified
	// Read Only: true
	ModifiedAt string `json:"modified_at,omitempty"`

	// User id of user who last modified this config
	// Read Only: true
	ModifiedBy string `json:"modified_by,omitempty"`

	// Merge first-time saml login to existing user account by email addresses. When a user logs in for the first time via saml this option will connect this user into their existing account by finding the account with a matching email address by testing the given types of credentials for existing users. Otherwise a new user account will be created for the user. This list (if provided) must be a comma separated list of string like 'email,ldap,google'
	NewUserMigrationTypes string `json:"new_user_migration_types,omitempty"`

	// Set user roles in Looker based on groups from Saml
	SetRolesFromGroups bool `json:"set_roles_from_groups,omitempty"`

	// Slug to identify configurations that are created in order to run a Saml config test
	// Read Only: true
	TestSlug string `json:"test_slug,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Name of user record attributes used to indicate email address field
	UserAttributeMapEmail string `json:"user_attribute_map_email,omitempty"`

	// Name of user record attributes used to indicate first name
	UserAttributeMapFirstName string `json:"user_attribute_map_first_name,omitempty"`

	// Name of user record attributes used to indicate last name
	UserAttributeMapLastName string `json:"user_attribute_map_last_name,omitempty"`

	// (Read-only) Array of mappings between Saml User Attributes and Looker User Attributes
	// Read Only: true
	UserAttributes []*SamlUserAttributeRead `json:"user_attributes"`

	// (Read/Write) Array of mappings between Saml User Attributes and arrays of Looker User Attribute ids
	UserAttributesWithIds []*SamlUserAttributeWrite `json:"user_attributes_with_ids"`
}

// Validate validates this saml config
func (m *SamlConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultNewUserGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultNewUserRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupsWithRoleIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAttributesWithIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlConfig) validateDefaultNewUserGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultNewUserGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultNewUserGroups); i++ {
		if swag.IsZero(m.DefaultNewUserGroups[i]) { // not required
			continue
		}

		if m.DefaultNewUserGroups[i] != nil {
			if err := m.DefaultNewUserGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("default_new_user_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) validateDefaultNewUserRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultNewUserRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultNewUserRoles); i++ {
		if swag.IsZero(m.DefaultNewUserRoles[i]) { // not required
			continue
		}

		if m.DefaultNewUserRoles[i] != nil {
			if err := m.DefaultNewUserRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("default_new_user_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) validateGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) validateGroupsWithRoleIds(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupsWithRoleIds) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupsWithRoleIds); i++ {
		if swag.IsZero(m.GroupsWithRoleIds[i]) { // not required
			continue
		}

		if m.GroupsWithRoleIds[i] != nil {
			if err := m.GroupsWithRoleIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups_with_role_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SamlConfig) validateUserAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.UserAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.UserAttributes); i++ {
		if swag.IsZero(m.UserAttributes[i]) { // not required
			continue
		}

		if m.UserAttributes[i] != nil {
			if err := m.UserAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) validateUserAttributesWithIds(formats strfmt.Registry) error {
	if swag.IsZero(m.UserAttributesWithIds) { // not required
		return nil
	}

	for i := 0; i < len(m.UserAttributesWithIds); i++ {
		if swag.IsZero(m.UserAttributesWithIds[i]) { // not required
			continue
		}

		if m.UserAttributesWithIds[i] != nil {
			if err := m.UserAttributesWithIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_attributes_with_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this saml config based on the context it is used
func (m *SamlConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultNewUserGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultNewUserRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupsWithRoleIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTestSlug(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAttributesWithIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlConfig) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *SamlConfig) contextValidateDefaultNewUserGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default_new_user_groups", "body", []*Group(m.DefaultNewUserGroups)); err != nil {
		return err
	}

	for i := 0; i < len(m.DefaultNewUserGroups); i++ {

		if m.DefaultNewUserGroups[i] != nil {
			if err := m.DefaultNewUserGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("default_new_user_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) contextValidateDefaultNewUserRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default_new_user_roles", "body", []*Role(m.DefaultNewUserRoles)); err != nil {
		return err
	}

	for i := 0; i < len(m.DefaultNewUserRoles); i++ {

		if m.DefaultNewUserRoles[i] != nil {
			if err := m.DefaultNewUserRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("default_new_user_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groups", "body", []*SamlGroupRead(m.Groups)); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) contextValidateGroupsWithRoleIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupsWithRoleIds); i++ {

		if m.GroupsWithRoleIds[i] != nil {
			if err := m.GroupsWithRoleIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups_with_role_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) contextValidateModifiedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modified_at", "body", string(m.ModifiedAt)); err != nil {
		return err
	}

	return nil
}

func (m *SamlConfig) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modified_by", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *SamlConfig) contextValidateTestSlug(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "test_slug", "body", string(m.TestSlug)); err != nil {
		return err
	}

	return nil
}

func (m *SamlConfig) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *SamlConfig) contextValidateUserAttributes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_attributes", "body", []*SamlUserAttributeRead(m.UserAttributes)); err != nil {
		return err
	}

	for i := 0; i < len(m.UserAttributes); i++ {

		if m.UserAttributes[i] != nil {
			if err := m.UserAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SamlConfig) contextValidateUserAttributesWithIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserAttributesWithIds); i++ {

		if m.UserAttributesWithIds[i] != nil {
			if err := m.UserAttributesWithIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_attributes_with_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SamlConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlConfig) UnmarshalBinary(b []byte) error {
	var res SamlConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
