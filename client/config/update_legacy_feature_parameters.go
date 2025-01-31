// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewUpdateLegacyFeatureParams creates a new UpdateLegacyFeatureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLegacyFeatureParams() *UpdateLegacyFeatureParams {
	return &UpdateLegacyFeatureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLegacyFeatureParamsWithTimeout creates a new UpdateLegacyFeatureParams object
// with the ability to set a timeout on a request.
func NewUpdateLegacyFeatureParamsWithTimeout(timeout time.Duration) *UpdateLegacyFeatureParams {
	return &UpdateLegacyFeatureParams{
		timeout: timeout,
	}
}

// NewUpdateLegacyFeatureParamsWithContext creates a new UpdateLegacyFeatureParams object
// with the ability to set a context for a request.
func NewUpdateLegacyFeatureParamsWithContext(ctx context.Context) *UpdateLegacyFeatureParams {
	return &UpdateLegacyFeatureParams{
		Context: ctx,
	}
}

// NewUpdateLegacyFeatureParamsWithHTTPClient creates a new UpdateLegacyFeatureParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLegacyFeatureParamsWithHTTPClient(client *http.Client) *UpdateLegacyFeatureParams {
	return &UpdateLegacyFeatureParams{
		HTTPClient: client,
	}
}

/* UpdateLegacyFeatureParams contains all the parameters to send to the API endpoint
   for the update legacy feature operation.

   Typically these are written to a http.Request.
*/
type UpdateLegacyFeatureParams struct {

	/* Body.

	   Legacy Feature
	*/
	Body *models.LegacyFeature

	/* LegacyFeatureID.

	   id of legacy feature

	   Format: int64
	*/
	LegacyFeatureID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update legacy feature params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLegacyFeatureParams) WithDefaults() *UpdateLegacyFeatureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update legacy feature params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLegacyFeatureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update legacy feature params
func (o *UpdateLegacyFeatureParams) WithTimeout(timeout time.Duration) *UpdateLegacyFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update legacy feature params
func (o *UpdateLegacyFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update legacy feature params
func (o *UpdateLegacyFeatureParams) WithContext(ctx context.Context) *UpdateLegacyFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update legacy feature params
func (o *UpdateLegacyFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update legacy feature params
func (o *UpdateLegacyFeatureParams) WithHTTPClient(client *http.Client) *UpdateLegacyFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update legacy feature params
func (o *UpdateLegacyFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update legacy feature params
func (o *UpdateLegacyFeatureParams) WithBody(body *models.LegacyFeature) *UpdateLegacyFeatureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update legacy feature params
func (o *UpdateLegacyFeatureParams) SetBody(body *models.LegacyFeature) {
	o.Body = body
}

// WithLegacyFeatureID adds the legacyFeatureID to the update legacy feature params
func (o *UpdateLegacyFeatureParams) WithLegacyFeatureID(legacyFeatureID int64) *UpdateLegacyFeatureParams {
	o.SetLegacyFeatureID(legacyFeatureID)
	return o
}

// SetLegacyFeatureID adds the legacyFeatureId to the update legacy feature params
func (o *UpdateLegacyFeatureParams) SetLegacyFeatureID(legacyFeatureID int64) {
	o.LegacyFeatureID = legacyFeatureID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLegacyFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param legacy_feature_id
	if err := r.SetPathParam("legacy_feature_id", swag.FormatInt64(o.LegacyFeatureID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
