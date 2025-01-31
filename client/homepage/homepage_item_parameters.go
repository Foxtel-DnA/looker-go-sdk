// Code generated by go-swagger; DO NOT EDIT.

package homepage

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

// NewHomepageItemParams creates a new HomepageItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHomepageItemParams() *HomepageItemParams {
	return &HomepageItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHomepageItemParamsWithTimeout creates a new HomepageItemParams object
// with the ability to set a timeout on a request.
func NewHomepageItemParamsWithTimeout(timeout time.Duration) *HomepageItemParams {
	return &HomepageItemParams{
		timeout: timeout,
	}
}

// NewHomepageItemParamsWithContext creates a new HomepageItemParams object
// with the ability to set a context for a request.
func NewHomepageItemParamsWithContext(ctx context.Context) *HomepageItemParams {
	return &HomepageItemParams{
		Context: ctx,
	}
}

// NewHomepageItemParamsWithHTTPClient creates a new HomepageItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewHomepageItemParamsWithHTTPClient(client *http.Client) *HomepageItemParams {
	return &HomepageItemParams{
		HTTPClient: client,
	}
}

/* HomepageItemParams contains all the parameters to send to the API endpoint
   for the homepage item operation.

   Typically these are written to a http.Request.
*/
type HomepageItemParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* HomepageItemID.

	   Id of homepage item

	   Format: int64
	*/
	HomepageItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the homepage item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HomepageItemParams) WithDefaults() *HomepageItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the homepage item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HomepageItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the homepage item params
func (o *HomepageItemParams) WithTimeout(timeout time.Duration) *HomepageItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the homepage item params
func (o *HomepageItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the homepage item params
func (o *HomepageItemParams) WithContext(ctx context.Context) *HomepageItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the homepage item params
func (o *HomepageItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the homepage item params
func (o *HomepageItemParams) WithHTTPClient(client *http.Client) *HomepageItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the homepage item params
func (o *HomepageItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the homepage item params
func (o *HomepageItemParams) WithFields(fields *string) *HomepageItemParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the homepage item params
func (o *HomepageItemParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHomepageItemID adds the homepageItemID to the homepage item params
func (o *HomepageItemParams) WithHomepageItemID(homepageItemID int64) *HomepageItemParams {
	o.SetHomepageItemID(homepageItemID)
	return o
}

// SetHomepageItemID adds the homepageItemId to the homepage item params
func (o *HomepageItemParams) SetHomepageItemID(homepageItemID int64) {
	o.HomepageItemID = homepageItemID
}

// WriteToRequest writes these params to a swagger request
func (o *HomepageItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homepage_item_id
	if err := r.SetPathParam("homepage_item_id", swag.FormatInt64(o.HomepageItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
