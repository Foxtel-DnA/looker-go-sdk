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

// SearchDashboardElementsReader is a Reader for the SearchDashboardElements structure.
type SearchDashboardElementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDashboardElementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchDashboardElementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchDashboardElementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchDashboardElementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchDashboardElementsOK creates a SearchDashboardElementsOK with default headers values
func NewSearchDashboardElementsOK() *SearchDashboardElementsOK {
	return &SearchDashboardElementsOK{}
}

/* SearchDashboardElementsOK describes a response with status code 200, with default header values.

Dashboard elements
*/
type SearchDashboardElementsOK struct {
	Payload []*models.DashboardElement
}

func (o *SearchDashboardElementsOK) Error() string {
	return fmt.Sprintf("[GET /dashboard_elements/search][%d] searchDashboardElementsOK  %+v", 200, o.Payload)
}
func (o *SearchDashboardElementsOK) GetPayload() []*models.DashboardElement {
	return o.Payload
}

func (o *SearchDashboardElementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDashboardElementsBadRequest creates a SearchDashboardElementsBadRequest with default headers values
func NewSearchDashboardElementsBadRequest() *SearchDashboardElementsBadRequest {
	return &SearchDashboardElementsBadRequest{}
}

/* SearchDashboardElementsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchDashboardElementsBadRequest struct {
	Payload *models.Error
}

func (o *SearchDashboardElementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboard_elements/search][%d] searchDashboardElementsBadRequest  %+v", 400, o.Payload)
}
func (o *SearchDashboardElementsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchDashboardElementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDashboardElementsNotFound creates a SearchDashboardElementsNotFound with default headers values
func NewSearchDashboardElementsNotFound() *SearchDashboardElementsNotFound {
	return &SearchDashboardElementsNotFound{}
}

/* SearchDashboardElementsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchDashboardElementsNotFound struct {
	Payload *models.Error
}

func (o *SearchDashboardElementsNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboard_elements/search][%d] searchDashboardElementsNotFound  %+v", 404, o.Payload)
}
func (o *SearchDashboardElementsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchDashboardElementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
