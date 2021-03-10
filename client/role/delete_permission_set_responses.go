// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// DeletePermissionSetReader is a Reader for the DeletePermissionSet structure.
type DeletePermissionSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePermissionSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePermissionSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePermissionSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePermissionSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeletePermissionSetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePermissionSetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePermissionSetNoContent creates a DeletePermissionSetNoContent with default headers values
func NewDeletePermissionSetNoContent() *DeletePermissionSetNoContent {
	return &DeletePermissionSetNoContent{}
}

/* DeletePermissionSetNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeletePermissionSetNoContent struct {
	Payload string
}

func (o *DeletePermissionSetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /permission_sets/{permission_set_id}][%d] deletePermissionSetNoContent  %+v", 204, o.Payload)
}
func (o *DeletePermissionSetNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeletePermissionSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePermissionSetBadRequest creates a DeletePermissionSetBadRequest with default headers values
func NewDeletePermissionSetBadRequest() *DeletePermissionSetBadRequest {
	return &DeletePermissionSetBadRequest{}
}

/* DeletePermissionSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeletePermissionSetBadRequest struct {
	Payload *models.Error
}

func (o *DeletePermissionSetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /permission_sets/{permission_set_id}][%d] deletePermissionSetBadRequest  %+v", 400, o.Payload)
}
func (o *DeletePermissionSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePermissionSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePermissionSetNotFound creates a DeletePermissionSetNotFound with default headers values
func NewDeletePermissionSetNotFound() *DeletePermissionSetNotFound {
	return &DeletePermissionSetNotFound{}
}

/* DeletePermissionSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeletePermissionSetNotFound struct {
	Payload *models.Error
}

func (o *DeletePermissionSetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /permission_sets/{permission_set_id}][%d] deletePermissionSetNotFound  %+v", 404, o.Payload)
}
func (o *DeletePermissionSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePermissionSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePermissionSetMethodNotAllowed creates a DeletePermissionSetMethodNotAllowed with default headers values
func NewDeletePermissionSetMethodNotAllowed() *DeletePermissionSetMethodNotAllowed {
	return &DeletePermissionSetMethodNotAllowed{}
}

/* DeletePermissionSetMethodNotAllowed describes a response with status code 405, with default header values.

Resource Can't Be Modified
*/
type DeletePermissionSetMethodNotAllowed struct {
	Payload *models.Error
}

func (o *DeletePermissionSetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /permission_sets/{permission_set_id}][%d] deletePermissionSetMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *DeletePermissionSetMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePermissionSetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePermissionSetTooManyRequests creates a DeletePermissionSetTooManyRequests with default headers values
func NewDeletePermissionSetTooManyRequests() *DeletePermissionSetTooManyRequests {
	return &DeletePermissionSetTooManyRequests{}
}

/* DeletePermissionSetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeletePermissionSetTooManyRequests struct {
	Payload *models.Error
}

func (o *DeletePermissionSetTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /permission_sets/{permission_set_id}][%d] deletePermissionSetTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeletePermissionSetTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePermissionSetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
