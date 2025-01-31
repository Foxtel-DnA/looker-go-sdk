// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewDashboardElementParams creates a new DashboardElementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDashboardElementParams() *DashboardElementParams {
	return &DashboardElementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardElementParamsWithTimeout creates a new DashboardElementParams object
// with the ability to set a timeout on a request.
func NewDashboardElementParamsWithTimeout(timeout time.Duration) *DashboardElementParams {
	return &DashboardElementParams{
		timeout: timeout,
	}
}

// NewDashboardElementParamsWithContext creates a new DashboardElementParams object
// with the ability to set a context for a request.
func NewDashboardElementParamsWithContext(ctx context.Context) *DashboardElementParams {
	return &DashboardElementParams{
		Context: ctx,
	}
}

// NewDashboardElementParamsWithHTTPClient creates a new DashboardElementParams object
// with the ability to set a custom HTTPClient for a request.
func NewDashboardElementParamsWithHTTPClient(client *http.Client) *DashboardElementParams {
	return &DashboardElementParams{
		HTTPClient: client,
	}
}

/* DashboardElementParams contains all the parameters to send to the API endpoint
   for the dashboard element operation.

   Typically these are written to a http.Request.
*/
type DashboardElementParams struct {

	/* DashboardElementID.

	   Id of dashboard element
	*/
	DashboardElementID string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dashboard element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DashboardElementParams) WithDefaults() *DashboardElementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dashboard element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DashboardElementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dashboard element params
func (o *DashboardElementParams) WithTimeout(timeout time.Duration) *DashboardElementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard element params
func (o *DashboardElementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard element params
func (o *DashboardElementParams) WithContext(ctx context.Context) *DashboardElementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard element params
func (o *DashboardElementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard element params
func (o *DashboardElementParams) WithHTTPClient(client *http.Client) *DashboardElementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard element params
func (o *DashboardElementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardElementID adds the dashboardElementID to the dashboard element params
func (o *DashboardElementParams) WithDashboardElementID(dashboardElementID string) *DashboardElementParams {
	o.SetDashboardElementID(dashboardElementID)
	return o
}

// SetDashboardElementID adds the dashboardElementId to the dashboard element params
func (o *DashboardElementParams) SetDashboardElementID(dashboardElementID string) {
	o.DashboardElementID = dashboardElementID
}

// WithFields adds the fields to the dashboard element params
func (o *DashboardElementParams) WithFields(fields *string) *DashboardElementParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the dashboard element params
func (o *DashboardElementParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardElementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard_element_id
	if err := r.SetPathParam("dashboard_element_id", o.DashboardElementID); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
