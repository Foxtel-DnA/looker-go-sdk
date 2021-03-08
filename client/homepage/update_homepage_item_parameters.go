// Code generated by go-swagger; DO NOT EDIT.

package homepage

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

// NewUpdateHomepageItemParams creates a new UpdateHomepageItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHomepageItemParams() *UpdateHomepageItemParams {
	return &UpdateHomepageItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHomepageItemParamsWithTimeout creates a new UpdateHomepageItemParams object
// with the ability to set a timeout on a request.
func NewUpdateHomepageItemParamsWithTimeout(timeout time.Duration) *UpdateHomepageItemParams {
	return &UpdateHomepageItemParams{
		timeout: timeout,
	}
}

// NewUpdateHomepageItemParamsWithContext creates a new UpdateHomepageItemParams object
// with the ability to set a context for a request.
func NewUpdateHomepageItemParamsWithContext(ctx context.Context) *UpdateHomepageItemParams {
	return &UpdateHomepageItemParams{
		Context: ctx,
	}
}

// NewUpdateHomepageItemParamsWithHTTPClient creates a new UpdateHomepageItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHomepageItemParamsWithHTTPClient(client *http.Client) *UpdateHomepageItemParams {
	return &UpdateHomepageItemParams{
		HTTPClient: client,
	}
}

/* UpdateHomepageItemParams contains all the parameters to send to the API endpoint
   for the update homepage item operation.

   Typically these are written to a http.Request.
*/
type UpdateHomepageItemParams struct {

	/* Body.

	   Homepage Item
	*/
	Body *models.HomepageItem

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* HomepageItemID.

	   Id of homepage item

	   Format: int64
	*/
	HomepageItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update homepage item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageItemParams) WithDefaults() *UpdateHomepageItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update homepage item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHomepageItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update homepage item params
func (o *UpdateHomepageItemParams) WithTimeout(timeout time.Duration) *UpdateHomepageItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update homepage item params
func (o *UpdateHomepageItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update homepage item params
func (o *UpdateHomepageItemParams) WithContext(ctx context.Context) *UpdateHomepageItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update homepage item params
func (o *UpdateHomepageItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update homepage item params
func (o *UpdateHomepageItemParams) WithHTTPClient(client *http.Client) *UpdateHomepageItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update homepage item params
func (o *UpdateHomepageItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update homepage item params
func (o *UpdateHomepageItemParams) WithBody(body *models.HomepageItem) *UpdateHomepageItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update homepage item params
func (o *UpdateHomepageItemParams) SetBody(body *models.HomepageItem) {
	o.Body = body
}

// WithFields adds the fields to the update homepage item params
func (o *UpdateHomepageItemParams) WithFields(fields *string) *UpdateHomepageItemParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update homepage item params
func (o *UpdateHomepageItemParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHomepageItemID adds the homepageItemID to the update homepage item params
func (o *UpdateHomepageItemParams) WithHomepageItemID(homepageItemID int64) *UpdateHomepageItemParams {
	o.SetHomepageItemID(homepageItemID)
	return o
}

// SetHomepageItemID adds the homepageItemId to the update homepage item params
func (o *UpdateHomepageItemParams) SetHomepageItemID(homepageItemID int64) {
	o.HomepageItemID = homepageItemID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHomepageItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homepage_item_id
	if err := r.SetPathParam("homepage_item_id", swag.FormatInt64(o.HomepageItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
