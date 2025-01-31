// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewDeleteGroupFromGroupParams creates a new DeleteGroupFromGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGroupFromGroupParams() *DeleteGroupFromGroupParams {
	return &DeleteGroupFromGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupFromGroupParamsWithTimeout creates a new DeleteGroupFromGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteGroupFromGroupParamsWithTimeout(timeout time.Duration) *DeleteGroupFromGroupParams {
	return &DeleteGroupFromGroupParams{
		timeout: timeout,
	}
}

// NewDeleteGroupFromGroupParamsWithContext creates a new DeleteGroupFromGroupParams object
// with the ability to set a context for a request.
func NewDeleteGroupFromGroupParamsWithContext(ctx context.Context) *DeleteGroupFromGroupParams {
	return &DeleteGroupFromGroupParams{
		Context: ctx,
	}
}

// NewDeleteGroupFromGroupParamsWithHTTPClient creates a new DeleteGroupFromGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGroupFromGroupParamsWithHTTPClient(client *http.Client) *DeleteGroupFromGroupParams {
	return &DeleteGroupFromGroupParams{
		HTTPClient: client,
	}
}

/* DeleteGroupFromGroupParams contains all the parameters to send to the API endpoint
   for the delete group from group operation.

   Typically these are written to a http.Request.
*/
type DeleteGroupFromGroupParams struct {

	/* DeletingGroupID.

	   Id of group to delete

	   Format: int64
	*/
	DeletingGroupID int64

	/* GroupID.

	   Id of group

	   Format: int64
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete group from group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupFromGroupParams) WithDefaults() *DeleteGroupFromGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete group from group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupFromGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete group from group params
func (o *DeleteGroupFromGroupParams) WithTimeout(timeout time.Duration) *DeleteGroupFromGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete group from group params
func (o *DeleteGroupFromGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete group from group params
func (o *DeleteGroupFromGroupParams) WithContext(ctx context.Context) *DeleteGroupFromGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete group from group params
func (o *DeleteGroupFromGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete group from group params
func (o *DeleteGroupFromGroupParams) WithHTTPClient(client *http.Client) *DeleteGroupFromGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete group from group params
func (o *DeleteGroupFromGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeletingGroupID adds the deletingGroupID to the delete group from group params
func (o *DeleteGroupFromGroupParams) WithDeletingGroupID(deletingGroupID int64) *DeleteGroupFromGroupParams {
	o.SetDeletingGroupID(deletingGroupID)
	return o
}

// SetDeletingGroupID adds the deletingGroupId to the delete group from group params
func (o *DeleteGroupFromGroupParams) SetDeletingGroupID(deletingGroupID int64) {
	o.DeletingGroupID = deletingGroupID
}

// WithGroupID adds the groupID to the delete group from group params
func (o *DeleteGroupFromGroupParams) WithGroupID(groupID int64) *DeleteGroupFromGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete group from group params
func (o *DeleteGroupFromGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupFromGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deleting_group_id
	if err := r.SetPathParam("deleting_group_id", swag.FormatInt64(o.DeletingGroupID)); err != nil {
		return err
	}

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
