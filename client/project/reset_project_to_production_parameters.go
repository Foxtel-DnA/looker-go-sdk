// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewResetProjectToProductionParams creates a new ResetProjectToProductionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetProjectToProductionParams() *ResetProjectToProductionParams {
	return &ResetProjectToProductionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetProjectToProductionParamsWithTimeout creates a new ResetProjectToProductionParams object
// with the ability to set a timeout on a request.
func NewResetProjectToProductionParamsWithTimeout(timeout time.Duration) *ResetProjectToProductionParams {
	return &ResetProjectToProductionParams{
		timeout: timeout,
	}
}

// NewResetProjectToProductionParamsWithContext creates a new ResetProjectToProductionParams object
// with the ability to set a context for a request.
func NewResetProjectToProductionParamsWithContext(ctx context.Context) *ResetProjectToProductionParams {
	return &ResetProjectToProductionParams{
		Context: ctx,
	}
}

// NewResetProjectToProductionParamsWithHTTPClient creates a new ResetProjectToProductionParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetProjectToProductionParamsWithHTTPClient(client *http.Client) *ResetProjectToProductionParams {
	return &ResetProjectToProductionParams{
		HTTPClient: client,
	}
}

/* ResetProjectToProductionParams contains all the parameters to send to the API endpoint
   for the reset project to production operation.

   Typically these are written to a http.Request.
*/
type ResetProjectToProductionParams struct {

	/* ProjectID.

	   Id of project
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset project to production params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetProjectToProductionParams) WithDefaults() *ResetProjectToProductionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset project to production params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetProjectToProductionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset project to production params
func (o *ResetProjectToProductionParams) WithTimeout(timeout time.Duration) *ResetProjectToProductionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset project to production params
func (o *ResetProjectToProductionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset project to production params
func (o *ResetProjectToProductionParams) WithContext(ctx context.Context) *ResetProjectToProductionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset project to production params
func (o *ResetProjectToProductionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset project to production params
func (o *ResetProjectToProductionParams) WithHTTPClient(client *http.Client) *ResetProjectToProductionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset project to production params
func (o *ResetProjectToProductionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the reset project to production params
func (o *ResetProjectToProductionParams) WithProjectID(projectID string) *ResetProjectToProductionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the reset project to production params
func (o *ResetProjectToProductionParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetProjectToProductionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
