// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// IntegrationHubReader is a Reader for the IntegrationHub structure.
type IntegrationHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIntegrationHubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewIntegrationHubBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewIntegrationHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationHubOK creates a IntegrationHubOK with default headers values
func NewIntegrationHubOK() *IntegrationHubOK {
	return &IntegrationHubOK{}
}

/*IntegrationHubOK handles this case with default header values.

Integration Hub
*/
type IntegrationHubOK struct {
	Payload *models.IntegrationHub
}

func (o *IntegrationHubOK) Error() string {
	return fmt.Sprintf("[GET /integration_hubs/{integration_hub_id}][%d] integrationHubOK  %+v", 200, o.Payload)
}

func (o *IntegrationHubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationHub)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationHubBadRequest creates a IntegrationHubBadRequest with default headers values
func NewIntegrationHubBadRequest() *IntegrationHubBadRequest {
	return &IntegrationHubBadRequest{}
}

/*IntegrationHubBadRequest handles this case with default header values.

Bad Request
*/
type IntegrationHubBadRequest struct {
	Payload *models.Error
}

func (o *IntegrationHubBadRequest) Error() string {
	return fmt.Sprintf("[GET /integration_hubs/{integration_hub_id}][%d] integrationHubBadRequest  %+v", 400, o.Payload)
}

func (o *IntegrationHubBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationHubNotFound creates a IntegrationHubNotFound with default headers values
func NewIntegrationHubNotFound() *IntegrationHubNotFound {
	return &IntegrationHubNotFound{}
}

/*IntegrationHubNotFound handles this case with default header values.

Not Found
*/
type IntegrationHubNotFound struct {
	Payload *models.Error
}

func (o *IntegrationHubNotFound) Error() string {
	return fmt.Sprintf("[GET /integration_hubs/{integration_hub_id}][%d] integrationHubNotFound  %+v", 404, o.Payload)
}

func (o *IntegrationHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}