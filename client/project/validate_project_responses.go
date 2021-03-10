// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// ValidateProjectReader is a Reader for the ValidateProject structure.
type ValidateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewValidateProjectTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateProjectOK creates a ValidateProjectOK with default headers values
func NewValidateProjectOK() *ValidateProjectOK {
	return &ValidateProjectOK{}
}

/* ValidateProjectOK describes a response with status code 200, with default header values.

Project validation results
*/
type ValidateProjectOK struct {
	Payload *models.ProjectValidation
}

func (o *ValidateProjectOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/validate][%d] validateProjectOK  %+v", 200, o.Payload)
}
func (o *ValidateProjectOK) GetPayload() *models.ProjectValidation {
	return o.Payload
}

func (o *ValidateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateProjectBadRequest creates a ValidateProjectBadRequest with default headers values
func NewValidateProjectBadRequest() *ValidateProjectBadRequest {
	return &ValidateProjectBadRequest{}
}

/* ValidateProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateProjectBadRequest struct {
	Payload *models.Error
}

func (o *ValidateProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/validate][%d] validateProjectBadRequest  %+v", 400, o.Payload)
}
func (o *ValidateProjectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateProjectNotFound creates a ValidateProjectNotFound with default headers values
func NewValidateProjectNotFound() *ValidateProjectNotFound {
	return &ValidateProjectNotFound{}
}

/* ValidateProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateProjectNotFound struct {
	Payload *models.Error
}

func (o *ValidateProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/validate][%d] validateProjectNotFound  %+v", 404, o.Payload)
}
func (o *ValidateProjectNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateProjectUnprocessableEntity creates a ValidateProjectUnprocessableEntity with default headers values
func NewValidateProjectUnprocessableEntity() *ValidateProjectUnprocessableEntity {
	return &ValidateProjectUnprocessableEntity{}
}

/* ValidateProjectUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ValidateProjectUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ValidateProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/validate][%d] validateProjectUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ValidateProjectUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ValidateProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateProjectTooManyRequests creates a ValidateProjectTooManyRequests with default headers values
func NewValidateProjectTooManyRequests() *ValidateProjectTooManyRequests {
	return &ValidateProjectTooManyRequests{}
}

/* ValidateProjectTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ValidateProjectTooManyRequests struct {
	Payload *models.Error
}

func (o *ValidateProjectTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/validate][%d] validateProjectTooManyRequests  %+v", 429, o.Payload)
}
func (o *ValidateProjectTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateProjectTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
