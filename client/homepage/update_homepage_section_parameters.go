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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewUpdateHomepageSectionParams creates a new UpdateHomepageSectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHomepageSectionParams() *UpdateHomepageSectionParams {
	return &UpdateHomepageSectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHomepageSectionParamsWithTimeout creates a new UpdateHomepageSectionParams object
// with the ability to set a timeout on a request.
func NewUpdateHomepageSectionParamsWithTimeout(timeout time.Duration) *UpdateHomepageSectionParams {
	return &UpdateHomepageSectionParams{
		timeout: timeout,
	}
}

// NewUpdateHomepageSectionParamsWithContext creates a new UpdateHomepageSectionParams object
// with the ability to set a context for a request.
func NewUpdateHomepageSectionParamsWithContext(ctx context.Context) *UpdateHomepageSectionParams {
	return &UpdateHomepageSectionParams{
		Context: ctx,
	}
}

// NewUpdateHomepageSectionParamsWithHTTPClient creates a new UpdateHomepageSectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHomepageSectionParamsWithHTTPClient(client *http.Client) *UpdateHomepageSectionParams {
	return &UpdateHomepageSectionParams{
		HTTPClient: client,
	}
}

/* UpdateHomepageSectionParams contains all the parameters to send to the API endpoint
   for the update homepage section operation.

   Typically these are written to a http.Request.
*/
type UpdateHomepageSectionParams struct {

	/* Body.

	   Homepage section
	*/
	Body *models.HomepageSection

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* HomepageSectionID.

	   Id of homepage section

	   Format: int64
	*/
	HomepageSectionID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update homepage section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageSectionParams) WithDefaults() *UpdateHomepageSectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update homepage section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageSectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update homepage section params
func (o *UpdateHomepageSectionParams) WithTimeout(timeout time.Duration) *UpdateHomepageSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update homepage section params
func (o *UpdateHomepageSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update homepage section params
func (o *UpdateHomepageSectionParams) WithContext(ctx context.Context) *UpdateHomepageSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update homepage section params
func (o *UpdateHomepageSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update homepage section params
func (o *UpdateHomepageSectionParams) WithHTTPClient(client *http.Client) *UpdateHomepageSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update homepage section params
func (o *UpdateHomepageSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update homepage section params
func (o *UpdateHomepageSectionParams) WithBody(body *models.HomepageSection) *UpdateHomepageSectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update homepage section params
func (o *UpdateHomepageSectionParams) SetBody(body *models.HomepageSection) {
	o.Body = body
}

// WithFields adds the fields to the update homepage section params
func (o *UpdateHomepageSectionParams) WithFields(fields *string) *UpdateHomepageSectionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update homepage section params
func (o *UpdateHomepageSectionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHomepageSectionID adds the homepageSectionID to the update homepage section params
func (o *UpdateHomepageSectionParams) WithHomepageSectionID(homepageSectionID int64) *UpdateHomepageSectionParams {
	o.SetHomepageSectionID(homepageSectionID)
	return o
}

// SetHomepageSectionID adds the homepageSectionId to the update homepage section params
func (o *UpdateHomepageSectionParams) SetHomepageSectionID(homepageSectionID int64) {
	o.HomepageSectionID = homepageSectionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHomepageSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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

	// path param homepage_section_id
	if err := r.SetPathParam("homepage_section_id", swag.FormatInt64(o.HomepageSectionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
