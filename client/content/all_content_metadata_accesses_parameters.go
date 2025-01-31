// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewAllContentMetadataAccessesParams creates a new AllContentMetadataAccessesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllContentMetadataAccessesParams() *AllContentMetadataAccessesParams {
	return &AllContentMetadataAccessesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllContentMetadataAccessesParamsWithTimeout creates a new AllContentMetadataAccessesParams object
// with the ability to set a timeout on a request.
func NewAllContentMetadataAccessesParamsWithTimeout(timeout time.Duration) *AllContentMetadataAccessesParams {
	return &AllContentMetadataAccessesParams{
		timeout: timeout,
	}
}

// NewAllContentMetadataAccessesParamsWithContext creates a new AllContentMetadataAccessesParams object
// with the ability to set a context for a request.
func NewAllContentMetadataAccessesParamsWithContext(ctx context.Context) *AllContentMetadataAccessesParams {
	return &AllContentMetadataAccessesParams{
		Context: ctx,
	}
}

// NewAllContentMetadataAccessesParamsWithHTTPClient creates a new AllContentMetadataAccessesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllContentMetadataAccessesParamsWithHTTPClient(client *http.Client) *AllContentMetadataAccessesParams {
	return &AllContentMetadataAccessesParams{
		HTTPClient: client,
	}
}

/* AllContentMetadataAccessesParams contains all the parameters to send to the API endpoint
   for the all content metadata accesses operation.

   Typically these are written to a http.Request.
*/
type AllContentMetadataAccessesParams struct {

	/* ContentMetadataID.

	   Id of content metadata

	   Format: int64
	*/
	ContentMetadataID int64

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all content metadata accesses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllContentMetadataAccessesParams) WithDefaults() *AllContentMetadataAccessesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all content metadata accesses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllContentMetadataAccessesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) WithTimeout(timeout time.Duration) *AllContentMetadataAccessesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) WithContext(ctx context.Context) *AllContentMetadataAccessesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) WithHTTPClient(client *http.Client) *AllContentMetadataAccessesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentMetadataID adds the contentMetadataID to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) WithContentMetadataID(contentMetadataID int64) *AllContentMetadataAccessesParams {
	o.SetContentMetadataID(contentMetadataID)
	return o
}

// SetContentMetadataID adds the contentMetadataId to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) SetContentMetadataID(contentMetadataID int64) {
	o.ContentMetadataID = contentMetadataID
}

// WithFields adds the fields to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) WithFields(fields *string) *AllContentMetadataAccessesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all content metadata accesses params
func (o *AllContentMetadataAccessesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllContentMetadataAccessesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param content_metadata_id
	qrContentMetadataID := o.ContentMetadataID
	qContentMetadataID := swag.FormatInt64(qrContentMetadataID)
	if qContentMetadataID != "" {

		if err := r.SetQueryParam("content_metadata_id", qContentMetadataID); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
