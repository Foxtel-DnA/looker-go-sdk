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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewCreateHomepageParams creates a new CreateHomepageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHomepageParams() *CreateHomepageParams {
	return &CreateHomepageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHomepageParamsWithTimeout creates a new CreateHomepageParams object
// with the ability to set a timeout on a request.
func NewCreateHomepageParamsWithTimeout(timeout time.Duration) *CreateHomepageParams {
	return &CreateHomepageParams{
		timeout: timeout,
	}
}

// NewCreateHomepageParamsWithContext creates a new CreateHomepageParams object
// with the ability to set a context for a request.
func NewCreateHomepageParamsWithContext(ctx context.Context) *CreateHomepageParams {
	return &CreateHomepageParams{
		Context: ctx,
	}
}

// NewCreateHomepageParamsWithHTTPClient creates a new CreateHomepageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHomepageParamsWithHTTPClient(client *http.Client) *CreateHomepageParams {
	return &CreateHomepageParams{
		HTTPClient: client,
	}
}

/* CreateHomepageParams contains all the parameters to send to the API endpoint
   for the create homepage operation.

   Typically these are written to a http.Request.
*/
type CreateHomepageParams struct {

	/* Body.

	   Homepage
	*/
	Body *models.Homepage

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHomepageParams) WithDefaults() *CreateHomepageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create homepage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHomepageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create homepage params
func (o *CreateHomepageParams) WithTimeout(timeout time.Duration) *CreateHomepageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create homepage params
func (o *CreateHomepageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create homepage params
func (o *CreateHomepageParams) WithContext(ctx context.Context) *CreateHomepageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create homepage params
func (o *CreateHomepageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create homepage params
func (o *CreateHomepageParams) WithHTTPClient(client *http.Client) *CreateHomepageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create homepage params
func (o *CreateHomepageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create homepage params
func (o *CreateHomepageParams) WithBody(body *models.Homepage) *CreateHomepageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create homepage params
func (o *CreateHomepageParams) SetBody(body *models.Homepage) {
	o.Body = body
}

// WithFields adds the fields to the create homepage params
func (o *CreateHomepageParams) WithFields(fields *string) *CreateHomepageParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create homepage params
func (o *CreateHomepageParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHomepageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
