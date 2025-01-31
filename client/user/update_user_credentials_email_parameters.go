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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewUpdateUserCredentialsEmailParams creates a new UpdateUserCredentialsEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserCredentialsEmailParams() *UpdateUserCredentialsEmailParams {
	return &UpdateUserCredentialsEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserCredentialsEmailParamsWithTimeout creates a new UpdateUserCredentialsEmailParams object
// with the ability to set a timeout on a request.
func NewUpdateUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *UpdateUserCredentialsEmailParams {
	return &UpdateUserCredentialsEmailParams{
		timeout: timeout,
	}
}

// NewUpdateUserCredentialsEmailParamsWithContext creates a new UpdateUserCredentialsEmailParams object
// with the ability to set a context for a request.
func NewUpdateUserCredentialsEmailParamsWithContext(ctx context.Context) *UpdateUserCredentialsEmailParams {
	return &UpdateUserCredentialsEmailParams{
		Context: ctx,
	}
}

// NewUpdateUserCredentialsEmailParamsWithHTTPClient creates a new UpdateUserCredentialsEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserCredentialsEmailParamsWithHTTPClient(client *http.Client) *UpdateUserCredentialsEmailParams {
	return &UpdateUserCredentialsEmailParams{
		HTTPClient: client,
	}
}

/* UpdateUserCredentialsEmailParams contains all the parameters to send to the API endpoint
   for the update user credentials email operation.

   Typically these are written to a http.Request.
*/
type UpdateUserCredentialsEmailParams struct {

	/* Body.

	   Email/Password Credential
	*/
	Body *models.CredentialsEmail

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

// WithDefaults hydrates default values in the update user credentials email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserCredentialsEmailParams) WithDefaults() *UpdateUserCredentialsEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user credentials email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserCredentialsEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithTimeout(timeout time.Duration) *UpdateUserCredentialsEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithContext(ctx context.Context) *UpdateUserCredentialsEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithHTTPClient(client *http.Client) *UpdateUserCredentialsEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithBody(body *models.CredentialsEmail) *UpdateUserCredentialsEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetBody(body *models.CredentialsEmail) {
	o.Body = body
}

// WithFields adds the fields to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithFields(fields *string) *UpdateUserCredentialsEmailParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithUserID(userID int64) *UpdateUserCredentialsEmailParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
