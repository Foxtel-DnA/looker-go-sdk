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
)

// NewDeleteLookmlModelParams creates a new DeleteLookmlModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLookmlModelParams() *DeleteLookmlModelParams {
	return &DeleteLookmlModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLookmlModelParamsWithTimeout creates a new DeleteLookmlModelParams object
// with the ability to set a timeout on a request.
func NewDeleteLookmlModelParamsWithTimeout(timeout time.Duration) *DeleteLookmlModelParams {
	return &DeleteLookmlModelParams{
		timeout: timeout,
	}
}

// NewDeleteLookmlModelParamsWithContext creates a new DeleteLookmlModelParams object
// with the ability to set a context for a request.
func NewDeleteLookmlModelParamsWithContext(ctx context.Context) *DeleteLookmlModelParams {
	return &DeleteLookmlModelParams{
		Context: ctx,
	}
}

// NewDeleteLookmlModelParamsWithHTTPClient creates a new DeleteLookmlModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLookmlModelParamsWithHTTPClient(client *http.Client) *DeleteLookmlModelParams {
	return &DeleteLookmlModelParams{
		HTTPClient: client,
	}
}

/* DeleteLookmlModelParams contains all the parameters to send to the API endpoint
   for the delete lookml model operation.

   Typically these are written to a http.Request.
*/
type DeleteLookmlModelParams struct {

	/* LookmlModelName.

	   Name of lookml model.
	*/
	LookmlModelName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete lookml model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLookmlModelParams) WithDefaults() *DeleteLookmlModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete lookml model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLookmlModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete lookml model params
func (o *DeleteLookmlModelParams) WithTimeout(timeout time.Duration) *DeleteLookmlModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lookml model params
func (o *DeleteLookmlModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lookml model params
func (o *DeleteLookmlModelParams) WithContext(ctx context.Context) *DeleteLookmlModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lookml model params
func (o *DeleteLookmlModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lookml model params
func (o *DeleteLookmlModelParams) WithHTTPClient(client *http.Client) *DeleteLookmlModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lookml model params
func (o *DeleteLookmlModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLookmlModelName adds the lookmlModelName to the delete lookml model params
func (o *DeleteLookmlModelParams) WithLookmlModelName(lookmlModelName string) *DeleteLookmlModelParams {
	o.SetLookmlModelName(lookmlModelName)
	return o
}

// SetLookmlModelName adds the lookmlModelName to the delete lookml model params
func (o *DeleteLookmlModelParams) SetLookmlModelName(lookmlModelName string) {
	o.LookmlModelName = lookmlModelName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLookmlModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lookml_model_name
	if err := r.SetPathParam("lookml_model_name", o.LookmlModelName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
