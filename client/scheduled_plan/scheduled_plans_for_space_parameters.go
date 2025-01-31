// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

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
	"github.com/go-openapi/swag"
)

// NewScheduledPlansForSpaceParams creates a new ScheduledPlansForSpaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduledPlansForSpaceParams() *ScheduledPlansForSpaceParams {
	return &ScheduledPlansForSpaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduledPlansForSpaceParamsWithTimeout creates a new ScheduledPlansForSpaceParams object
// with the ability to set a timeout on a request.
func NewScheduledPlansForSpaceParamsWithTimeout(timeout time.Duration) *ScheduledPlansForSpaceParams {
	return &ScheduledPlansForSpaceParams{
		timeout: timeout,
	}
}

// NewScheduledPlansForSpaceParamsWithContext creates a new ScheduledPlansForSpaceParams object
// with the ability to set a context for a request.
func NewScheduledPlansForSpaceParamsWithContext(ctx context.Context) *ScheduledPlansForSpaceParams {
	return &ScheduledPlansForSpaceParams{
		Context: ctx,
	}
}

// NewScheduledPlansForSpaceParamsWithHTTPClient creates a new ScheduledPlansForSpaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduledPlansForSpaceParamsWithHTTPClient(client *http.Client) *ScheduledPlansForSpaceParams {
	return &ScheduledPlansForSpaceParams{
		HTTPClient: client,
	}
}

/* ScheduledPlansForSpaceParams contains all the parameters to send to the API endpoint
   for the scheduled plans for space operation.

   Typically these are written to a http.Request.
*/
type ScheduledPlansForSpaceParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* SpaceID.

	   Space Id

	   Format: int64
	*/
	SpaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduled plans for space params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduledPlansForSpaceParams) WithDefaults() *ScheduledPlansForSpaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduled plans for space params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduledPlansForSpaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) WithTimeout(timeout time.Duration) *ScheduledPlansForSpaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) WithContext(ctx context.Context) *ScheduledPlansForSpaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) WithHTTPClient(client *http.Client) *ScheduledPlansForSpaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) WithFields(fields *string) *ScheduledPlansForSpaceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSpaceID adds the spaceID to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) WithSpaceID(spaceID int64) *ScheduledPlansForSpaceParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the scheduled plans for space params
func (o *ScheduledPlansForSpaceParams) SetSpaceID(spaceID int64) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduledPlansForSpaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("space_id", swag.FormatInt64(o.SpaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
