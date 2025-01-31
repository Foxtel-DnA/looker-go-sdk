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

// DashboardLayoutComponentReader is a Reader for the DashboardLayoutComponent structure.
type DashboardLayoutComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardLayoutComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardLayoutComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardLayoutComponentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardLayoutComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardLayoutComponentOK creates a DashboardLayoutComponentOK with default headers values
func NewDashboardLayoutComponentOK() *DashboardLayoutComponentOK {
	return &DashboardLayoutComponentOK{}
}

/* DashboardLayoutComponentOK describes a response with status code 200, with default header values.

DashboardLayoutComponent
*/
type DashboardLayoutComponentOK struct {
	Payload *models.DashboardLayoutComponent
}

func (o *DashboardLayoutComponentOK) Error() string {
	return fmt.Sprintf("[GET /dashboard_layout_components/{dashboard_layout_component_id}][%d] dashboardLayoutComponentOK  %+v", 200, o.Payload)
}
func (o *DashboardLayoutComponentOK) GetPayload() *models.DashboardLayoutComponent {
	return o.Payload
}

func (o *DashboardLayoutComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardLayoutComponent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardLayoutComponentBadRequest creates a DashboardLayoutComponentBadRequest with default headers values
func NewDashboardLayoutComponentBadRequest() *DashboardLayoutComponentBadRequest {
	return &DashboardLayoutComponentBadRequest{}
}

/* DashboardLayoutComponentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DashboardLayoutComponentBadRequest struct {
	Payload *models.Error
}

func (o *DashboardLayoutComponentBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboard_layout_components/{dashboard_layout_component_id}][%d] dashboardLayoutComponentBadRequest  %+v", 400, o.Payload)
}
func (o *DashboardLayoutComponentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardLayoutComponentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardLayoutComponentNotFound creates a DashboardLayoutComponentNotFound with default headers values
func NewDashboardLayoutComponentNotFound() *DashboardLayoutComponentNotFound {
	return &DashboardLayoutComponentNotFound{}
}

/* DashboardLayoutComponentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DashboardLayoutComponentNotFound struct {
	Payload *models.Error
}

func (o *DashboardLayoutComponentNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboard_layout_components/{dashboard_layout_component_id}][%d] dashboardLayoutComponentNotFound  %+v", 404, o.Payload)
}
func (o *DashboardLayoutComponentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardLayoutComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
