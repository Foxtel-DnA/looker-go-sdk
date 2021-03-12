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

// DashboardFilterReader is a Reader for the DashboardFilter structure.
type DashboardFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardFilterOK creates a DashboardFilterOK with default headers values
func NewDashboardFilterOK() *DashboardFilterOK {
	return &DashboardFilterOK{}
}

/* DashboardFilterOK describes a response with status code 200, with default header values.

Dashboard Filter
*/
type DashboardFilterOK struct {
	Payload *models.DashboardFilter
}

func (o *DashboardFilterOK) Error() string {
	return fmt.Sprintf("[GET /dashboard_filters/{dashboard_filter_id}][%d] dashboardFilterOK  %+v", 200, o.Payload)
}
func (o *DashboardFilterOK) GetPayload() *models.DashboardFilter {
	return o.Payload
}

func (o *DashboardFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardFilterBadRequest creates a DashboardFilterBadRequest with default headers values
func NewDashboardFilterBadRequest() *DashboardFilterBadRequest {
	return &DashboardFilterBadRequest{}
}

/* DashboardFilterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DashboardFilterBadRequest struct {
	Payload *models.Error
}

func (o *DashboardFilterBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboard_filters/{dashboard_filter_id}][%d] dashboardFilterBadRequest  %+v", 400, o.Payload)
}
func (o *DashboardFilterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardFilterNotFound creates a DashboardFilterNotFound with default headers values
func NewDashboardFilterNotFound() *DashboardFilterNotFound {
	return &DashboardFilterNotFound{}
}

/* DashboardFilterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DashboardFilterNotFound struct {
	Payload *models.Error
}

func (o *DashboardFilterNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboard_filters/{dashboard_filter_id}][%d] dashboardFilterNotFound  %+v", 404, o.Payload)
}
func (o *DashboardFilterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
