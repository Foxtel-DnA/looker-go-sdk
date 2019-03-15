// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllGroupsParams creates a new AllGroupsParams object
// with the default values initialized.
func NewAllGroupsParams() *AllGroupsParams {
	var ()
	return &AllGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllGroupsParamsWithTimeout creates a new AllGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllGroupsParamsWithTimeout(timeout time.Duration) *AllGroupsParams {
	var ()
	return &AllGroupsParams{

		timeout: timeout,
	}
}

// NewAllGroupsParamsWithContext creates a new AllGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllGroupsParamsWithContext(ctx context.Context) *AllGroupsParams {
	var ()
	return &AllGroupsParams{

		Context: ctx,
	}
}

// NewAllGroupsParamsWithHTTPClient creates a new AllGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllGroupsParamsWithHTTPClient(client *http.Client) *AllGroupsParams {
	var ()
	return &AllGroupsParams{
		HTTPClient: client,
	}
}

/*AllGroupsParams contains all the parameters to send to the API endpoint
for the all groups operation typically these are written to a http.Request
*/
type AllGroupsParams struct {

	/*CanAddToContentMetadata
	  Select only groups that either can/cannot be given access to content.

	*/
	CanAddToContentMetadata *bool
	/*ContentMetadataID
	  Id of content metadata to which groups must have access.

	*/
	ContentMetadataID *int64
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Ids
	  Optional of ids to get specific groups.

	*/
	Ids []int64
	/*Page
	  Requested page.

	*/
	Page *int64
	/*PerPage
	  Results per page.

	*/
	PerPage *int64
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all groups params
func (o *AllGroupsParams) WithTimeout(timeout time.Duration) *AllGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all groups params
func (o *AllGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all groups params
func (o *AllGroupsParams) WithContext(ctx context.Context) *AllGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all groups params
func (o *AllGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all groups params
func (o *AllGroupsParams) WithHTTPClient(client *http.Client) *AllGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all groups params
func (o *AllGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCanAddToContentMetadata adds the canAddToContentMetadata to the all groups params
func (o *AllGroupsParams) WithCanAddToContentMetadata(canAddToContentMetadata *bool) *AllGroupsParams {
	o.SetCanAddToContentMetadata(canAddToContentMetadata)
	return o
}

// SetCanAddToContentMetadata adds the canAddToContentMetadata to the all groups params
func (o *AllGroupsParams) SetCanAddToContentMetadata(canAddToContentMetadata *bool) {
	o.CanAddToContentMetadata = canAddToContentMetadata
}

// WithContentMetadataID adds the contentMetadataID to the all groups params
func (o *AllGroupsParams) WithContentMetadataID(contentMetadataID *int64) *AllGroupsParams {
	o.SetContentMetadataID(contentMetadataID)
	return o
}

// SetContentMetadataID adds the contentMetadataId to the all groups params
func (o *AllGroupsParams) SetContentMetadataID(contentMetadataID *int64) {
	o.ContentMetadataID = contentMetadataID
}

// WithFields adds the fields to the all groups params
func (o *AllGroupsParams) WithFields(fields *string) *AllGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all groups params
func (o *AllGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIds adds the ids to the all groups params
func (o *AllGroupsParams) WithIds(ids []int64) *AllGroupsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the all groups params
func (o *AllGroupsParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WithPage adds the page to the all groups params
func (o *AllGroupsParams) WithPage(page *int64) *AllGroupsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the all groups params
func (o *AllGroupsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the all groups params
func (o *AllGroupsParams) WithPerPage(perPage *int64) *AllGroupsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the all groups params
func (o *AllGroupsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSorts adds the sorts to the all groups params
func (o *AllGroupsParams) WithSorts(sorts *string) *AllGroupsParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the all groups params
func (o *AllGroupsParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WriteToRequest writes these params to a swagger request
func (o *AllGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CanAddToContentMetadata != nil {

		// query param can_add_to_content_metadata
		var qrCanAddToContentMetadata bool
		if o.CanAddToContentMetadata != nil {
			qrCanAddToContentMetadata = *o.CanAddToContentMetadata
		}
		qCanAddToContentMetadata := swag.FormatBool(qrCanAddToContentMetadata)
		if qCanAddToContentMetadata != "" {
			if err := r.SetQueryParam("can_add_to_content_metadata", qCanAddToContentMetadata); err != nil {
				return err
			}
		}

	}

	if o.ContentMetadataID != nil {

		// query param content_metadata_id
		var qrContentMetadataID int64
		if o.ContentMetadataID != nil {
			qrContentMetadataID = *o.ContentMetadataID
		}
		qContentMetadataID := swag.FormatInt64(qrContentMetadataID)
		if qContentMetadataID != "" {
			if err := r.SetQueryParam("content_metadata_id", qContentMetadataID); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	var valuesIds []string
	for _, v := range o.Ids {
		valuesIds = append(valuesIds, swag.FormatInt64(v))
	}

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string
		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {
			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}