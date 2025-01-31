// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewCreateDashboardFilterParams creates a new CreateDashboardFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDashboardFilterParams() *CreateDashboardFilterParams {
	return &CreateDashboardFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDashboardFilterParamsWithTimeout creates a new CreateDashboardFilterParams object
// with the ability to set a timeout on a request.
func NewCreateDashboardFilterParamsWithTimeout(timeout time.Duration) *CreateDashboardFilterParams {
	return &CreateDashboardFilterParams{
		timeout: timeout,
	}
}

// NewCreateDashboardFilterParamsWithContext creates a new CreateDashboardFilterParams object
// with the ability to set a context for a request.
func NewCreateDashboardFilterParamsWithContext(ctx context.Context) *CreateDashboardFilterParams {
	return &CreateDashboardFilterParams{
		Context: ctx,
	}
}

// NewCreateDashboardFilterParamsWithHTTPClient creates a new CreateDashboardFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDashboardFilterParamsWithHTTPClient(client *http.Client) *CreateDashboardFilterParams {
	return &CreateDashboardFilterParams{
		HTTPClient: client,
	}
}

/* CreateDashboardFilterParams contains all the parameters to send to the API endpoint
   for the create dashboard filter operation.

   Typically these are written to a http.Request.
*/
type CreateDashboardFilterParams struct {

	/* Body.

	   Dashboard Filter
	*/
	Body *models.CreateDashboardFilter

	/* Fields.

	   Requested fields
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create dashboard filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDashboardFilterParams) WithDefaults() *CreateDashboardFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create dashboard filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDashboardFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create dashboard filter params
func (o *CreateDashboardFilterParams) WithTimeout(timeout time.Duration) *CreateDashboardFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create dashboard filter params
func (o *CreateDashboardFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create dashboard filter params
func (o *CreateDashboardFilterParams) WithContext(ctx context.Context) *CreateDashboardFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create dashboard filter params
func (o *CreateDashboardFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create dashboard filter params
func (o *CreateDashboardFilterParams) WithHTTPClient(client *http.Client) *CreateDashboardFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create dashboard filter params
func (o *CreateDashboardFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create dashboard filter params
func (o *CreateDashboardFilterParams) WithBody(body *models.CreateDashboardFilter) *CreateDashboardFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create dashboard filter params
func (o *CreateDashboardFilterParams) SetBody(body *models.CreateDashboardFilter) {
	o.Body = body
}

// WithFields adds the fields to the create dashboard filter params
func (o *CreateDashboardFilterParams) WithFields(fields *string) *CreateDashboardFilterParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create dashboard filter params
func (o *CreateDashboardFilterParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDashboardFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
