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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewScheduledPlanRunOnceByIDParams creates a new ScheduledPlanRunOnceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduledPlanRunOnceByIDParams() *ScheduledPlanRunOnceByIDParams {
	return &ScheduledPlanRunOnceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduledPlanRunOnceByIDParamsWithTimeout creates a new ScheduledPlanRunOnceByIDParams object
// with the ability to set a timeout on a request.
func NewScheduledPlanRunOnceByIDParamsWithTimeout(timeout time.Duration) *ScheduledPlanRunOnceByIDParams {
	return &ScheduledPlanRunOnceByIDParams{
		timeout: timeout,
	}
}

// NewScheduledPlanRunOnceByIDParamsWithContext creates a new ScheduledPlanRunOnceByIDParams object
// with the ability to set a context for a request.
func NewScheduledPlanRunOnceByIDParamsWithContext(ctx context.Context) *ScheduledPlanRunOnceByIDParams {
	return &ScheduledPlanRunOnceByIDParams{
		Context: ctx,
	}
}

// NewScheduledPlanRunOnceByIDParamsWithHTTPClient creates a new ScheduledPlanRunOnceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduledPlanRunOnceByIDParamsWithHTTPClient(client *http.Client) *ScheduledPlanRunOnceByIDParams {
	return &ScheduledPlanRunOnceByIDParams{
		HTTPClient: client,
	}
}

/* ScheduledPlanRunOnceByIDParams contains all the parameters to send to the API endpoint
   for the scheduled plan run once by id operation.

   Typically these are written to a http.Request.
*/
type ScheduledPlanRunOnceByIDParams struct {

	/* Body.

	   Property values to apply to the newly copied scheduled plan before running it
	*/
	Body *models.WriteScheduledPlan

	/* ScheduledPlanID.

	   Id of schedule plan to copy and run

	   Format: int64
	*/
	ScheduledPlanID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scheduled plan run once by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduledPlanRunOnceByIDParams) WithDefaults() *ScheduledPlanRunOnceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scheduled plan run once by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduledPlanRunOnceByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) WithTimeout(timeout time.Duration) *ScheduledPlanRunOnceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) WithContext(ctx context.Context) *ScheduledPlanRunOnceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) WithHTTPClient(client *http.Client) *ScheduledPlanRunOnceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) WithBody(body *models.WriteScheduledPlan) *ScheduledPlanRunOnceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) SetBody(body *models.WriteScheduledPlan) {
	o.Body = body
}

// WithScheduledPlanID adds the scheduledPlanID to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) WithScheduledPlanID(scheduledPlanID int64) *ScheduledPlanRunOnceByIDParams {
	o.SetScheduledPlanID(scheduledPlanID)
	return o
}

// SetScheduledPlanID adds the scheduledPlanId to the scheduled plan run once by id params
func (o *ScheduledPlanRunOnceByIDParams) SetScheduledPlanID(scheduledPlanID int64) {
	o.ScheduledPlanID = scheduledPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduledPlanRunOnceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scheduled_plan_id
	if err := r.SetPathParam("scheduled_plan_id", swag.FormatInt64(o.ScheduledPlanID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
