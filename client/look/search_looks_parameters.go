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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchLooksParams creates a new SearchLooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchLooksParams() *SearchLooksParams {
	return &SearchLooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchLooksParamsWithTimeout creates a new SearchLooksParams object
// with the ability to set a timeout on a request.
func NewSearchLooksParamsWithTimeout(timeout time.Duration) *SearchLooksParams {
	return &SearchLooksParams{
		timeout: timeout,
	}
}

// NewSearchLooksParamsWithContext creates a new SearchLooksParams object
// with the ability to set a context for a request.
func NewSearchLooksParamsWithContext(ctx context.Context) *SearchLooksParams {
	return &SearchLooksParams{
		Context: ctx,
	}
}

// NewSearchLooksParamsWithHTTPClient creates a new SearchLooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchLooksParamsWithHTTPClient(client *http.Client) *SearchLooksParams {
	return &SearchLooksParams{
		HTTPClient: client,
	}
}

/* SearchLooksParams contains all the parameters to send to the API endpoint
   for the search looks operation.

   Typically these are written to a http.Request.
*/
type SearchLooksParams struct {

	/* ContentFavoriteID.

	   Select looks with a particular content favorite id

	   Format: int64
	*/
	ContentFavoriteID *int64

	/* Curate.

	   Exclude items that exist only in personal spaces other than the users
	*/
	Curate *bool

	/* Deleted.

	   Select soft-deleted looks
	*/
	Deleted *bool

	/* Description.

	   Match Look description.
	*/
	Description *string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* FilterOr.

	   Combine given search criteria in a boolean OR expression
	*/
	FilterOr *bool

	/* ID.

	   Match look id.
	*/
	ID *string

	/* Limit.

	   Number of results to return. (used with offset and takes priority over page and per_page)

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   Number of results to skip before returning any. (used with limit and takes priority over page and per_page)

	   Format: int64
	*/
	Offset *int64

	/* Page.

	   Requested page.

	   Format: int64
	*/
	Page *int64

	/* PerPage.

	   Results per page.

	   Format: int64
	*/
	PerPage *int64

	/* QueryID.

	   Select looks that reference a particular query by query_id

	   Format: int64
	*/
	QueryID *int64

	/* Sorts.

	   One or more fields to sort results by. Sortable fields: [:title, :user_id, :id, :created_at, :space_id, :folder_id, :description, :updated_at, :last_updater_id, :view_count, :favorite_count, :content_favorite_id, :deleted, :deleted_at, :last_viewed_at, :last_accessed_at, :query_id]
	*/
	Sorts *string

	/* SpaceID.

	   Select looks in a particular space.
	*/
	SpaceID *string

	/* Title.

	   Match Look title.
	*/
	Title *string

	/* UserID.

	   Select looks created by a particular user.
	*/
	UserID *string

	/* ViewCount.

	   Select looks with particular view_count value
	*/
	ViewCount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search looks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchLooksParams) WithDefaults() *SearchLooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search looks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchLooksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search looks params
