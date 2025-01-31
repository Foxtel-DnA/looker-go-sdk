// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// DashboardDashboardFiltersReader is a Reader for the DashboardDashboardFilters structure.
type DashboardDashboardFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardDashboardFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardDashboardFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardDashboardFiltersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardDashboardFiltersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardDashboardFiltersOK creates a DashboardDashboardFiltersOK with default headers values
func NewDashboardDashboardFiltersOK() *DashboardDashboardFiltersOK {
	return &DashboardDashboardFiltersOK{}
}

/* DashboardDashboardFiltersOK describes a response with status code 200, with default header values.

Dashboard Filter
*/
type DashboardDashboardFiltersOK struct {
	Payload []*models.DashboardFilter
}

func (o *DashboardDashboardFiltersOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_filters][%d] dashboardDashboardFiltersOK  %+v", 200, o.Payload)
}
func (o *DashboardDashboardFiltersOK) GetPayload() []*models.DashboardFilter {
	return o.Payload
}

func (o *DashboardDashboardFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardFiltersBadRequest creates a DashboardDashboardFiltersBadRequest with default headers values
func NewDashboardDashboardFiltersBadRequest() *DashboardDashboardFiltersBadRequest {
	return &DashboardDashboardFiltersBadRequest{}
}

/* DashboardDashboardFiltersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DashboardDashboardFiltersBadRequest struct {
	Payload *models.Error
}

func (o *DashboardDashboardFiltersBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_filters][%d] dashboardDashboardFiltersBadRequest  %+v", 400, o.Payload)
}
func (o *DashboardDashboardFiltersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardDashboardFiltersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardFiltersNotFound creates a DashboardDashboardFiltersNotFound with default headers values
func NewDashboardDashboardFiltersNotFound() *DashboardDashboardFiltersNotFound {
	return &DashboardDashboardFiltersNotFound{}
}

/* DashboardDashboardFiltersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DashboardDashboardFiltersNotFound struct {
	Payload *models.Error
}

func (o *DashboardDashboardFiltersNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_filters][%d] dashboardDashboardFiltersNotFound  %+v", 404, o.Payload)
}
func (o *DashboardDashboardFiltersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardDashboardFiltersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
