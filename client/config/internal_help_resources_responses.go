// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// InternalHelpResourcesReader is a Reader for the InternalHelpResources structure.
type InternalHelpResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalHelpResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalHelpResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalHelpResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalHelpResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInternalHelpResourcesOK creates a InternalHelpResourcesOK with default headers values
func NewInternalHelpResourcesOK() *InternalHelpResourcesOK {
	return &InternalHelpResourcesOK{}
}

/* InternalHelpResourcesOK describes a response with status code 200, with default header values.

Internal Help Resources
*/
type InternalHelpResourcesOK struct {
	Payload *models.InternalHelpResources
}

func (o *InternalHelpResourcesOK) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_enabled][%d] internalHelpResourcesOK  %+v", 200, o.Payload)
}
func (o *InternalHelpResourcesOK) GetPayload() *models.InternalHelpResources {
	return o.Payload
}

func (o *InternalHelpResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalHelpResources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalHelpResourcesBadRequest creates a InternalHelpResourcesBadRequest with default headers values
func NewInternalHelpResourcesBadRequest() *InternalHelpResourcesBadRequest {
	return &InternalHelpResourcesBadRequest{}
}

/* InternalHelpResourcesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalHelpResourcesBadRequest struct {
	Payload *models.Error
}

func (o *InternalHelpResourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_enabled][%d] internalHelpResourcesBadRequest  %+v", 400, o.Payload)
}
func (o *InternalHelpResourcesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalHelpResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalHelpResourcesNotFound creates a InternalHelpResourcesNotFound with default headers values
func NewInternalHelpResourcesNotFound() *InternalHelpResourcesNotFound {
	return &InternalHelpResourcesNotFound{}
}

/* InternalHelpResourcesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalHelpResourcesNotFound struct {
	Payload *models.Error
}

func (o *InternalHelpResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_enabled][%d] internalHelpResourcesNotFound  %+v", 404, o.Payload)
}
func (o *InternalHelpResourcesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalHelpResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
