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

// User user
//
// swagger:model User
type User struct {

	// User can be directly assigned a role.
	// Read Only: true
	AllowDirectRoles *bool `json:"allow_direct_roles,omitempty"`

	// User can be a direct member of a normal Looker group.
	// Read Only: true
	AllowNormalGroupMembership *bool `json:"allow_normal_group_membership,omitempty"`

	// User can inherit roles from a normal Looker group.
	// Read Only: true
	AllowRolesFromNormalGroups *bool `json:"allow_roles_from_normal_groups,omitempty"`

	// URL for the avatar image (may be generic)
	// Read Only: true
	// Format: uri
	AvatarURL strfmt.URI `json:"avatar_url,omitempty"`

	// URL for the avatar image (may be generic), does not specify size
	// Read Only: true
	// Format: uri
	AvatarURLWithoutSizing strfmt.URI `json:"avatar_url_without_sizing,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// API 3 credentials
	// Read Only: true
	CredentialsApi3 []*CredentialsApi3 `json:"credentials_api3"`

	// Email/Password login credentials
	// Read Only: true
	CredentialsEmail *CredentialsEmail `json:"credentials_email,omitempty"`

	// Embed credentials
	// Read Only: true
	CredentialsEmbed []*CredentialsEmbed `json:"credentials_embed"`

	// Google auth credentials
	// Read Only: true
	CredentialsGoogle *CredentialsGoogle `json:"credentials_google,omitempty"`

	// LDAP credentials
	// Read Only: true
	CredentialsLdap *CredentialsLDAP `json:"credentials_ldap,omitempty"`

	// LookerOpenID credentials. Used for login by Looker Analysts
	// Read Only: true
	CredentialsLookerOpenid *CredentialsLookerOpenid `json:"credentials_looker_openid,omitempty"`

	// OpenID Connect auth credentials
	// Read Only: true
	CredentialsOidc *CredentialsOIDC `json:"credentials_oidc,omitempty"`

	// Saml auth credentials
	// Read Only: true
	CredentialsSaml *CredentialsSaml `json:"credentials_saml,omitempty"`

	// Two-factor credentials
	// Read Only: true
	CredentialsTotp *CredentialsTotp `json:"credentials_totp,omitempty"`

	// Full name for display (available only if both first_name and last_name are set)
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// EMail address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// (Embed only) ID of user's group space based on the external_group_id optionally specified during embed user login
	// Read Only: true
	EmbedGroupSpaceID int64 `json:"embed_group_space_id,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Array of ids of the groups for this user
	// Read Only: true
	GroupIds []int64 `json:"group_ids"`

	// ID string for user's home folder
	HomeFolderID string `json:"home_folder_id,omitempty"`

	// ID string for user's home space
	HomeSpaceID string `json:"home_space_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Account has been disabled
	IsDisabled bool `json:"is_disabled,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// User's preferred locale. User locale takes precedence over Looker's system-wide default locale. Locale determines language of display strings and date and numeric formatting in API responses. Locale string must be a 2 letter language code or a combination of language code and region code: 'en' or 'en-US', for example.
	Locale string `json:"locale,omitempty"`

	// Array of strings representing the Looker versions that this user has used (this only goes back as far as '3.54.0')
	// Read Only: true
	LookerVersions []string `json:"looker_versions"`

	// User's dev workspace has been checked for presence of applicable production projects
	ModelsDirValidated bool `json:"models_dir_validated,omitempty"`

	// ID of user's personal folder
	// Read Only: true
	PersonalFolderID int64 `json:"personal_folder_id,omitempty"`

	// ID of user's personal space
	// Read Only: true
	PersonalSpaceID int64 `json:"personal_space_id,omitempty"`

	// User is identified as an employee of Looker
	// Read Only: true
	PresumedLookerEmployee *bool `json:"presumed_looker_employee,omitempty"`

	// Array of ids of the roles for this user
	// Read Only: true
	RoleIds []int64 `json:"role_ids"`

	// User's roles are managed by an external directory like SAML or LDAP and can not be changed directly.
	// Read Only: true
	RolesExternallyManaged *bool `json:"roles_externally_managed,omitempty"`

	// Active sessions
	// Read Only: true
	Sessions []*Session `json:"sessions"`

	// Per user dictionary of undocumented state information owned by the Looker UI.
	UIState map[string]string `json:"ui_state,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// User is identified as an employee of Looker who has been verified via Looker corporate authentication
	// Read Only: true
	VerifiedLookerEmployee *bool `json:"verified_looker_employee,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatarURLWithoutSizing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsApi3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsEmbed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsGoogle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsLdap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsLookerOpenid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsSaml(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsTotp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAvatarURL(formats strfmt.Registry) error {
	if swag.IsZero(m.AvatarURL) { // not required
		return nil
	}

	if err := validate.FormatOf("avatar_url", "body", "uri", m.AvatarURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateAvatarURLWithoutSizing(formats strfmt.Registry) error {
	if swag.IsZero(m.AvatarURLWithoutSizing) { // not required
		return nil
	}

	if err := validate.FormatOf("avatar_url_without_sizing", "body", "uri", m.AvatarURLWithoutSizing.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCredentialsApi3(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsApi3) { // not required
		return nil
	}

	for i := 0; i < len(m.CredentialsApi3); i++ {
		if swag.IsZero(m.CredentialsApi3[i]) { // not required
			continue
		}

		if m.CredentialsApi3[i] != nil {
			if err := m.CredentialsApi3[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_api3" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateCredentialsEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsEmail) { // not required
		return nil
	}

	if m.CredentialsEmail != nil {
		if err := m.CredentialsEmail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_email")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsEmbed(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsEmbed) { // not required
		return nil
	}

	for i := 0; i < len(m.CredentialsEmbed); i++ {
		if swag.IsZero(m.CredentialsEmbed[i]) { // not required
			continue
		}

		if m.CredentialsEmbed[i] != nil {
			if err := m.CredentialsEmbed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_embed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateCredentialsGoogle(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsGoogle) { // not required
		return nil
	}

	if m.CredentialsGoogle != nil {
		if err := m.CredentialsGoogle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_google")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsLdap(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsLdap) { // not required
		return nil
	}

	if m.CredentialsLdap != nil {
		if err := m.CredentialsLdap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_ldap")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsLookerOpenid(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsLookerOpenid) { // not required
		return nil
	}

	if m.CredentialsLookerOpenid != nil {
		if err := m.CredentialsLookerOpenid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_looker_openid")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsOidc(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsOidc) { // not required
		return nil
	}

	if m.CredentialsOidc != nil {
		if err := m.CredentialsOidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_oidc")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsSaml(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsSaml) { // not required
		return nil
	}

	if m.CredentialsSaml != nil {
		if err := m.CredentialsSaml.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_saml")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsTotp(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsTotp) { // not required
		return nil
	}

	if m.CredentialsTotp != nil {
		if err := m.CredentialsTotp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_totp")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateSessions(formats strfmt.Registry) error {
	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowDirectRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAllowNormalGroupMembership(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAllowRolesFromNormalGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvatarURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvatarURLWithoutSizing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsApi3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsEmbed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsGoogle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsLdap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsLookerOpenid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsOidc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsSaml(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsTotp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmbedGroupSpaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLookerVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersonalFolderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersonalSpaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePresumedLookerEmployee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRolesExternallyManaged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerifiedLookerEmployee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateAllowDirectRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "allow_direct_roles", "body", m.AllowDirectRoles); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateAllowNormalGroupMembership(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "allow_normal_group_membership", "body", m.AllowNormalGroupMembership); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateAllowRolesFromNormalGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "allow_roles_from_normal_groups", "body", m.AllowRolesFromNormalGroups); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateAvatarURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "avatar_url", "body", strfmt.URI(m.AvatarURL)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateAvatarURLWithoutSizing(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "avatar_url_without_sizing", "body", strfmt.URI(m.AvatarURLWithoutSizing)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *User) contextValidateCredentialsApi3(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "credentials_api3", "body", []*CredentialsApi3(m.CredentialsApi3)); err != nil {
		return err
	}

	for i := 0; i < len(m.CredentialsApi3); i++ {

		if m.CredentialsApi3[i] != nil {
			if err := m.CredentialsApi3[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_api3" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateCredentialsEmail(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsEmail != nil {
		if err := m.CredentialsEmail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_email")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsEmbed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "credentials_embed", "body", []*CredentialsEmbed(m.CredentialsEmbed)); err != nil {
		return err
	}

	for i := 0; i < len(m.CredentialsEmbed); i++ {

		if m.CredentialsEmbed[i] != nil {
			if err := m.CredentialsEmbed[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_embed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateCredentialsGoogle(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsGoogle != nil {
		if err := m.CredentialsGoogle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_google")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsLdap(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsLdap != nil {
		if err := m.CredentialsLdap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_ldap")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsLookerOpenid(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsLookerOpenid != nil {
		if err := m.CredentialsLookerOpenid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_looker_openid")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsOidc(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsOidc != nil {
		if err := m.CredentialsOidc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_oidc")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsSaml(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsSaml != nil {
		if err := m.CredentialsSaml.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_saml")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateCredentialsTotp(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsTotp != nil {
		if err := m.CredentialsTotp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_totp")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "email", "body", string(m.Email)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateEmbedGroupSpaceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "embed_group_space_id", "body", int64(m.EmbedGroupSpaceID)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateGroupIds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "group_ids", "body", []int64(m.GroupIds)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateLookerVersions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "looker_versions", "body", []string(m.LookerVersions)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidatePersonalFolderID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "personal_folder_id", "body", int64(m.PersonalFolderID)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidatePersonalSpaceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "personal_space_id", "body", int64(m.PersonalSpaceID)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidatePresumedLookerEmployee(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "presumed_looker_employee", "body", m.PresumedLookerEmployee); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateRoleIds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "role_ids", "body", []int64(m.RoleIds)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateRolesExternallyManaged(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "roles_externally_managed", "body", m.RolesExternallyManaged); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateSessions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sessions", "body", []*Session(m.Sessions)); err != nil {
		return err
	}

	for i := 0; i < len(m.Sessions); i++ {

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *User) contextValidateVerifiedLookerEmployee(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "verified_looker_employee", "body", m.VerifiedLookerEmployee); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
