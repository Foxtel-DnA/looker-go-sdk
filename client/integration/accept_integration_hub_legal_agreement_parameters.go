// Code generated by go-swagger; DO NOT EDIT.

package integration

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

// NewAcceptIntegrationHubLegalAgreementParams creates a new AcceptIntegrationHubLegalAgreementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptIntegrationHubLegalAgreementParams() *AcceptIntegrationHubLegalAgreementParams {
	return &AcceptIntegrationHubLegalAgreementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptIntegrationHubLegalAgreementParamsWithTimeout creates a new AcceptIntegrationHubLegalAgreementParams object
// with the ability to set a timeout on a request.
func NewAcceptIntegrationHubLegalAgreementParamsWithTimeout(timeout time.Duration) *AcceptIntegrationHubLegalAgreementParams {
	return &AcceptIntegrationHubLegalAgreementParams{
		timeout: timeout,
	}
}

// NewAcceptIntegrationHubLegalAgreementParamsWithContext creates a new AcceptIntegrationHubLegalAgreementParams object
// with the ability to set a context for a request.
func NewAcceptIntegrationHubLegalAgreementParamsWithContext(ctx context.Context) *AcceptIntegrationHubLegalAgreementParams {
	return &AcceptIntegrationHubLegalAgreementParams{
		Context: ctx,
	}
}

// NewAcceptIntegrationHubLegalAgreementParamsWithHTTPClient creates a new AcceptIntegrationHubLegalAgreementParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptIntegrationHubLegalAgreementParamsWithHTTPClient(client *http.Client) *AcceptIntegrationHubLegalAgreementParams {
	return &AcceptIntegrationHubLegalAgreementParams{
		HTTPClient: client,
	}
}

/* AcceptIntegrationHubLegalAgreementParams contains all the parameters to send to the API endpoint
   for the accept integration hub legal agreement operation.

   Typically these are written to a http.Request.
*/
type AcceptIntegrationHubLegalAgreementParams struct {

	/* IntegrationHubID.

	   Id of integration_hub

	   Format: int64
	*/
	IntegrationHubID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept integration hub legal agreement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptIntegrationHubLegalAgreementParams) WithDefaults() *AcceptIntegrationHubLegalAgreementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept integration hub legal agreement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptIntegrationHubLegalAgreementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) WithTimeout(timeout time.Duration) *AcceptIntegrationHubLegalAgreementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) WithContext(ctx context.Context) *AcceptIntegrationHubLegalAgreementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) WithHTTPClient(client *http.Client) *AcceptIntegrationHubLegalAgreementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationHubID adds the integrationHubID to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) WithIntegrationHubID(integrationHubID int64) *AcceptIntegrationHubLegalAgreementParams {
	o.SetIntegrationHubID(integrationHubID)
	return o
}

// SetIntegrationHubID adds the integrationHubId to the accept integration hub legal agreement params
func (o *AcceptIntegrationHubLegalAgreementParams) SetIntegrationHubID(integrationHubID int64) {
	o.IntegrationHubID = integrationHubID
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptIntegrationHubLegalAgreementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integration_hub_id
	if err := r.SetPathParam("integration_hub_id", swag.FormatInt64(o.IntegrationHubID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
