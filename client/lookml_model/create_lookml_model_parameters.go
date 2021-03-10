// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

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

// NewCreateLookmlModelParams creates a new CreateLookmlModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLookmlModelParams() *CreateLookmlModelParams {
	return &CreateLookmlModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLookmlModelParamsWithTimeout creates a new CreateLookmlModelParams object
// with the ability to set a timeout on a request.
func NewCreateLookmlModelParamsWithTimeout(timeout time.Duration) *CreateLookmlModelParams {
	return &CreateLookmlModelParams{
		timeout: timeout,
	}
}

// NewCreateLookmlModelParamsWithContext creates a new CreateLookmlModelParams object
// with the ability to set a context for a request.
func NewCreateLookmlModelParamsWithContext(ctx context.Context) *CreateLookmlModelParams {
	return &CreateLookmlModelParams{
		Context: ctx,
	}
}

// NewCreateLookmlModelParamsWithHTTPClient creates a new CreateLookmlModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLookmlModelParamsWithHTTPClient(client *http.Client) *CreateLookmlModelParams {
	return &CreateLookmlModelParams{
		HTTPClient: client,
	}
}

/* CreateLookmlModelParams contains all the parameters to send to the API endpoint
   for the create lookml model operation.

   Typically these are written to a http.Request.
*/
type CreateLookmlModelParams struct {

	/* Body.

	   LookML Model
	*/
	Body *models.LookmlModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create lookml model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLookmlModelParams) WithDefaults() *CreateLookmlModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create lookml model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLookmlModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create lookml model params
func (o *CreateLookmlModelParams) WithTimeout(timeout time.Duration) *CreateLookmlModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create lookml model params
func (o *CreateLookmlModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create lookml model params
func (o *CreateLookmlModelParams) WithContext(ctx context.Context) *CreateLookmlModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create lookml model params
func (o *CreateLookmlModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create lookml model params
func (o *CreateLookmlModelParams) WithHTTPClient(client *http.Client) *CreateLookmlModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create lookml model params
func (o *CreateLookmlModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create lookml model params
func (o *CreateLookmlModelParams) WithBody(body *models.LookmlModel) *CreateLookmlModelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create lookml model params
func (o *CreateLookmlModelParams) SetBody(body *models.LookmlModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLookmlModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
