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

	"your-damain.com/swagger/looker-api-golang/models"
)

// NewUpdateCloudStorageConfigurationParams creates a new UpdateCloudStorageConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCloudStorageConfigurationParams() *UpdateCloudStorageConfigurationParams {
	return &UpdateCloudStorageConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCloudStorageConfigurationParamsWithTimeout creates a new UpdateCloudStorageConfigurationParams object
// with the ability to set a timeout on a request.
func NewUpdateCloudStorageConfigurationParamsWithTimeout(timeout time.Duration) *UpdateCloudStorageConfigurationParams {
	return &UpdateCloudStorageConfigurationParams{
		timeout: timeout,
	}
}

// NewUpdateCloudStorageConfigurationParamsWithContext creates a new UpdateCloudStorageConfigurationParams object
// with the ability to set a context for a request.
func NewUpdateCloudStorageConfigurationParamsWithContext(ctx context.Context) *UpdateCloudStorageConfigurationParams {
	return &UpdateCloudStorageConfigurationParams{
		Context: ctx,
	}
}

// NewUpdateCloudStorageConfigurationParamsWithHTTPClient creates a new UpdateCloudStorageConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCloudStorageConfigurationParamsWithHTTPClient(client *http.Client) *UpdateCloudStorageConfigurationParams {
	return &UpdateCloudStorageConfigurationParams{
		HTTPClient: client,
	}
}

/* UpdateCloudStorageConfigurationParams contains all the parameters to send to the API endpoint
   for the update cloud storage configuration operation.

   Typically these are written to a http.Request.
*/
type UpdateCloudStorageConfigurationParams struct {

	/* Body.

	   Options for Cloud Storage Configuration
	*/
	Body *models.BackupConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cloud storage configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCloudStorageConfigurationParams) WithDefaults() *UpdateCloudStorageConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cloud storage configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCloudStorageConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) WithTimeout(timeout time.Duration) *UpdateCloudStorageConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) WithContext(ctx context.Context) *UpdateCloudStorageConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) WithHTTPClient(client *http.Client) *UpdateCloudStorageConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) WithBody(body *models.BackupConfiguration) *UpdateCloudStorageConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update cloud storage configuration params
func (o *UpdateCloudStorageConfigurationParams) SetBody(body *models.BackupConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCloudStorageConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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