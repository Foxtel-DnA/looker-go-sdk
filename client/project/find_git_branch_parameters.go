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

// NewFindGitBranchParams creates a new FindGitBranchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindGitBranchParams() *FindGitBranchParams {
	return &FindGitBranchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindGitBranchParamsWithTimeout creates a new FindGitBranchParams object
// with the ability to set a timeout on a request.
func NewFindGitBranchParamsWithTimeout(timeout time.Duration) *FindGitBranchParams {
	return &FindGitBranchParams{
		timeout: timeout,
	}
}

// NewFindGitBranchParamsWithContext creates a new FindGitBranchParams object
// with the ability to set a context for a request.
func NewFindGitBranchParamsWithContext(ctx context.Context) *FindGitBranchParams {
	return &FindGitBranchParams{
		Context: ctx,
	}
}

// NewFindGitBranchParamsWithHTTPClient creates a new FindGitBranchParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindGitBranchParamsWithHTTPClient(client *http.Client) *FindGitBranchParams {
	return &FindGitBranchParams{
		HTTPClient: client,
	}
}

/* FindGitBranchParams contains all the parameters to send to the API endpoint
   for the find git branch operation.

   Typically these are written to a http.Request.
*/
type FindGitBranchParams struct {

	/* BranchName.

	   Branch Name
	*/
	BranchName string

	/* ProjectID.

	   Project Id
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find git branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindGitBranchParams) WithDefaults() *FindGitBranchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find git branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindGitBranchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find git branch params
func (o *FindGitBranchParams) WithTimeout(timeout time.Duration) *FindGitBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find git branch params
func (o *FindGitBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find git branch params
func (o *FindGitBranchParams) WithContext(ctx context.Context) *FindGitBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find git branch params
func (o *FindGitBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find git branch params
func (o *FindGitBranchParams) WithHTTPClient(client *http.Client) *FindGitBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find git branch params
func (o *FindGitBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranchName adds the branchName to the find git branch params
func (o *FindGitBranchParams) WithBranchName(branchName string) *FindGitBranchParams {
	o.SetBranchName(branchName)
	return o
}

// SetBranchName adds the branchName to the find git branch params
func (o *FindGitBranchParams) SetBranchName(branchName string) {
	o.BranchName = branchName
}

// WithProjectID adds the projectID to the find git branch params
func (o *FindGitBranchParams) WithProjectID(projectID string) *FindGitBranchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the find git branch params
func (o *FindGitBranchParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *FindGitBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branch_name
	if err := r.SetPathParam("branch_name", o.BranchName); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
