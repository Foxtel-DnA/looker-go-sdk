// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewCreateUserAttributeParams creates a new CreateUserAttributeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserAttributeParams() *CreateUserAttributeParams {
	return &CreateUserAttributeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserAttributeParamsWithTimeout creates a new CreateUserAttributeParams object
// with the ability to set a timeout on a request.
func NewCreateUserAttributeParamsWithTimeout(timeout time.Duration) *CreateUserAttributeParams {
	return &CreateUserAttributeParams{
		timeout: timeout,
	}
}

// NewCreateUserAttributeParamsWithContext creates a new CreateUserAttributeParams object
// with the ability to set a context for a request.
func NewCreateUserAttributeParamsWithContext(ctx context.Context) *CreateUserAttributeParams {
	return &CreateUserAttributeParams{
		Context: ctx,
	}
}

// NewCreateUserAttributeParamsWithHTTPClient creates a new CreateUserAttributeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserAttributeParamsWithHTTPClient(client *http.Client) *CreateUserAttributeParams {
	return &CreateUserAttributeParams{
		HTTPClient: client,
	}
}

/* CreateUserAttributeParams contains all the parameters to send to the API endpoint
   for the create user attribute operation.

   Typically these are written to a http.Request.
*/
type CreateUserAttributeParams struct {

	/* Body.

	   User Attribute
	*/
	Body *models.UserAttribute

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user attribute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserAttributeParams) WithDefaults() *CreateUserAttributeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user attribute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserAttributeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user attribute params
func (o *CreateUserAttributeParams) WithTimeout(timeout time.Duration) *CreateUserAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user attribute params
func (o *CreateUserAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user attribute params
func (o *CreateUserAttributeParams) WithContext(ctx context.Context) *CreateUserAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user attribute params
func (o *CreateUserAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user attribute params
func (o *CreateUserAttributeParams) WithHTTPClient(client *http.Client) *CreateUserAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user attribute params
func (o *CreateUserAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user attribute params
func (o *CreateUserAttributeParams) WithBody(body *models.UserAttribute) *CreateUserAttributeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user attribute params
func (o *CreateUserAttributeParams) SetBody(body *models.UserAttribute) {
	o.Body = body
}

// WithFields adds the fields to the create user attribute params
func (o *CreateUserAttributeParams) WithFields(fields *string) *CreateUserAttributeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create user attribute params
func (o *CreateUserAttributeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
