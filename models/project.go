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

// Project project
//
// swagger:model Project
type Project struct {

	// Validation policy: If true, the project can be committed with warnings when `validation_required` is true. (`allow_warnings` does nothing if `validation_required` is false).
	AllowWarnings bool `json:"allow_warnings,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// (Write-Only) Optional secret token with which to authenticate requests to the webhook deploy endpoint. If not set, endpoint is unauthenticated.
	DeploySecret string `json:"deploy_secret,omitempty"`

	// Port that HTTP(S) application server is running on (for PRs, file browsing, etc.)
	GitApplicationServerHTTPPort int64 `json:"git_application_server_http_port,omitempty"`

	// Scheme that is running on application server (for PRs, file browsing, etc.) Valid values are: "http", "https".
	GitApplicationServerHTTPScheme string `json:"git_application_server_http_scheme,omitempty"`

	// (Write-Only) Git password for HTTPS authentication. (For production only, if using user attributes.)
	GitPassword string `json:"git_password,omitempty"`

	// User attribute name for password in per-user HTTPS authentication.
	GitPasswordUserAttribute string `json:"git_password_user_attribute,omitempty"`

	// Git production branch name. Defaults to master.
	GitProductionBranchName string `json:"git_production_branch_name,omitempty"`

	// If true, advanced git release management is enabled for this project
	GitReleaseMgmtEnabled bool `json:"git_release_mgmt_enabled,omitempty"`

	// Git remote repository url
	GitRemoteURL string `json:"git_remote_url,omitempty"`

	// Name of the git service provider
	GitServiceName string `json:"git_service_name,omitempty"`

	// Git username for HTTPS authentication. (For production only, if using user attributes.)
	GitUsername string `json:"git_username,omitempty"`

	// User attribute name for username in per-user HTTPS authentication.
	GitUsernameUserAttribute string `json:"git_username_user_attribute,omitempty"`

	// Project Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// If true the project is an example project and cannot be modified
	// Read Only: true
	IsExample *bool `json:"is_example,omitempty"`

	// Project display name
	Name string `json:"name,omitempty"`

	// The git pull request policy for this project. Valid values are: "off", "links", "recommended", "required".
	PullRequestMode string `json:"pull_request_mode,omitempty"`

	// (Write-Only) When true, unsets the deploy secret to allow unauthenticated access to the webhook deploy endpoint.
	UnsetDeploySecret bool `json:"unset_deploy_secret,omitempty"`

	// If true the project is configured with a git repository
	// Read Only: true
	UsesGit *bool `json:"uses_git,omitempty"`

	// Validation policy: If true, the project must pass validation checks before project changes can be committed to the git repository
	ValidationRequired bool `json:"validation_required,omitempty"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this project based on the context it is used
func (m *Project) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsExample(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsesGit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Project) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Project) contextValidateIsExample(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_example", "body", m.IsExample); err != nil {
		return err
	}

	return nil
}

func (m *Project) contextValidateUsesGit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uses_git", "body", m.UsesGit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
