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

// Look look
//
// swagger:model Look
type Look struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Content Favorite Id
	// Read Only: true
	ContentFavoriteID int64 `json:"content_favorite_id,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// Time that the Look was created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Whether or not a look is 'soft' deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Time that the Look was deleted.
	// Read Only: true
	// Format: date-time
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	// Id of User that deleted the look.
	// Read Only: true
	DeleterID int64 `json:"deleter_id,omitempty"`

	// Description
	Description string `json:"description,omitempty"`

	// Embed Url
	// Read Only: true
	EmbedURL string `json:"embed_url,omitempty"`

	// Excel File Url
	// Read Only: true
	ExcelFileURL string `json:"excel_file_url,omitempty"`

	// Number of times favorited
	// Read Only: true
	FavoriteCount int64 `json:"favorite_count,omitempty"`

	// Folder of this Look
	// Read Only: true
	Folder *FolderBase `json:"folder,omitempty"`

	// Folder Id
	FolderID string `json:"folder_id,omitempty"`

	// Google Spreadsheet Formula
	// Read Only: true
	GoogleSpreadsheetFormula string `json:"google_spreadsheet_formula,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Image Embed Url
	// Read Only: true
	ImageEmbedURL string `json:"image_embed_url,omitempty"`

	// auto-run query when Look viewed
	IsRunOnLoad bool `json:"is_run_on_load,omitempty"`

	// Time that the Look was last accessed by any user
	// Read Only: true
	// Format: date-time
	LastAccessedAt strfmt.DateTime `json:"last_accessed_at,omitempty"`

	// Id of User that last updated the look.
	// Read Only: true
	LastUpdaterID int64 `json:"last_updater_id,omitempty"`

	// Time last viewed in the Looker web UI
	// Read Only: true
	// Format: date-time
	LastViewedAt strfmt.DateTime `json:"last_viewed_at,omitempty"`

	// Model
	// Read Only: true
	Model *LookModel `json:"model,omitempty"`

	// Is Public
	Public bool `json:"public,omitempty"`

	// Public Slug
	// Read Only: true
	PublicSlug string `json:"public_slug,omitempty"`

	// Public Url
	// Read Only: true
	PublicURL string `json:"public_url,omitempty"`

	// Query Id
	QueryID int64 `json:"query_id,omitempty"`

	// Short Url
	// Read Only: true
	ShortURL string `json:"short_url,omitempty"`

	// Space of this Look
	// Read Only: true
	Space *SpaceBase `json:"space,omitempty"`

	// Space Id
	SpaceID string `json:"space_id,omitempty"`

	// Look Title
	Title string `json:"title,omitempty"`

	// Time that the Look was updated.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// (DEPRECATED) User
	// Read Only: true
	User *UserIDOnly `json:"user,omitempty"`

	// User Id
	UserID int64 `json:"user_id,omitempty"`

	// Number of times viewed in the Looker web UI
	// Read Only: true
	ViewCount int64 `json:"view_count,omitempty"`
}

// Validate validates this look
func (m *Look) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAccessedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastViewedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Look) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateDeletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateFolder(formats strfmt.Registry) error {
	if swag.IsZero(m.Folder) { // not required
		return nil
	}

	if m.Folder != nil {
		if err := m.Folder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("folder")
			}
			return err
		}
	}

	return nil
}

func (m *Look) validateLastAccessedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastAccessedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_accessed_at", "body", "date-time", m.LastAccessedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateLastViewedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastViewedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_viewed_at", "body", "date-time", m.LastViewedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *Look) validateSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

func (m *Look) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this look based on the context it is used
func (m *Look) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentFavoriteID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentMetadataID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmbedURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcelFileURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFavoriteCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFolder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogleSpreadsheetFormula(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageEmbedURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastAccessedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdaterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastViewedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicSlug(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShortURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Look) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Look) contextValidateContentFavoriteID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content_favorite_id", "body", int64(m.ContentFavoriteID)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateContentMetadataID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content_metadata_id", "body", int64(m.ContentMetadataID)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deleted_at", "body", strfmt.DateTime(m.DeletedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateDeleterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deleter_id", "body", int64(m.DeleterID)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateEmbedURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "embed_url", "body", string(m.EmbedURL)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateExcelFileURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "excel_file_url", "body", string(m.ExcelFileURL)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateFavoriteCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "favorite_count", "body", int64(m.FavoriteCount)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateFolder(ctx context.Context, formats strfmt.Registry) error {

	if m.Folder != nil {
		if err := m.Folder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("folder")
			}
			return err
		}
	}

	return nil
}

func (m *Look) contextValidateGoogleSpreadsheetFormula(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "google_spreadsheet_formula", "body", string(m.GoogleSpreadsheetFormula)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateImageEmbedURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "image_embed_url", "body", string(m.ImageEmbedURL)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateLastAccessedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_accessed_at", "body", strfmt.DateTime(m.LastAccessedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateLastUpdaterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updater_id", "body", int64(m.LastUpdaterID)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateLastViewedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_viewed_at", "body", strfmt.DateTime(m.LastViewedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if m.Model != nil {
		if err := m.Model.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *Look) contextValidatePublicSlug(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "public_slug", "body", string(m.PublicSlug)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidatePublicURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "public_url", "body", string(m.PublicURL)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateShortURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "short_url", "body", string(m.ShortURL)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateSpace(ctx context.Context, formats strfmt.Registry) error {

	if m.Space != nil {
		if err := m.Space.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

func (m *Look) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Look) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *Look) contextValidateViewCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "view_count", "body", int64(m.ViewCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Look) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Look) UnmarshalBinary(b []byte) error {
	var res Look
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
