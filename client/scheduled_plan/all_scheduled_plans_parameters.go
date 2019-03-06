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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllScheduledPlansParams creates a new AllScheduledPlansParams object
// with the default values initialized.
func NewAllScheduledPlansParams() *AllScheduledPlansParams {
	var ()
	return &AllScheduledPlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllScheduledPlansParamsWithTimeout creates a new AllScheduledPlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllScheduledPlansParamsWithTimeout(timeout time.Duration) *AllScheduledPlansParams {
	var ()
	return &AllScheduledPlansParams{

		timeout: timeout,
	}
}

// NewAllScheduledPlansParamsWithContext creates a new AllScheduledPlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllScheduledPlansParamsWithContext(ctx context.Context) *AllScheduledPlansParams {
	var ()
	return &AllScheduledPlansParams{

		Context: ctx,
	}
}

// NewAllScheduledPlansParamsWithHTTPClient creates a new AllScheduledPlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllScheduledPlansParamsWithHTTPClient(client *http.Client) *AllScheduledPlansParams {
	var ()
	return &AllScheduledPlansParams{
		HTTPClient: client,
	}
}

/*AllScheduledPlansParams contains all the parameters to send to the API endpoint
for the all scheduled plans operation typically these are written to a http.Request
*/
type AllScheduledPlansParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*UserID
	  User Id (default is requesting user if not specified)

	*/
	UserID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all scheduled plans params
func (o *AllScheduledPlansParams) WithTimeout(timeout time.Duration) *AllScheduledPlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all scheduled plans params
func (o *AllScheduledPlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all scheduled plans params
func (o *AllScheduledPlansParams) WithContext(ctx context.Context) *AllScheduledPlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all scheduled plans params
func (o *AllScheduledPlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all scheduled plans params
func (o *AllScheduledPlansParams) WithHTTPClient(client *http.Client) *AllScheduledPlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all scheduled plans params
func (o *AllScheduledPlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all scheduled plans params
func (o *AllScheduledPlansParams) WithFields(fields *string) *AllScheduledPlansParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all scheduled plans params
func (o *AllScheduledPlansParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the all scheduled plans params
func (o *AllScheduledPlansParams) WithUserID(userID *int64) *AllScheduledPlansParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the all scheduled plans params
func (o *AllScheduledPlansParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AllScheduledPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UserID != nil {

		// query param user_id
		var qrUserID int64
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatInt64(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
