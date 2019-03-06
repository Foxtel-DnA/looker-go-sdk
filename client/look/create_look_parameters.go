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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewCreateLookParams creates a new CreateLookParams object
// with the default values initialized.
func NewCreateLookParams() *CreateLookParams {
	var ()
	return &CreateLookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLookParamsWithTimeout creates a new CreateLookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLookParamsWithTimeout(timeout time.Duration) *CreateLookParams {
	var ()
	return &CreateLookParams{

		timeout: timeout,
	}
}

// NewCreateLookParamsWithContext creates a new CreateLookParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLookParamsWithContext(ctx context.Context) *CreateLookParams {
	var ()
	return &CreateLookParams{

		Context: ctx,
	}
}

// NewCreateLookParamsWithHTTPClient creates a new CreateLookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLookParamsWithHTTPClient(client *http.Client) *CreateLookParams {
	var ()
	return &CreateLookParams{
		HTTPClient: client,
	}
}

/*CreateLookParams contains all the parameters to send to the API endpoint
for the create look operation typically these are written to a http.Request
*/
type CreateLookParams struct {

	/*Body
	  Look

	*/
	Body *models.LookWithQuery
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create look params
func (o *CreateLookParams) WithTimeout(timeout time.Duration) *CreateLookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create look params
func (o *CreateLookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create look params
func (o *CreateLookParams) WithContext(ctx context.Context) *CreateLookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create look params
func (o *CreateLookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create look params
func (o *CreateLookParams) WithHTTPClient(client *http.Client) *CreateLookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create look params
func (o *CreateLookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create look params
func (o *CreateLookParams) WithBody(body *models.LookWithQuery) *CreateLookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create look params
func (o *CreateLookParams) SetBody(body *models.LookWithQuery) {
	o.Body = body
}

// WithFields adds the fields to the create look params
func (o *CreateLookParams) WithFields(fields *string) *CreateLookParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create look params
func (o *CreateLookParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
