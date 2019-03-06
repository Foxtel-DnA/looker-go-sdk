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

// DeleteIntegrationHubReader is a Reader for the DeleteIntegrationHub structure.
type DeleteIntegrationHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteIntegrationHubNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteIntegrationHubBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteIntegrationHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIntegrationHubNoContent creates a DeleteIntegrationHubNoContent with default headers values
func NewDeleteIntegrationHubNoContent() *DeleteIntegrationHubNoContent {
	return &DeleteIntegrationHubNoContent{}
}

/*DeleteIntegrationHubNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteIntegrationHubNoContent struct {
	Payload string
}

func (o *DeleteIntegrationHubNoContent) Error() string {
	return fmt.Sprintf("[DELETE /integration_hubs/{integration_hub_id}][%d] deleteIntegrationHubNoContent  %+v", 204, o.Payload)
}

func (o *DeleteIntegrationHubNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationHubBadRequest creates a DeleteIntegrationHubBadRequest with default headers values
func NewDeleteIntegrationHubBadRequest() *DeleteIntegrationHubBadRequest {
	return &DeleteIntegrationHubBadRequest{}
}

/*DeleteIntegrationHubBadRequest handles this case with default header values.

Bad Request
*/
type DeleteIntegrationHubBadRequest struct {
	Payload *models.Error
}

func (o *DeleteIntegrationHubBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /integration_hubs/{integration_hub_id}][%d] deleteIntegrationHubBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationHubBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationHubNotFound creates a DeleteIntegrationHubNotFound with default headers values
func NewDeleteIntegrationHubNotFound() *DeleteIntegrationHubNotFound {
	return &DeleteIntegrationHubNotFound{}
}

/*DeleteIntegrationHubNotFound handles this case with default header values.

Not Found
*/
type DeleteIntegrationHubNotFound struct {
	Payload *models.Error
}

func (o *DeleteIntegrationHubNotFound) Error() string {
	return fmt.Sprintf("[DELETE /integration_hubs/{integration_hub_id}][%d] deleteIntegrationHubNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
