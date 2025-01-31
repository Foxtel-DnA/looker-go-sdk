// Code generated by go-swagger; DO NOT EDIT.

package look

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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewUpdateLookParams creates a new UpdateLookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLookParams() *UpdateLookParams {
	return &UpdateLookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLookParamsWithTimeout creates a new UpdateLookParams object
// with the ability to set a timeout on a request.
func NewUpdateLookParamsWithTimeout(timeout time.Duration) *UpdateLookParams {
	return &UpdateLookParams{
		timeout: timeout,
	}
}

// NewUpdateLookParamsWithContext creates a new UpdateLookParams object
// with the ability to set a context for a request.
func NewUpdateLookParamsWithContext(ctx context.Context) *UpdateLookParams {
	return &UpdateLookParams{
		Context: ctx,
	}
}

// NewUpdateLookParamsWithHTTPClient creates a new UpdateLookParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLookParamsWithHTTPClient(client *http.Client) *UpdateLookParams {
	return &UpdateLookParams{
		HTTPClient: client,
	}
}

/* UpdateLookParams contains all the parameters to send to the API endpoint
   for the update look operation.

   Typically these are written to a http.Request.
*/
type UpdateLookParams struct {

	/* Body.

	   Look
	*/
	Body *models.LookWithQuery

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* LookID.

	   Id of look

	   Format: int64
	*/
	LookID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update look params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLookParams) WithDefaults() *UpdateLookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update look params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update look params
func (o *UpdateLookParams) WithTimeout(timeout time.Duration) *UpdateLookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update look params
func (o *UpdateLookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update look params
func (o *UpdateLookParams) WithContext(ctx context.Context) *UpdateLookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update look params
func (o *UpdateLookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update look params
func (o *UpdateLookParams) WithHTTPClient(client *http.Client) *UpdateLookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update look params
func (o *UpdateLookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update look params
func (o *UpdateLookParams) WithBody(body *models.LookWithQuery) *UpdateLookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update look params
func (o *UpdateLookParams) SetBody(body *models.LookWithQuery) {
	o.Body = body
}

// WithFields adds the fields to the update look params
func (o *UpdateLookParams) WithFields(fields *string) *UpdateLookParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update look params
func (o *UpdateLookParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLookID adds the lookID to the update look params
func (o *UpdateLookParams) WithLookID(lookID int64) *UpdateLookParams {
	o.SetLookID(lookID)
	return o
}

// SetLookID adds the lookId to the update look params
func (o *UpdateLookParams) SetLookID(lookID int64) {
	o.LookID = lookID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

	// path param look_id
	if err := r.SetPathParam("look_id", swag.FormatInt64(o.LookID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
