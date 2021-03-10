// Code generated by go-swagger; DO NOT EDIT.

package space

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

// NewSpaceLooksParams creates a new SpaceLooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSpaceLooksParams() *SpaceLooksParams {
	return &SpaceLooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceLooksParamsWithTimeout creates a new SpaceLooksParams object
// with the ability to set a timeout on a request.
func NewSpaceLooksParamsWithTimeout(timeout time.Duration) *SpaceLooksParams {
	return &SpaceLooksParams{
		timeout: timeout,
	}
}

// NewSpaceLooksParamsWithContext creates a new SpaceLooksParams object
// with the ability to set a context for a request.
func NewSpaceLooksParamsWithContext(ctx context.Context) *SpaceLooksParams {
	return &SpaceLooksParams{
		Context: ctx,
	}
}

// NewSpaceLooksParamsWithHTTPClient creates a new SpaceLooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewSpaceLooksParamsWithHTTPClient(client *http.Client) *SpaceLooksParams {
	return &SpaceLooksParams{
		HTTPClient: client,
	}
}

/* SpaceLooksParams contains all the parameters to send to the API endpoint
   for the space looks operation.

   Typically these are written to a http.Request.
*/
type SpaceLooksParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* SpaceID.

	   Id of space
	*/
	SpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the space looks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpaceLooksParams) WithDefaults() *SpaceLooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the space looks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpaceLooksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the space looks params
func (o *SpaceLooksParams) WithTimeout(timeout time.Duration) *SpaceLooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the space looks params
func (o *SpaceLooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the space looks params
func (o *SpaceLooksParams) WithContext(ctx context.Context) *SpaceLooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the space looks params
func (o *SpaceLooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the space looks params
func (o *SpaceLooksParams) WithHTTPClient(client *http.Client) *SpaceLooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the space looks params
func (o *SpaceLooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the space looks params
func (o *SpaceLooksParams) WithFields(fields *string) *SpaceLooksParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the space looks params
func (o *SpaceLooksParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSpaceID adds the spaceID to the space looks params
func (o *SpaceLooksParams) WithSpaceID(spaceID string) *SpaceLooksParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the space looks params
func (o *SpaceLooksParams) SetSpaceID(spaceID string) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceLooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