func (o *SearchLooksParams) WithTimeout(timeout time.Duration) *SearchLooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search looks params
func (o *SearchLooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search looks params
func (o *SearchLooksParams) WithContext(ctx context.Context) *SearchLooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search looks params
func (o *SearchLooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search looks params
func (o *SearchLooksParams) WithHTTPClient(client *http.Client) *SearchLooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search looks params
func (o *SearchLooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentFavoriteID adds the contentFavoriteID to the search looks params
func (o *SearchLooksParams) WithContentFavoriteID(contentFavoriteID *int64) *SearchLooksParams {
	o.SetContentFavoriteID(contentFavoriteID)
	return o
}

// SetContentFavoriteID adds the contentFavoriteId to the search looks params
func (o *SearchLooksParams) SetContentFavoriteID(contentFavoriteID *int64) {
	o.ContentFavoriteID = contentFavoriteID
}

// WithCurate adds the curate to the search looks params
func (o *SearchLooksParams) WithCurate(curate *bool) *SearchLooksParams {
	o.SetCurate(curate)
	return o
}

// SetCurate adds the curate to the search looks params
func (o *SearchLooksParams) SetCurate(curate *bool) {
	o.Curate = curate
}

// WithDeleted adds the deleted to the search looks params
func (o *SearchLooksParams) WithDeleted(deleted *bool) *SearchLooksParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the search looks params
func (o *SearchLooksParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDescription adds the description to the search looks params
func (o *SearchLooksParams) WithDescription(description *string) *SearchLooksParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the search looks params
func (o *SearchLooksParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the search looks params
func (o *SearchLooksParams) WithFields(fields *string) *SearchLooksParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search looks params
func (o *SearchLooksParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilterOr adds the filterOr to the search looks params
func (o *SearchLooksParams) WithFilterOr(filterOr *bool) *SearchLooksParams {
	o.SetFilterOr(filterOr)
	return o
}

// SetFilterOr adds the filterOr to the search looks params
func (o *SearchLooksParams) SetFilterOr(filterOr *bool) {
	o.FilterOr = filterOr
}

// WithID adds the id to the search looks params
func (o *SearchLooksParams) WithID(id *string) *SearchLooksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search looks params
func (o *SearchLooksParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the search looks params
func (o *SearchLooksParams) WithLimit(limit *int64) *SearchLooksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search looks params
func (o *SearchLooksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search looks params
func (o *SearchLooksParams) WithOffset(offset *int64) *SearchLooksParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search looks params
func (o *SearchLooksParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the search looks params
func (o *SearchLooksParams) WithPage(page *int64) *SearchLooksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search looks params
func (o *SearchLooksParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the search looks params
func (o *SearchLooksParams) WithPerPage(perPage *int64) *SearchLooksParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the search looks params
func (o *SearchLooksParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithQueryID adds the queryID to the search looks params
func (o *SearchLooksParams) WithQueryID(queryID *int64) *SearchLooksParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the search looks params
func (o *SearchLooksParams) SetQueryID(queryID *int64) {
	o.QueryID = queryID
}

// WithSorts adds the sorts to the search looks params
func (o *SearchLooksParams) WithSorts(sorts *string) *SearchLooksParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search looks params
func (o *SearchLooksParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithSpaceID adds the spaceID to the search looks params
func (o *SearchLooksParams) WithSpaceID(spaceID *string) *SearchLooksParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the search looks params
func (o *SearchLooksParams) SetSpaceID(spaceID *string) {
	o.SpaceID = spaceID
}

// WithTitle adds the title to the search looks params
func (o *SearchLooksParams) WithTitle(title *string) *SearchLooksParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the search looks params
func (o *SearchLooksParams) SetTitle(title *string) {
	o.Title = title
}

// WithUserID adds the userID to the search looks params
func (o *SearchLooksParams) WithUserID(userID *string) *SearchLooksParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the search looks params
func (o *SearchLooksParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithViewCount adds the viewCount to the search looks params
func (o *SearchLooksParams) WithViewCount(viewCount *string) *SearchLooksParams {
	o.SetViewCount(viewCount)
	return o
}

// SetViewCount adds the viewCount to the search looks params
func (o *SearchLooksParams) SetViewCount(viewCount *string) {
	o.ViewCount = viewCount
}

// WriteToRequest writes these params to a swagger request
func (o *SearchLooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentFavoriteID != nil {

		// query param content_favorite_id
		var qrContentFavoriteID int64

		if o.ContentFavoriteID != nil {
			qrContentFavoriteID = *o.ContentFavoriteID
		}
		qContentFavoriteID := swag.FormatInt64(qrContentFavoriteID)
		if qContentFavoriteID != "" {

			if err := r.SetQueryParam("content_favorite_id", qContentFavoriteID); err != nil {
				return err
			}
		}
	}

	if o.Curate != nil {

		// query param curate
		var qrCurate bool

		if o.Curate != nil {
			qrCurate = *o.Curate
		}
		qCurate := swag.FormatBool(qrCurate)
		if qCurate != "" {

			if err := r.SetQueryParam("curate", qCurate); err != nil {
				return err
			}
		}
	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
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

	if o.FilterOr != nil {

		// query param filter_or
		var qrFilterOr bool

		if o.FilterOr != nil {
			qrFilterOr = *o.FilterOr
		}
		qFilterOr := swag.FormatBool(qrFilterOr)
		if qFilterOr != "" {

			if err := r.SetQueryParam("filter_or", qFilterOr); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.QueryID != nil {

		// query param query_id
		var qrQueryID int64

		if o.QueryID != nil {
			qrQueryID = *o.QueryID
		}
		qQueryID := swag.FormatInt64(qrQueryID)
		if qQueryID != "" {

			if err := r.SetQueryParam("query_id", qQueryID); err != nil {
				return err
			}
		}
	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string

		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {

			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}
	}

	if o.SpaceID != nil {

		// query param space_id
		var qrSpaceID string

		if o.SpaceID != nil {
			qrSpaceID = *o.SpaceID
		}
		qSpaceID := qrSpaceID
		if qSpaceID != "" {

			if err := r.SetQueryParam("space_id", qSpaceID); err != nil {
				return err
			}
		}
	}

	if o.Title != nil {

		// query param title
		var qrTitle string

		if o.Title != nil {
			qrTitle = *o.Title
		}
		qTitle := qrTitle
		if qTitle != "" {

			if err := r.SetQueryParam("title", qTitle); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.ViewCount != nil {

		// query param view_count
		var qrViewCount string

		if o.ViewCount != nil {
			qrViewCount = *o.ViewCount
		}
		qViewCount := qrViewCount
		if qViewCount != "" {

			if err := r.SetQueryParam("view_count", qViewCount); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
