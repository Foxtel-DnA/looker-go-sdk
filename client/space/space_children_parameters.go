// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSpaceChildrenParams creates a new SpaceChildrenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSpaceChildrenParams() *SpaceChildrenParams {
	return &SpaceChildrenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceChildrenParamsWithTimeout creates a new SpaceChildrenParams object
// with the ability to set a timeout on a request.
func NewSpaceChildrenParamsWithTimeout(timeout time.Duration) *SpaceChildrenParams {
	return &SpaceChildrenParams{
		timeout: timeout,
	}
}

// NewSpaceChildrenParamsWithContext creates a new SpaceChildrenParams object
// with the ability to set a context for a request.
func NewSpaceChildrenParamsWithContext(ctx context.Context) *SpaceChildrenParams {
	return &SpaceChildrenParams{
		Context: ctx,
	}
}

// NewSpaceChildrenParamsWithHTTPClient creates a new SpaceChildrenParams object
// with the ability to set a custom HTTPClient for a request.
func NewSpaceChildrenParamsWithHTTPClient(client *http.Client) *SpaceChildrenParams {
	return &SpaceChildrenParams{
		HTTPClient: client,
	}
}

/* SpaceChildrenParams contains all the parameters to send to the API endpoint
   for the space children operation.

   Typically these are written to a http.Request.
*/
type SpaceChildrenParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* Page.

	   Requested page.

	   Format: int64
	*/
	Page *int64

	/* PerPage.

	   Results per page.

	   Format: int64
	*/
	PerPage *int64

	/* Sorts.

	   Fields to sort by.
	*/
	Sorts *string

	/* SpaceID.

	   Id of space
	*/
	SpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the space children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpaceChildrenParams) WithDefaults() *SpaceChildrenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the space children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpaceChildrenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the space children params
func (o *SpaceChildrenParams) WithTimeout(timeout time.Duration) *SpaceChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the space children params
func (o *SpaceChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the space children params
func (o *SpaceChildrenParams) WithContext(ctx context.Context) *SpaceChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the space children params
func (o *SpaceChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the space children params
func (o *SpaceChildrenParams) WithHTTPClient(client *http.Client) *SpaceChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the space children params
func (o *SpaceChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the space children params
func (o *SpaceChildrenParams) WithFields(fields *string) *SpaceChildrenParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the space children params
func (o *SpaceChildrenParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPage adds the page to the space children params
func (o *SpaceChildrenParams) WithPage(page *int64) *SpaceChildrenParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the space children params
func (o *SpaceChildrenParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the space children params
func (o *SpaceChildrenParams) WithPerPage(perPage *int64) *SpaceChildrenParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the space children params
func (o *SpaceChildrenParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSorts adds the sorts to the space children params
func (o *SpaceChildrenParams) WithSorts(sorts *string) *SpaceChildrenParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the space children params
func (o *SpaceChildrenParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithSpaceID adds the spaceID to the space children params
func (o *SpaceChildrenParams) WithSpaceID(spaceID string) *SpaceChildrenParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the space children params
func (o *SpaceChildrenParams) SetSpaceID(spaceID string) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
