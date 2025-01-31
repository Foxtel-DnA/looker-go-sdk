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
)

// NewSearchHomepagesParams creates a new SearchHomepagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchHomepagesParams() *SearchHomepagesParams {
	return &SearchHomepagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchHomepagesParamsWithTimeout creates a new SearchHomepagesParams object
// with the ability to set a timeout on a request.
func NewSearchHomepagesParamsWithTimeout(timeout time.Duration) *SearchHomepagesParams {
	return &SearchHomepagesParams{
		timeout: timeout,
	}
}

// NewSearchHomepagesParamsWithContext creates a new SearchHomepagesParams object
// with the ability to set a context for a request.
func NewSearchHomepagesParamsWithContext(ctx context.Context) *SearchHomepagesParams {
	return &SearchHomepagesParams{
		Context: ctx,
	}
}

// NewSearchHomepagesParamsWithHTTPClient creates a new SearchHomepagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchHomepagesParamsWithHTTPClient(client *http.Client) *SearchHomepagesParams {
	return &SearchHomepagesParams{
		HTTPClient: client,
	}
}

/* SearchHomepagesParams contains all the parameters to send to the API endpoint
   for the search homepages operation.

   Typically these are written to a http.Request.
*/
type SearchHomepagesParams struct {

	/* CreatedAt.

	   Matches the timestamp for when the homepage was created.
	*/
	CreatedAt *string

	/* CreatorID.

	   Filter on homepages created by a particular user.
	*/
	CreatorID *string

	/* Favorited.

	   Return favorited homepages when true.
	*/
	Favorited *bool

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* FilterOr.

	   Combine given search criteria in a boolean OR expression
	*/
	FilterOr *bool

	/* FirstName.

	   The first name of the user who created this homepage.
	*/
	FirstName *string

	/* LastName.

	   The last name of the user who created this homepage.
	*/
	LastName *string

	/* Limit.

	   The maximum number of items to return. (used with offset and takes priority over page and per_page)

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   The number of items to skip before returning any. (used with limit and takes priority over page and per_page)

	   Format: int64
	*/
	Offset *int64

	/* Page.

	   The page to return.

	   Format: int64
	*/
	Page *int64

	/* PerPage.

	   The number of items in the returned page.

	   Format: int64
	*/
	PerPage *int64

	/* Sorts.

	   The fields to sort the results by.
	*/
	Sorts *string

	/* Title.

	   Matches homepage title.
	*/
	Title *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search homepages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchHomepagesParams) WithDefaults() *SearchHomepagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search homepages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchHomepagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search homepages params
func (o *SearchHomepagesParams) WithTimeout(timeout time.Duration) *SearchHomepagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search homepages params
func (o *SearchHomepagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search homepages params
func (o *SearchHomepagesParams) WithContext(ctx context.Context) *SearchHomepagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search homepages params
func (o *SearchHomepagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search homepages params
func (o *SearchHomepagesParams) WithHTTPClient(client *http.Client) *SearchHomepagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search homepages params
func (o *SearchHomepagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAt adds the createdAt to the search homepages params
func (o *SearchHomepagesParams) WithCreatedAt(createdAt *string) *SearchHomepagesParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the search homepages params
func (o *SearchHomepagesParams) SetCreatedAt(createdAt *string) {
	o.CreatedAt = createdAt
}

// WithCreatorID adds the creatorID to the search homepages params
func (o *SearchHomepagesParams) WithCreatorID(creatorID *string) *SearchHomepagesParams {
	o.SetCreatorID(creatorID)
	return o
}

// SetCreatorID adds the creatorId to the search homepages params
func (o *SearchHomepagesParams) SetCreatorID(creatorID *string) {
	o.CreatorID = creatorID
}

// WithFavorited adds the favorited to the search homepages params
func (o *SearchHomepagesParams) WithFavorited(favorited *bool) *SearchHomepagesParams {
	o.SetFavorited(favorited)
	return o
}

// SetFavorited adds the favorited to the search homepages params
func (o *SearchHomepagesParams) SetFavorited(favorited *bool) {
	o.Favorited = favorited
}

// WithFields adds the fields to the search homepages params
func (o *SearchHomepagesParams) WithFields(fields *string) *SearchHomepagesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search homepages params
func (o *SearchHomepagesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilterOr adds the filterOr to the search homepages params
func (o *SearchHomepagesParams) WithFilterOr(filterOr *bool) *SearchHomepagesParams {
	o.SetFilterOr(filterOr)
	return o
}

// SetFilterOr adds the filterOr to the search homepages params
func (o *SearchHomepagesParams) SetFilterOr(filterOr *bool) {
	o.FilterOr = filterOr
}

// WithFirstName adds the firstName to the search homepages params
func (o *SearchHomepagesParams) WithFirstName(firstName *string) *SearchHomepagesParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the search homepages params
func (o *SearchHomepagesParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithLastName adds the lastName to the search homepages params
func (o *SearchHomepagesParams) WithLastName(lastName *string) *SearchHomepagesParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the search homepages params
func (o *SearchHomepagesParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithLimit adds the limit to the search homepages params
func (o *SearchHomepagesParams) WithLimit(limit *int64) *SearchHomepagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search homepages params
func (o *SearchHomepagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search homepages params
func (o *SearchHomepagesParams) WithOffset(offset *int64) *SearchHomepagesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search homepages params
func (o *SearchHomepagesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the search homepages params
func (o *SearchHomepagesParams) WithPage(page *int64) *SearchHomepagesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search homepages params
func (o *SearchHomepagesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the search homepages params
func (o *SearchHomepagesParams) WithPerPage(perPage *int64) *SearchHomepagesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the search homepages params
func (o *SearchHomepagesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSorts adds the sorts to the search homepages params
func (o *SearchHomepagesParams) WithSorts(sorts *string) *SearchHomepagesParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search homepages params
func (o *SearchHomepagesParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithTitle adds the title to the search homepages params
func (o *SearchHomepagesParams) WithTitle(title *string) *SearchHomepagesParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the search homepages params
func (o *SearchHomepagesParams) SetTitle(title *string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *SearchHomepagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedAt != nil {

		// query param created_at
		var qrCreatedAt string

		if o.CreatedAt != nil {
			qrCreatedAt = *o.CreatedAt
		}
		qCreatedAt := qrCreatedAt
		if qCreatedAt != "" {

			if err := r.SetQueryParam("created_at", qCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.CreatorID != nil {

		// query param creator_id
		var qrCreatorID string

		if o.CreatorID != nil {
			qrCreatorID = *o.CreatorID
		}
		qCreatorID := qrCreatorID
		if qCreatorID != "" {

			if err := r.SetQueryParam("creator_id", qCreatorID); err != nil {
				return err
			}
		}
	}

	if o.Favorited != nil {

		// query param favorited
		var qrFavorited bool

		if o.Favorited != nil {
			qrFavorited = *o.Favorited
		}
		qFavorited := swag.FormatBool(qrFavorited)
		if qFavorited != "" {

			if err := r.SetQueryParam("favorited", qFavorited); err != nil {
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

	if o.FirstName != nil {

		// query param first_name
		var qrFirstName string

		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {

			if err := r.SetQueryParam("first_name", qFirstName); err != nil {
				return err
			}
		}
	}

	if o.LastName != nil {

		// query param last_name
		var qrLastName string

		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {

			if err := r.SetQueryParam("last_name", qLastName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
