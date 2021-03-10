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

	"github.com/billtrust/looker-go-sdk/models"
)

// NewUpdateInternalHelpResourcesParams creates a new UpdateInternalHelpResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateInternalHelpResourcesParams() *UpdateInternalHelpResourcesParams {
	return &UpdateInternalHelpResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInternalHelpResourcesParamsWithTimeout creates a new UpdateInternalHelpResourcesParams object
// with the ability to set a timeout on a request.
func NewUpdateInternalHelpResourcesParamsWithTimeout(timeout time.Duration) *UpdateInternalHelpResourcesParams {
	return &UpdateInternalHelpResourcesParams{
		timeout: timeout,
	}
}

// NewUpdateInternalHelpResourcesParamsWithContext creates a new UpdateInternalHelpResourcesParams object
// with the ability to set a context for a request.
func NewUpdateInternalHelpResourcesParamsWithContext(ctx context.Context) *UpdateInternalHelpResourcesParams {
	return &UpdateInternalHelpResourcesParams{
		Context: ctx,
	}
}

// NewUpdateInternalHelpResourcesParamsWithHTTPClient creates a new UpdateInternalHelpResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateInternalHelpResourcesParamsWithHTTPClient(client *http.Client) *UpdateInternalHelpResourcesParams {
	return &UpdateInternalHelpResourcesParams{
		HTTPClient: client,
	}
}

/* UpdateInternalHelpResourcesParams contains all the parameters to send to the API endpoint
   for the update internal help resources operation.

   Typically these are written to a http.Request.
*/
type UpdateInternalHelpResourcesParams struct {

	/* Body.

	   Custom Welcome Email
	*/
	Body *models.InternalHelpResources

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update internal help resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInternalHelpResourcesParams) WithDefaults() *UpdateInternalHelpResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update internal help resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInternalHelpResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) WithTimeout(timeout time.Duration) *UpdateInternalHelpResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) WithContext(ctx context.Context) *UpdateInternalHelpResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) WithHTTPClient(client *http.Client) *UpdateInternalHelpResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) WithBody(body *models.InternalHelpResources) *UpdateInternalHelpResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update internal help resources params
func (o *UpdateInternalHelpResourcesParams) SetBody(body *models.InternalHelpResources) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInternalHelpResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
