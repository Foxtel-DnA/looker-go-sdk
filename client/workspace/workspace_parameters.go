// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewWorkspaceParams creates a new WorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkspaceParams() *WorkspaceParams {
	return &WorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkspaceParamsWithTimeout creates a new WorkspaceParams object
// with the ability to set a timeout on a request.
func NewWorkspaceParamsWithTimeout(timeout time.Duration) *WorkspaceParams {
	return &WorkspaceParams{
		timeout: timeout,
	}
}

// NewWorkspaceParamsWithContext creates a new WorkspaceParams object
// with the ability to set a context for a request.
func NewWorkspaceParamsWithContext(ctx context.Context) *WorkspaceParams {
	return &WorkspaceParams{
		Context: ctx,
	}
}

// NewWorkspaceParamsWithHTTPClient creates a new WorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkspaceParamsWithHTTPClient(client *http.Client) *WorkspaceParams {
	return &WorkspaceParams{
		HTTPClient: client,
	}
}

/* WorkspaceParams contains all the parameters to send to the API endpoint
   for the workspace operation.

   Typically these are written to a http.Request.
*/
type WorkspaceParams struct {

	/* WorkspaceID.

	   Id of the workspace
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceParams) WithDefaults() *WorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workspace params
func (o *WorkspaceParams) WithTimeout(timeout time.Duration) *WorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workspace params
func (o *WorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workspace params
func (o *WorkspaceParams) WithContext(ctx context.Context) *WorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workspace params
func (o *WorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workspace params
func (o *WorkspaceParams) WithHTTPClient(client *http.Client) *WorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workspace params
func (o *WorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the workspace params
func (o *WorkspaceParams) WithWorkspaceID(workspaceID string) *WorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the workspace params
func (o *WorkspaceParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
