// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewAllUserCredentialsEmbedsParams creates a new AllUserCredentialsEmbedsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllUserCredentialsEmbedsParams() *AllUserCredentialsEmbedsParams {
	return &AllUserCredentialsEmbedsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllUserCredentialsEmbedsParamsWithTimeout creates a new AllUserCredentialsEmbedsParams object
// with the ability to set a timeout on a request.
func NewAllUserCredentialsEmbedsParamsWithTimeout(timeout time.Duration) *AllUserCredentialsEmbedsParams {
	return &AllUserCredentialsEmbedsParams{
		timeout: timeout,
	}
}

// NewAllUserCredentialsEmbedsParamsWithContext creates a new AllUserCredentialsEmbedsParams object
// with the ability to set a context for a request.
func NewAllUserCredentialsEmbedsParamsWithContext(ctx context.Context) *AllUserCredentialsEmbedsParams {
	return &AllUserCredentialsEmbedsParams{
		Context: ctx,
	}
}

// NewAllUserCredentialsEmbedsParamsWithHTTPClient creates a new AllUserCredentialsEmbedsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllUserCredentialsEmbedsParamsWithHTTPClient(client *http.Client) *AllUserCredentialsEmbedsParams {
	return &AllUserCredentialsEmbedsParams{
		HTTPClient: client,
	}
}

/* AllUserCredentialsEmbedsParams contains all the parameters to send to the API endpoint
   for the all user credentials embeds operation.

   Typically these are written to a http.Request.
*/
type AllUserCredentialsEmbedsParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* UserID.

	   id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all user credentials embeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserCredentialsEmbedsParams) WithDefaults() *AllUserCredentialsEmbedsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all user credentials embeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserCredentialsEmbedsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) WithTimeout(timeout time.Duration) *AllUserCredentialsEmbedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) WithContext(ctx context.Context) *AllUserCredentialsEmbedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) WithHTTPClient(client *http.Client) *AllUserCredentialsEmbedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) WithFields(fields *string) *AllUserCredentialsEmbedsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) WithUserID(userID int64) *AllUserCredentialsEmbedsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the all user credentials embeds params
func (o *AllUserCredentialsEmbedsParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AllUserCredentialsEmbedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
