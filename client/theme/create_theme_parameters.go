// Code generated by go-swagger; DO NOT EDIT.

package theme

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

// NewCreateThemeParams creates a new CreateThemeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateThemeParams() *CreateThemeParams {
	return &CreateThemeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateThemeParamsWithTimeout creates a new CreateThemeParams object
// with the ability to set a timeout on a request.
func NewCreateThemeParamsWithTimeout(timeout time.Duration) *CreateThemeParams {
	return &CreateThemeParams{
		timeout: timeout,
	}
}

// NewCreateThemeParamsWithContext creates a new CreateThemeParams object
// with the ability to set a context for a request.
func NewCreateThemeParamsWithContext(ctx context.Context) *CreateThemeParams {
	return &CreateThemeParams{
		Context: ctx,
	}
}

// NewCreateThemeParamsWithHTTPClient creates a new CreateThemeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateThemeParamsWithHTTPClient(client *http.Client) *CreateThemeParams {
	return &CreateThemeParams{
		HTTPClient: client,
	}
}

/* CreateThemeParams contains all the parameters to send to the API endpoint
   for the create theme operation.

   Typically these are written to a http.Request.
*/
type CreateThemeParams struct {

	/* Body.

	   Theme
	*/
	Body *models.Theme

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create theme params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateThemeParams) WithDefaults() *CreateThemeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create theme params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateThemeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create theme params
func (o *CreateThemeParams) WithTimeout(timeout time.Duration) *CreateThemeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create theme params
func (o *CreateThemeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create theme params
func (o *CreateThemeParams) WithContext(ctx context.Context) *CreateThemeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create theme params
func (o *CreateThemeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create theme params
func (o *CreateThemeParams) WithHTTPClient(client *http.Client) *CreateThemeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create theme params
func (o *CreateThemeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create theme params
func (o *CreateThemeParams) WithBody(body *models.Theme) *CreateThemeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create theme params
func (o *CreateThemeParams) SetBody(body *models.Theme) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateThemeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
