// Code generated by go-swagger; DO NOT EDIT.

package folder

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
)

// NewDeleteFolderParams creates a new DeleteFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFolderParams() *DeleteFolderParams {
	return &DeleteFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFolderParamsWithTimeout creates a new DeleteFolderParams object
// with the ability to set a timeout on a request.
func NewDeleteFolderParamsWithTimeout(timeout time.Duration) *DeleteFolderParams {
	return &DeleteFolderParams{
		timeout: timeout,
	}
}

// NewDeleteFolderParamsWithContext creates a new DeleteFolderParams object
// with the ability to set a context for a request.
func NewDeleteFolderParamsWithContext(ctx context.Context) *DeleteFolderParams {
	return &DeleteFolderParams{
		Context: ctx,
	}
}

// NewDeleteFolderParamsWithHTTPClient creates a new DeleteFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFolderParamsWithHTTPClient(client *http.Client) *DeleteFolderParams {
	return &DeleteFolderParams{
		HTTPClient: client,
	}
}

/* DeleteFolderParams contains all the parameters to send to the API endpoint
   for the delete folder operation.

   Typically these are written to a http.Request.
*/
type DeleteFolderParams struct {

	/* FolderID.

	   Id of folder
	*/
	FolderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFolderParams) WithDefaults() *DeleteFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFolderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete folder params
func (o *DeleteFolderParams) WithTimeout(timeout time.Duration) *DeleteFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete folder params
func (o *DeleteFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete folder params
func (o *DeleteFolderParams) WithContext(ctx context.Context) *DeleteFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete folder params
func (o *DeleteFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete folder params
func (o *DeleteFolderParams) WithHTTPClient(client *http.Client) *DeleteFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete folder params
func (o *DeleteFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderID adds the folderID to the delete folder params
func (o *DeleteFolderParams) WithFolderID(folderID string) *DeleteFolderParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the delete folder params
func (o *DeleteFolderParams) SetFolderID(folderID string) {
	o.FolderID = folderID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param folder_id
	if err := r.SetPathParam("folder_id", o.FolderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
