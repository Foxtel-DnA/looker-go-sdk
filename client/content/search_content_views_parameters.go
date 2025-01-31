// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewSearchContentViewsParams creates a new SearchContentViewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchContentViewsParams() *SearchContentViewsParams {
	return &SearchContentViewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchContentViewsParamsWithTimeout creates a new SearchContentViewsParams object
// with the ability to set a timeout on a request.
func NewSearchContentViewsParamsWithTimeout(timeout time.Duration) *SearchContentViewsParams {
	return &SearchContentViewsParams{
		timeout: timeout,
	}
}

// NewSearchContentViewsParamsWithContext creates a new SearchContentViewsParams object
// with the ability to set a context for a request.
func NewSearchContentViewsParamsWithContext(ctx context.Context) *SearchContentViewsParams {
	return &SearchContentViewsParams{
		Context: ctx,
	}
}

// NewSearchContentViewsParamsWithHTTPClient creates a new SearchContentViewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchContentViewsParamsWithHTTPClient(client *http.Client) *SearchContentViewsParams {
	return &SearchContentViewsParams{
		HTTPClient: client,
	}
}

/* SearchContentViewsParams contains all the parameters to send to the API endpoint
   for the search content views operation.

   Typically these are written to a http.Request.
*/
type SearchContentViewsParams struct {

	/* AllTime.

	   True if only all time view records should be returned
	*/
	AllTime *bool

	/* ContentMetadataID.

	   Match content metadata id

	   Format: int64
	*/
	ContentMetadataID *int64

	/* DashboardID.

	   Match dashboard_id
	*/
	DashboardID *string

	/* Fields.

	   Requested fields
	*/
	Fields *string

	/* FilterOr.

	   Combine given search criteria in a boolean OR expression
	*/
	FilterOr *bool

	/* GroupID.

	   Match Group Id

	   Format: int64
	*/
	GroupID *int64

	/* Limit.

	   Number of results to return. Use with `offset` to manage pagination of results

	   Format: int64
	*/
	Limit *int64

	/* LookID.

	   Match look_id
	*/
	LookID *string

	/* Offset.

	   Number of results to skip before returning data

	   Format: int64
	*/
	Offset *int64

	/* Sorts.

	   Fields to sort by
	*/
	Sorts *string

	/* StartOfWeekDate.

	   Match start of week date (format is "YYYY-MM-DD")
	*/
	StartOfWeekDate *string

	/* UserID.

	   Match user id

	   Format: int64
	*/
	UserID *int64

	/* ViewCount.

	   Match view count

	   Format: int64
	*/
	ViewCount *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search content views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchContentViewsParams) WithDefaults() *SearchContentViewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search content views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchContentViewsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search content views params
