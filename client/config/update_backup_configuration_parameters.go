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
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewUpdateBackupConfigurationParams creates a new UpdateBackupConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBackupConfigurationParams() *UpdateBackupConfigurationParams {
	return &UpdateBackupConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBackupConfigurationParamsWithTimeout creates a new UpdateBackupConfigurationParams object
// with the ability to set a timeout on a request.
func NewUpdateBackupConfigurationParamsWithTimeout(timeout time.Duration) *UpdateBackupConfigurationParams {
	return &UpdateBackupConfigurationParams{
		timeout: timeout,
	}
}

// NewUpdateBackupConfigurationParamsWithContext creates a new UpdateBackupConfigurationParams object
// with the ability to set a context for a request.
func NewUpdateBackupConfigurationParamsWithContext(ctx context.Context) *UpdateBackupConfigurationParams {
	return &UpdateBackupConfigurationParams{
		Context: ctx,
	}
}

// NewUpdateBackupConfigurationParamsWithHTTPClient creates a new UpdateBackupConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBackupConfigurationParamsWithHTTPClient(client *http.Client) *UpdateBackupConfigurationParams {
	return &UpdateBackupConfigurationParams{
		HTTPClient: client,
	}
}

/* UpdateBackupConfigurationParams contains all the parameters to send to the API endpoint
   for the update backup configuration operation.

   Typically these are written to a http.Request.
*/
type UpdateBackupConfigurationParams struct {

	/* Body.

	   Options for Backup Configuration
	*/
	Body *models.BackupConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update backup configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBackupConfigurationParams) WithDefaults() *UpdateBackupConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update backup configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBackupConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update backup configuration params
func (o *UpdateBackupConfigurationParams) WithTimeout(timeout time.Duration) *UpdateBackupConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update backup configuration params
func (o *UpdateBackupConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update backup configuration params
func (o *UpdateBackupConfigurationParams) WithContext(ctx context.Context) *UpdateBackupConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update backup configuration params
func (o *UpdateBackupConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update backup configuration params
func (o *UpdateBackupConfigurationParams) WithHTTPClient(client *http.Client) *UpdateBackupConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update backup configuration params
func (o *UpdateBackupConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update backup configuration params
func (o *UpdateBackupConfigurationParams) WithBody(body *models.BackupConfiguration) *UpdateBackupConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update backup configuration params
func (o *UpdateBackupConfigurationParams) SetBody(body *models.BackupConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBackupConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
