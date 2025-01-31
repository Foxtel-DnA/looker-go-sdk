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

// DashboardAggregateTableLookmlReader is a Reader for the DashboardAggregateTableLookml structure.
type DashboardAggregateTableLookmlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardAggregateTableLookmlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardAggregateTableLookmlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardAggregateTableLookmlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardAggregateTableLookmlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardAggregateTableLookmlOK creates a DashboardAggregateTableLookmlOK with default headers values
func NewDashboardAggregateTableLookmlOK() *DashboardAggregateTableLookmlOK {
	return &DashboardAggregateTableLookmlOK{}
}

/* DashboardAggregateTableLookmlOK describes a response with status code 200, with default header values.

JSON for Aggregate Table LookML
*/
type DashboardAggregateTableLookmlOK struct {
	Payload *models.DashboardAggregateTableLookml
}

func (o *DashboardAggregateTableLookmlOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/aggregate_table_lookml/{dashboard_id}][%d] dashboardAggregateTableLookmlOK  %+v", 200, o.Payload)
}
func (o *DashboardAggregateTableLookmlOK) GetPayload() *models.DashboardAggregateTableLookml {
	return o.Payload
}

func (o *DashboardAggregateTableLookmlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardAggregateTableLookml)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardAggregateTableLookmlBadRequest creates a DashboardAggregateTableLookmlBadRequest with default headers values
func NewDashboardAggregateTableLookmlBadRequest() *DashboardAggregateTableLookmlBadRequest {
	return &DashboardAggregateTableLookmlBadRequest{}
}

/* DashboardAggregateTableLookmlBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DashboardAggregateTableLookmlBadRequest struct {
	Payload *models.Error
}

func (o *DashboardAggregateTableLookmlBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboards/aggregate_table_lookml/{dashboard_id}][%d] dashboardAggregateTableLookmlBadRequest  %+v", 400, o.Payload)
}
func (o *DashboardAggregateTableLookmlBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardAggregateTableLookmlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardAggregateTableLookmlNotFound creates a DashboardAggregateTableLookmlNotFound with default headers values
func NewDashboardAggregateTableLookmlNotFound() *DashboardAggregateTableLookmlNotFound {
	return &DashboardAggregateTableLookmlNotFound{}
}

/* DashboardAggregateTableLookmlNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DashboardAggregateTableLookmlNotFound struct {
	Payload *models.Error
}

func (o *DashboardAggregateTableLookmlNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/aggregate_table_lookml/{dashboard_id}][%d] dashboardAggregateTableLookmlNotFound  %+v", 404, o.Payload)
}
func (o *DashboardAggregateTableLookmlNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardAggregateTableLookmlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
