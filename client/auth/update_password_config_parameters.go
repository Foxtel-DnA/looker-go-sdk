// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewUpdatePasswordConfigParams creates a new UpdatePasswordConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePasswordConfigParams() *UpdatePasswordConfigParams {
	return &UpdatePasswordConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePasswordConfigParamsWithTimeout creates a new UpdatePasswordConfigParams object
// with the ability to set a timeout on a request.
func NewUpdatePasswordConfigParamsWithTimeout(timeout time.Duration) *UpdatePasswordConfigParams {
	return &UpdatePasswordConfigParams{
		timeout: timeout,
	}
}

// NewUpdatePasswordConfigParamsWithContext creates a new UpdatePasswordConfigParams object
// with the ability to set a context for a request.
func NewUpdatePasswordConfigParamsWithContext(ctx context.Context) *UpdatePasswordConfigParams {
	return &UpdatePasswordConfigParams{
		Context: ctx,
	}
}

// NewUpdatePasswordConfigParamsWithHTTPClient creates a new UpdatePasswordConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePasswordConfigParamsWithHTTPClient(client *http.Client) *UpdatePasswordConfigParams {
	return &UpdatePasswordConfigParams{
		HTTPClient: client,
	}
}

/* UpdatePasswordConfigParams contains all the parameters to send to the API endpoint
   for the update password config operation.

   Typically these are written to a http.Request.
*/
type UpdatePasswordConfigParams struct {

	/* Body.

	   Password Config
	*/
	Body *models.PasswordConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update password config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePasswordConfigParams) WithDefaults() *UpdatePasswordConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update password config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePasswordConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update password config params
func (o *UpdatePasswordConfigParams) WithTimeout(timeout time.Duration) *UpdatePasswordConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update password config params
func (o *UpdatePasswordConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update password config params
func (o *UpdatePasswordConfigParams) WithContext(ctx context.Context) *UpdatePasswordConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update password config params
func (o *UpdatePasswordConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update password config params
func (o *UpdatePasswordConfigParams) WithHTTPClient(client *http.Client) *UpdatePasswordConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update password config params
func (o *UpdatePasswordConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update password config params
func (o *UpdatePasswordConfigParams) WithBody(body *models.PasswordConfig) *UpdatePasswordConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update password config params
func (o *UpdatePasswordConfigParams) SetBody(body *models.PasswordConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePasswordConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
