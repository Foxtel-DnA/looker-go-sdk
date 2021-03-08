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

// NewUpdateHomepageParams creates a new UpdateHomepageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHomepageParams() *UpdateHomepageParams {
	return &UpdateHomepageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHomepageParamsWithTimeout creates a new UpdateHomepageParams object
// with the ability to set a timeout on a request.
func NewUpdateHomepageParamsWithTimeout(timeout time.Duration) *UpdateHomepageParams {
	return &UpdateHomepageParams{
		timeout: timeout,
	}
}

// NewUpdateHomepageParamsWithContext creates a new UpdateHomepageParams object
// with the ability to set a context for a request.
func NewUpdateHomepageParamsWithContext(ctx context.Context) *UpdateHomepageParams {
	return &UpdateHomepageParams{
		Context: ctx,
	}
}

// NewUpdateHomepageParamsWithHTTPClient creates a new UpdateHomepageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHomepageParamsWithHTTPClient(client *http.Client) *UpdateHomepageParams {
	return &UpdateHomepageParams{
		HTTPClient: client,
	}
}

/* UpdateHomepageParams contains all the parameters to send to the API endpoint
   for the update homepage operation.

   Typically these are written to a http.Request.
*/
type UpdateHomepageParams struct {

	/* Body.

	   Homepage
	*/
	Body *models.Homepage

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* HomepageID.

	   Id of homepage

	   Format: int64
	*/
	HomepageID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageParams) WithDefaults() *UpdateHomepageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update homepage params
func (o *UpdateHomepageParams) WithTimeout(timeout time.Duration) *UpdateHomepageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update homepage params
func (o *UpdateHomepageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update homepage params
func (o *UpdateHomepageParams) WithContext(ctx context.Context) *UpdateHomepageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update homepage params
func (o *UpdateHomepageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update homepage params
func (o *UpdateHomepageParams) WithHTTPClient(client *http.Client) *UpdateHomepageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update homepage params
func (o *UpdateHomepageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update homepage params
func (o *UpdateHomepageParams) WithBody(body *models.Homepage) *UpdateHomepageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update homepage params
func (o *UpdateHomepageParams) SetBody(body *models.Homepage) {
	o.Body = body
}

// WithFields adds the fields to the update homepage params
func (o *UpdateHomepageParams) WithFields(fields *string) *UpdateHomepageParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update homepage params
func (o *UpdateHomepageParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHomepageID adds the homepageID to the update homepage params
func (o *UpdateHomepageParams) WithHomepageID(homepageID int64) *UpdateHomepageParams {
	o.SetHomepageID(homepageID)
	return o
}

// SetHomepageID adds the homepageId to the update homepage params
func (o *UpdateHomepageParams) SetHomepageID(homepageID int64) {
	o.HomepageID = homepageID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHomepageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homepage_id
	if err := r.SetPathParam("homepage_id", swag.FormatInt64(o.HomepageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
