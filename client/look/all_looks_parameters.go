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
)

// NewAllLooksParams creates a new AllLooksParams object
// with the default values initialized.
func NewAllLooksParams() *AllLooksParams {
	var ()
	return &AllLooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllLooksParamsWithTimeout creates a new AllLooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllLooksParamsWithTimeout(timeout time.Duration) *AllLooksParams {
	var ()
	return &AllLooksParams{

		timeout: timeout,
	}
}

// NewAllLooksParamsWithContext creates a new AllLooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllLooksParamsWithContext(ctx context.Context) *AllLooksParams {
	var ()
	return &AllLooksParams{

		Context: ctx,
	}
}

// NewAllLooksParamsWithHTTPClient creates a new AllLooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllLooksParamsWithHTTPClient(client *http.Client) *AllLooksParams {
	var ()
	return &AllLooksParams{
		HTTPClient: client,
	}
}

/*AllLooksParams contains all the parameters to send to the API endpoint
for the all looks operation typically these are written to a http.Request
*/
type AllLooksParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all looks params
func (o *AllLooksParams) WithTimeout(timeout time.Duration) *AllLooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all looks params
func (o *AllLooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all looks params
func (o *AllLooksParams) WithContext(ctx context.Context) *AllLooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all looks params
func (o *AllLooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all looks params
func (o *AllLooksParams) WithHTTPClient(client *http.Client) *AllLooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all looks params
func (o *AllLooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all looks params
func (o *AllLooksParams) WithFields(fields *string) *AllLooksParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all looks params
func (o *AllLooksParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllLooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}