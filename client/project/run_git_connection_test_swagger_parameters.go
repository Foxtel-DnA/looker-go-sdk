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

// NewRunGitConnectionTestParams creates a new RunGitConnectionTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunGitConnectionTestParams() *RunGitConnectionTestParams {
	return &RunGitConnectionTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunGitConnectionTestParamsWithTimeout creates a new RunGitConnectionTestParams object
// with the ability to set a timeout on a request.
func NewRunGitConnectionTestParamsWithTimeout(timeout time.Duration) *RunGitConnectionTestParams {
	return &RunGitConnectionTestParams{
		timeout: timeout,
	}
}

// NewRunGitConnectionTestParamsWithContext creates a new RunGitConnectionTestParams object
// with the ability to set a context for a request.
func NewRunGitConnectionTestParamsWithContext(ctx context.Context) *RunGitConnectionTestParams {
	return &RunGitConnectionTestParams{
		Context: ctx,
	}
}

// NewRunGitConnectionTestParamsWithHTTPClient creates a new RunGitConnectionTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunGitConnectionTestParamsWithHTTPClient(client *http.Client) *RunGitConnectionTestParams {
	return &RunGitConnectionTestParams{
		HTTPClient: client,
	}
}

/* RunGitConnectionTestParams contains all the parameters to send to the API endpoint
   for the run git connection test operation.

   Typically these are written to a http.Request.
*/
type RunGitConnectionTestParams struct {

	/* ProjectID.

	   Project Id
	*/
	ProjectID string

	/* RemoteURL.

	   (Optional: leave blank for root project) The remote url for remote dependency to test.
	*/
	RemoteURL *string

	/* TestID.

	   Test Id
	*/
	TestID string

	/* UseProduction.

	   (Optional: leave blank for dev credentials) Whether to use git production credentials.
	*/
	UseProduction *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run git connection test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunGitConnectionTestParams) WithDefaults() *RunGitConnectionTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run git connection test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunGitConnectionTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run git connection test params
func (o *RunGitConnectionTestParams) WithTimeout(timeout time.Duration) *RunGitConnectionTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run git connection test params
func (o *RunGitConnectionTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run git connection test params
func (o *RunGitConnectionTestParams) WithContext(ctx context.Context) *RunGitConnectionTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run git connection test params
func (o *RunGitConnectionTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run git connection test params
func (o *RunGitConnectionTestParams) WithHTTPClient(client *http.Client) *RunGitConnectionTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run git connection test params
func (o *RunGitConnectionTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the run git connection test params
func (o *RunGitConnectionTestParams) WithProjectID(projectID string) *RunGitConnectionTestParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the run git connection test params
func (o *RunGitConnectionTestParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRemoteURL adds the remoteURL to the run git connection test params
func (o *RunGitConnectionTestParams) WithRemoteURL(remoteURL *string) *RunGitConnectionTestParams {
	o.SetRemoteURL(remoteURL)
	return o
}

// SetRemoteURL adds the remoteUrl to the run git connection test params
func (o *RunGitConnectionTestParams) SetRemoteURL(remoteURL *string) {
	o.RemoteURL = remoteURL
}

// WithTestID adds the testID to the run git connection test params
func (o *RunGitConnectionTestParams) WithTestID(testID string) *RunGitConnectionTestParams {
	o.SetTestID(testID)
	return o
}

// SetTestID adds the testId to the run git connection test params
func (o *RunGitConnectionTestParams) SetTestID(testID string) {
	o.TestID = testID
}

// WithUseProduction adds the useProduction to the run git connection test params
func (o *RunGitConnectionTestParams) WithUseProduction(useProduction *string) *RunGitConnectionTestParams {
	o.SetUseProduction(useProduction)
	return o
}

// SetUseProduction adds the useProduction to the run git connection test params
func (o *RunGitConnectionTestParams) SetUseProduction(useProduction *string) {
	o.UseProduction = useProduction
}

// WriteToRequest writes these params to a swagger request
func (o *RunGitConnectionTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.RemoteURL != nil {

		// query param remote_url
		var qrRemoteURL string

		if o.RemoteURL != nil {
			qrRemoteURL = *o.RemoteURL
		}
		qRemoteURL := qrRemoteURL
		if qRemoteURL != "" {

			if err := r.SetQueryParam("remote_url", qRemoteURL); err != nil {
				return err
			}
		}
	}

	// path param test_id
	if err := r.SetPathParam("test_id", o.TestID); err != nil {
		return err
	}

	if o.UseProduction != nil {

		// query param use_production
		var qrUseProduction string

		if o.UseProduction != nil {
			qrUseProduction = *o.UseProduction
		}
		qUseProduction := qrUseProduction
		if qUseProduction != "" {

			if err := r.SetQueryParam("use_production", qUseProduction); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
