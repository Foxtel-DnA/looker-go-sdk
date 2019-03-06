// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewUpdateWhitelabelConfigurationParams creates a new UpdateWhitelabelConfigurationParams object
// with the default values initialized.
func NewUpdateWhitelabelConfigurationParams() *UpdateWhitelabelConfigurationParams {
	var ()
	return &UpdateWhitelabelConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWhitelabelConfigurationParamsWithTimeout creates a new UpdateWhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWhitelabelConfigurationParamsWithTimeout(timeout time.Duration) *UpdateWhitelabelConfigurationParams {
	var ()
	return &UpdateWhitelabelConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateWhitelabelConfigurationParamsWithContext creates a new UpdateWhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWhitelabelConfigurationParamsWithContext(ctx context.Context) *UpdateWhitelabelConfigurationParams {
	var ()
	return &UpdateWhitelabelConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateWhitelabelConfigurationParamsWithHTTPClient creates a new UpdateWhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWhitelabelConfigurationParamsWithHTTPClient(client *http.Client) *UpdateWhitelabelConfigurationParams {
	var ()
	return &UpdateWhitelabelConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateWhitelabelConfigurationParams contains all the parameters to send to the API endpoint
for the update whitelabel configuration operation typically these are written to a http.Request
*/
type UpdateWhitelabelConfigurationParams struct {

	/*Body
	  Whitelabel configuration

	*/
	Body *models.WhitelabelConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) WithTimeout(timeout time.Duration) *UpdateWhitelabelConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) WithContext(ctx context.Context) *UpdateWhitelabelConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) WithHTTPClient(client *http.Client) *UpdateWhitelabelConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) WithBody(body *models.WhitelabelConfiguration) *UpdateWhitelabelConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update whitelabel configuration params
func (o *UpdateWhitelabelConfigurationParams) SetBody(body *models.WhitelabelConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWhitelabelConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
