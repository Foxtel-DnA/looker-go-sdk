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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserCredentialsLookerOpenidParams creates a new DeleteUserCredentialsLookerOpenidParams object
// with the default values initialized.
func NewDeleteUserCredentialsLookerOpenidParams() *DeleteUserCredentialsLookerOpenidParams {
	var ()
	return &DeleteUserCredentialsLookerOpenidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsLookerOpenidParamsWithTimeout creates a new DeleteUserCredentialsLookerOpenidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserCredentialsLookerOpenidParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsLookerOpenidParams {
	var ()
	return &DeleteUserCredentialsLookerOpenidParams{

		timeout: timeout,
	}
}

// NewDeleteUserCredentialsLookerOpenidParamsWithContext creates a new DeleteUserCredentialsLookerOpenidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserCredentialsLookerOpenidParamsWithContext(ctx context.Context) *DeleteUserCredentialsLookerOpenidParams {
	var ()
	return &DeleteUserCredentialsLookerOpenidParams{

		Context: ctx,
	}
}

// NewDeleteUserCredentialsLookerOpenidParamsWithHTTPClient creates a new DeleteUserCredentialsLookerOpenidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserCredentialsLookerOpenidParamsWithHTTPClient(client *http.Client) *DeleteUserCredentialsLookerOpenidParams {
	var ()
	return &DeleteUserCredentialsLookerOpenidParams{
		HTTPClient: client,
	}
}

/*DeleteUserCredentialsLookerOpenidParams contains all the parameters to send to the API endpoint
for the delete user credentials looker openid operation typically these are written to a http.Request
*/
type DeleteUserCredentialsLookerOpenidParams struct {

	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) WithTimeout(timeout time.Duration) *DeleteUserCredentialsLookerOpenidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) WithContext(ctx context.Context) *DeleteUserCredentialsLookerOpenidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) WithHTTPClient(client *http.Client) *DeleteUserCredentialsLookerOpenidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) WithUserID(userID int64) *DeleteUserCredentialsLookerOpenidParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user credentials looker openid params
func (o *DeleteUserCredentialsLookerOpenidParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsLookerOpenidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}