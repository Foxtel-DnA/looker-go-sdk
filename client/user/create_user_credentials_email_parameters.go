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

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewCreateUserCredentialsEmailParams creates a new CreateUserCredentialsEmailParams object
// with the default values initialized.
func NewCreateUserCredentialsEmailParams() *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsEmailParamsWithTimeout creates a new CreateUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{

		timeout: timeout,
	}
}

// NewCreateUserCredentialsEmailParamsWithContext creates a new CreateUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserCredentialsEmailParamsWithContext(ctx context.Context) *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{

		Context: ctx,
	}
}

// NewCreateUserCredentialsEmailParamsWithHTTPClient creates a new CreateUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserCredentialsEmailParamsWithHTTPClient(client *http.Client) *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{
		HTTPClient: client,
	}
}

/*CreateUserCredentialsEmailParams contains all the parameters to send to the API endpoint
for the create user credentials email operation typically these are written to a http.Request
*/
type CreateUserCredentialsEmailParams struct {

	/*Body
	  Email/Password Credential

	*/
	Body *models.CredentialsEmail
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithTimeout(timeout time.Duration) *CreateUserCredentialsEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithContext(ctx context.Context) *CreateUserCredentialsEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithHTTPClient(client *http.Client) *CreateUserCredentialsEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithBody(body *models.CredentialsEmail) *CreateUserCredentialsEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetBody(body *models.CredentialsEmail) {
	o.Body = body
}

// WithFields adds the fields to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithFields(fields *string) *CreateUserCredentialsEmailParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithUserID(userID int64) *CreateUserCredentialsEmailParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
