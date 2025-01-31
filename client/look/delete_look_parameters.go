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
)

// NewDeleteLookParams creates a new DeleteLookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLookParams() *DeleteLookParams {
	return &DeleteLookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLookParamsWithTimeout creates a new DeleteLookParams object
// with the ability to set a timeout on a request.
func NewDeleteLookParamsWithTimeout(timeout time.Duration) *DeleteLookParams {
	return &DeleteLookParams{
		timeout: timeout,
	}
}

// NewDeleteLookParamsWithContext creates a new DeleteLookParams object
// with the ability to set a context for a request.
func NewDeleteLookParamsWithContext(ctx context.Context) *DeleteLookParams {
	return &DeleteLookParams{
		Context: ctx,
	}
}

// NewDeleteLookParamsWithHTTPClient creates a new DeleteLookParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLookParamsWithHTTPClient(client *http.Client) *DeleteLookParams {
	return &DeleteLookParams{
		HTTPClient: client,
	}
}

/* DeleteLookParams contains all the parameters to send to the API endpoint
   for the delete look operation.

   Typically these are written to a http.Request.
*/
type DeleteLookParams struct {

	/* LookID.

	   Id of look

	   Format: int64
	*/
	LookID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete look params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLookParams) WithDefaults() *DeleteLookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete look params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete look params
func (o *DeleteLookParams) WithTimeout(timeout time.Duration) *DeleteLookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete look params
func (o *DeleteLookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete look params
func (o *DeleteLookParams) WithContext(ctx context.Context) *DeleteLookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete look params
func (o *DeleteLookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete look params
func (o *DeleteLookParams) WithHTTPClient(client *http.Client) *DeleteLookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete look params
func (o *DeleteLookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLookID adds the lookID to the delete look params
func (o *DeleteLookParams) WithLookID(lookID int64) *DeleteLookParams {
	o.SetLookID(lookID)
	return o
}

// SetLookID adds the lookId to the delete look params
func (o *DeleteLookParams) SetLookID(lookID int64) {
	o.LookID = lookID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param look_id
	if err := r.SetPathParam("look_id", swag.FormatInt64(o.LookID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
