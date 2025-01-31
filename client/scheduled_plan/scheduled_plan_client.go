// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scheduled plan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scheduled plan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllScheduledPlans(params *AllScheduledPlansParams, opts ...ClientOption) (*AllScheduledPlansOK, error)

	CreateScheduledPlan(params *CreateScheduledPlanParams, opts ...ClientOption) (*CreateScheduledPlanOK, error)

	DeleteScheduledPlan(params *DeleteScheduledPlanParams, opts ...ClientOption) (*DeleteScheduledPlanNoContent, error)

	ScheduledPlan(params *ScheduledPlanParams, opts ...ClientOption) (*ScheduledPlanOK, error)

	ScheduledPlanRunOnce(params *ScheduledPlanRunOnceParams, opts ...ClientOption) (*ScheduledPlanRunOnceOK, error)

	ScheduledPlanRunOnceByID(params *ScheduledPlanRunOnceByIDParams, opts ...ClientOption) (*ScheduledPlanRunOnceByIDOK, error)

	ScheduledPlansForDashboard(params *ScheduledPlansForDashboardParams, opts ...ClientOption) (*ScheduledPlansForDashboardOK, error)

	ScheduledPlansForLook(params *ScheduledPlansForLookParams, opts ...ClientOption) (*ScheduledPlansForLookOK, error)

	ScheduledPlansForLookmlDashboard(params *ScheduledPlansForLookmlDashboardParams, opts ...ClientOption) (*ScheduledPlansForLookmlDashboardOK, error)

	ScheduledPlansForSpace(params *ScheduledPlansForSpaceParams, opts ...ClientOption) (*ScheduledPlansForSpaceOK, error)

	UpdateScheduledPlan(params *UpdateScheduledPlanParams, opts ...ClientOption) (*UpdateScheduledPlanOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AllScheduledPlans gets all scheduled plans

  ### List All Scheduled Plans

Returns all scheduled plans which belong to the caller or given user.

If no user_id is provided, this function returns the scheduled plans owned by the caller.


To list all schedules for all users, pass `all_users=true`.


The caller must have `see_schedules` permission to see other users' scheduled plans.



*/
func (a *Client) AllScheduledPlans(params *AllScheduledPlansParams, opts ...ClientOption) (*AllScheduledPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllScheduledPlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_scheduled_plans",
		Method:             "GET",
		PathPattern:        "/scheduled_plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllScheduledPlansReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllScheduledPlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_scheduled_plans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateScheduledPlan creates scheduled plan

  ### Create a Scheduled Plan

Create a scheduled plan to render a Look or Dashboard on a recurring schedule.

To create a scheduled plan, you MUST provide values for the following fields:
`name`
and
`look_id`, `dashboard_id`, `lookml_dashboard_id`, or `query_id`
and
`cron_tab` or `datagroup`
and
at least one scheduled_plan_destination

A scheduled plan MUST have at least one scheduled_plan_destination defined.

When `look_id` is set, `require_no_results`, `require_results`, and `require_change` are all required.

If `create_scheduled_plan` fails with a 422 error, be sure to look at the error messages in the response which will explain exactly what fields are missing or values that are incompatible.

The queries that provide the data for the look or dashboard are run in the context of user account that owns the scheduled plan.

When `run_as_recipient` is `false` or not specified, the queries that provide the data for the
look or dashboard are run in the context of user account that owns the scheduled plan.

When `run_as_recipient` is `true` and all the email recipients are Looker user accounts, the
queries are run in the context of each recipient, so different recipients may see different
data from the same scheduled render of a look or dashboard. For more details, see [Run As Recipient](https://looker.com/docs/r/admin/run-as-recipient).

Admins can create and modify scheduled plans on behalf of other users by specifying a user id.
Non-admin users may not create or modify scheduled plans by or for other users.

#### Email Permissions:

For details about permissions required to schedule delivery to email and the safeguards
Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).


#### Scheduled Plan Destination Formats

Scheduled plan destinations must specify the data format to produce and send to the destination.

Formats:

| format | Description
| :-----------: | :--- |
| json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata.
| json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query
| inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination.
| csv | Comma separated values with a header
| txt | Tab separated values with a header
| html | Simple html
| xlsx | MS Excel spreadsheet
| wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document
| assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document
| wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image
||

Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.



*/
func (a *Client) CreateScheduledPlan(params *CreateScheduledPlanParams, opts ...ClientOption) (*CreateScheduledPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScheduledPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_scheduled_plan",
		Method:             "POST",
		PathPattern:        "/scheduled_plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateScheduledPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateScheduledPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_scheduled_plan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteScheduledPlan deletes scheduled plan

  ### Delete a Scheduled Plan

Normal users can only delete their own scheduled plans.
Admins can delete other users' scheduled plans.
This delete cannot be undone.

*/
func (a *Client) DeleteScheduledPlan(params *DeleteScheduledPlanParams, opts ...ClientOption) (*DeleteScheduledPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteScheduledPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_scheduled_plan",
		Method:             "DELETE",
		PathPattern:        "/scheduled_plans/{scheduled_plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScheduledPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteScheduledPlanNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_scheduled_plan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlan gets scheduled plan

  ### Get Information About a Scheduled Plan

Admins can fetch information about other users' Scheduled Plans.

*/
func (a *Client) ScheduledPlan(params *ScheduledPlanParams, opts ...ClientOption) (*ScheduledPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plan",
		Method:             "GET",
		PathPattern:        "/scheduled_plans/{scheduled_plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlanRunOnce runs scheduled plan once

  ### Run a Scheduled Plan Immediately

Create a scheduled plan that runs only once, and immediately.

This can be useful for testing a Scheduled Plan before committing to a production schedule.

Admins can create scheduled plans on behalf of other users by specifying a user id.

This API is rate limited to prevent it from being used for relay spam or DoS attacks

#### Email Permissions:

For details about permissions required to schedule delivery to email and the safeguards
Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).


#### Scheduled Plan Destination Formats

Scheduled plan destinations must specify the data format to produce and send to the destination.

Formats:

| format | Description
| :-----------: | :--- |
| json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata.
| json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query
| inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination.
| csv | Comma separated values with a header
| txt | Tab separated values with a header
| html | Simple html
| xlsx | MS Excel spreadsheet
| wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document
| assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document
| wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image
||

Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.



*/
func (a *Client) ScheduledPlanRunOnce(params *ScheduledPlanRunOnceParams, opts ...ClientOption) (*ScheduledPlanRunOnceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlanRunOnceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plan_run_once",
		Method:             "POST",
		PathPattern:        "/scheduled_plans/run_once",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlanRunOnceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlanRunOnceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plan_run_once: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlanRunOnceByID runs scheduled plan once by Id

  ### Run a Scheduled Plan By Id Immediately
This function creates a run-once schedule plan based on an existing scheduled plan,
applies modifications (if any) to the new scheduled plan, and runs the new schedule plan immediately.
This can be useful for testing modifications to an existing scheduled plan before committing to a production schedule.

This function internally performs the following operations:

1. Copies the properties of the existing scheduled plan into a new scheduled plan
2. Copies any properties passed in the JSON body of this request into the new scheduled plan (replacing the original values)
3. Creates the new scheduled plan
4. Runs the new scheduled plan

The original scheduled plan is not modified by this operation.
Admins can create, modify, and run scheduled plans on behalf of other users by specifying a user id.
Non-admins can only create, modify, and run their own scheduled plans.

#### Email Permissions:

For details about permissions required to schedule delivery to email and the safeguards
Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).


#### Scheduled Plan Destination Formats

Scheduled plan destinations must specify the data format to produce and send to the destination.

Formats:

| format | Description
| :-----------: | :--- |
| json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata.
| json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query
| inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination.
| csv | Comma separated values with a header
| txt | Tab separated values with a header
| html | Simple html
| xlsx | MS Excel spreadsheet
| wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document
| assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document
| wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image
||

Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.



This API is rate limited to prevent it from being used for relay spam or DoS attacks


*/
func (a *Client) ScheduledPlanRunOnceByID(params *ScheduledPlanRunOnceByIDParams, opts ...ClientOption) (*ScheduledPlanRunOnceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlanRunOnceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plan_run_once_by_id",
		Method:             "POST",
		PathPattern:        "/scheduled_plans/{scheduled_plan_id}/run_once",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlanRunOnceByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlanRunOnceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plan_run_once_by_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlansForDashboard scheduleds plans for dashboard

  ### Get Scheduled Plans for a Dashboard

Returns all scheduled plans for a dashboard which belong to the caller or given user.

If no user_id is provided, this function returns the scheduled plans owned by the caller.


To list all schedules for all users, pass `all_users=true`.


The caller must have `see_schedules` permission to see other users' scheduled plans.



*/
func (a *Client) ScheduledPlansForDashboard(params *ScheduledPlansForDashboardParams, opts ...ClientOption) (*ScheduledPlansForDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlansForDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plans_for_dashboard",
		Method:             "GET",
		PathPattern:        "/scheduled_plans/dashboard/{dashboard_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlansForDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlansForDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plans_for_dashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlansForLook scheduleds plans for look

  ### Get Scheduled Plans for a Look

Returns all scheduled plans for a look which belong to the caller or given user.

If no user_id is provided, this function returns the scheduled plans owned by the caller.


To list all schedules for all users, pass `all_users=true`.


The caller must have `see_schedules` permission to see other users' scheduled plans.



*/
func (a *Client) ScheduledPlansForLook(params *ScheduledPlansForLookParams, opts ...ClientOption) (*ScheduledPlansForLookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlansForLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plans_for_look",
		Method:             "GET",
		PathPattern:        "/scheduled_plans/look/{look_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlansForLookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlansForLookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plans_for_look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlansForLookmlDashboard scheduleds plans for look m l dashboard

  ### Get Scheduled Plans for a LookML Dashboard

Returns all scheduled plans for a LookML Dashboard which belong to the caller or given user.

If no user_id is provided, this function returns the scheduled plans owned by the caller.


To list all schedules for all users, pass `all_users=true`.


The caller must have `see_schedules` permission to see other users' scheduled plans.



*/
func (a *Client) ScheduledPlansForLookmlDashboard(params *ScheduledPlansForLookmlDashboardParams, opts ...ClientOption) (*ScheduledPlansForLookmlDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlansForLookmlDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plans_for_lookml_dashboard",
		Method:             "GET",
		PathPattern:        "/scheduled_plans/lookml_dashboard/{lookml_dashboard_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlansForLookmlDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlansForLookmlDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plans_for_lookml_dashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScheduledPlansForSpace scheduleds plans for space

  ### Get Scheduled Plans for a Space

Returns scheduled plans owned by the caller for a given space id.

*/
func (a *Client) ScheduledPlansForSpace(params *ScheduledPlansForSpaceParams, opts ...ClientOption) (*ScheduledPlansForSpaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduledPlansForSpaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "scheduled_plans_for_space",
		Method:             "GET",
		PathPattern:        "/scheduled_plans/space/{space_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScheduledPlansForSpaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduledPlansForSpaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for scheduled_plans_for_space: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateScheduledPlan updates scheduled plan

  ### Update a Scheduled Plan

Admins can update other users' Scheduled Plans.

Note: Any scheduled plan destinations specified in an update will **replace** all scheduled plan destinations
currently defined for the scheduled plan.

For Example: If a scheduled plan has destinations A, B, and C, and you call update on this scheduled plan
specifying only B in the destinations, then destinations A and C will be deleted by the update.

Updating a scheduled plan to assign null or an empty array to the scheduled_plan_destinations property is an error, as a scheduled plan must always have at least one destination.

If you omit the scheduled_plan_destinations property from the object passed to update, then the destinations
defined on the original scheduled plan will remain unchanged.

#### Email Permissions:

For details about permissions required to schedule delivery to email and the safeguards
Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).


#### Scheduled Plan Destination Formats

Scheduled plan destinations must specify the data format to produce and send to the destination.

Formats:

| format | Description
| :-----------: | :--- |
| json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata.
| json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query
| inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination.
| csv | Comma separated values with a header
| txt | Tab separated values with a header
| html | Simple html
| xlsx | MS Excel spreadsheet
| wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document
| assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document
| wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image
||

Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.



*/
func (a *Client) UpdateScheduledPlan(params *UpdateScheduledPlanParams, opts ...ClientOption) (*UpdateScheduledPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScheduledPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_scheduled_plan",
		Method:             "PATCH",
		PathPattern:        "/scheduled_plans/{scheduled_plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateScheduledPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateScheduledPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_scheduled_plan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