func (o *SearchContentViewsParams) WithTimeout(timeout time.Duration) *SearchContentViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search content views params
func (o *SearchContentViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search content views params
func (o *SearchContentViewsParams) WithContext(ctx context.Context) *SearchContentViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search content views params
func (o *SearchContentViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search content views params
func (o *SearchContentViewsParams) WithHTTPClient(client *http.Client) *SearchContentViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search content views params
func (o *SearchContentViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllTime adds the allTime to the search content views params
func (o *SearchContentViewsParams) WithAllTime(allTime *bool) *SearchContentViewsParams {
	o.SetAllTime(allTime)
	return o
}

// SetAllTime adds the allTime to the search content views params
func (o *SearchContentViewsParams) SetAllTime(allTime *bool) {
	o.AllTime = allTime
}

// WithContentMetadataID adds the contentMetadataID to the search content views params
func (o *SearchContentViewsParams) WithContentMetadataID(contentMetadataID *int64) *SearchContentViewsParams {
	o.SetContentMetadataID(contentMetadataID)
	return o
}

// SetContentMetadataID adds the contentMetadataId to the search content views params
func (o *SearchContentViewsParams) SetContentMetadataID(contentMetadataID *int64) {
	o.ContentMetadataID = contentMetadataID
}

// WithDashboardID adds the dashboardID to the search content views params
func (o *SearchContentViewsParams) WithDashboardID(dashboardID *string) *SearchContentViewsParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the search content views params
func (o *SearchContentViewsParams) SetDashboardID(dashboardID *string) {
	o.DashboardID = dashboardID
}

// WithFields adds the fields to the search content views params
func (o *SearchContentViewsParams) WithFields(fields *string) *SearchContentViewsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search content views params
func (o *SearchContentViewsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilterOr adds the filterOr to the search content views params
func (o *SearchContentViewsParams) WithFilterOr(filterOr *bool) *SearchContentViewsParams {
	o.SetFilterOr(filterOr)
	return o
}

// SetFilterOr adds the filterOr to the search content views params
func (o *SearchContentViewsParams) SetFilterOr(filterOr *bool) {
	o.FilterOr = filterOr
}

// WithGroupID adds the groupID to the search content views params
func (o *SearchContentViewsParams) WithGroupID(groupID *int64) *SearchContentViewsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the search content views params
func (o *SearchContentViewsParams) SetGroupID(groupID *int64) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the search content views params
func (o *SearchContentViewsParams) WithLimit(limit *int64) *SearchContentViewsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search content views params
func (o *SearchContentViewsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLookID adds the lookID to the search content views params
func (o *SearchContentViewsParams) WithLookID(lookID *string) *SearchContentViewsParams {
	o.SetLookID(lookID)
	return o
}

// SetLookID adds the lookId to the search content views params
func (o *SearchContentViewsParams) SetLookID(lookID *string) {
	o.LookID = lookID
}

// WithOffset adds the offset to the search content views params
func (o *SearchContentViewsParams) WithOffset(offset *int64) *SearchContentViewsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search content views params
func (o *SearchContentViewsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSorts adds the sorts to the search content views params
func (o *SearchContentViewsParams) WithSorts(sorts *string) *SearchContentViewsParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search content views params
func (o *SearchContentViewsParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithStartOfWeekDate adds the startOfWeekDate to the search content views params
func (o *SearchContentViewsParams) WithStartOfWeekDate(startOfWeekDate *string) *SearchContentViewsParams {
	o.SetStartOfWeekDate(startOfWeekDate)
	return o
}

// SetStartOfWeekDate adds the startOfWeekDate to the search content views params
func (o *SearchContentViewsParams) SetStartOfWeekDate(startOfWeekDate *string) {
	o.StartOfWeekDate = startOfWeekDate
}

// WithUserID adds the userID to the search content views params
func (o *SearchContentViewsParams) WithUserID(userID *int64) *SearchContentViewsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the search content views params
func (o *SearchContentViewsParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WithViewCount adds the viewCount to the search content views params
func (o *SearchContentViewsParams) WithViewCount(viewCount *int64) *SearchContentViewsParams {
	o.SetViewCount(viewCount)
	return o
}

// SetViewCount adds the viewCount to the search content views params
func (o *SearchContentViewsParams) SetViewCount(viewCount *int64) {
	o.ViewCount = viewCount
}

// WriteToRequest writes these params to a swagger request
func (o *SearchContentViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllTime != nil {

		// query param all_time
		var qrAllTime bool

		if o.AllTime != nil {
			qrAllTime = *o.AllTime
		}
		qAllTime := swag.FormatBool(qrAllTime)
		if qAllTime != "" {

			if err := r.SetQueryParam("all_time", qAllTime); err != nil {
				return err
			}
		}
	}

	if o.ContentMetadataID != nil {

		// query param content_metadata_id
		var qrContentMetadataID int64

		if o.ContentMetadataID != nil {
			qrContentMetadataID = *o.ContentMetadataID
		}
		qContentMetadataID := swag.FormatInt64(qrContentMetadataID)
		if qContentMetadataID != "" {

			if err := r.SetQueryParam("content_metadata_id", qContentMetadataID); err != nil {
				return err
			}
		}
	}

	if o.DashboardID != nil {

		// query param dashboard_id
		var qrDashboardID string

		if o.DashboardID != nil {
			qrDashboardID = *o.DashboardID
		}
		qDashboardID := qrDashboardID
		if qDashboardID != "" {

			if err := r.SetQueryParam("dashboard_id", qDashboardID); err != nil {
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

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID int64

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := swag.FormatInt64(qrGroupID)
		if qGroupID != "" {

			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
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

	if o.LookID != nil {

		// query param look_id
		var qrLookID string

		if o.LookID != nil {
			qrLookID = *o.LookID
		}
		qLookID := qrLookID
		if qLookID != "" {

			if err := r.SetQueryParam("look_id", qLookID); err != nil {
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

	if o.StartOfWeekDate != nil {

		// query param start_of_week_date
		var qrStartOfWeekDate string

		if o.StartOfWeekDate != nil {
			qrStartOfWeekDate = *o.StartOfWeekDate
		}
		qStartOfWeekDate := qrStartOfWeekDate
		if qStartOfWeekDate != "" {

			if err := r.SetQueryParam("start_of_week_date", qStartOfWeekDate); err != nil {
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

	if o.ViewCount != nil {

		// query param view_count
		var qrViewCount int64

		if o.ViewCount != nil {
			qrViewCount = *o.ViewCount
		}
		qViewCount := swag.FormatInt64(qrViewCount)
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
