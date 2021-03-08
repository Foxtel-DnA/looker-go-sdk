// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// ContentValidationReader is a Reader for the ContentValidation structure.
type ContentValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContentValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContentValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContentValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContentValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewContentValidationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewContentValidationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContentValidationOK creates a ContentValidationOK with default headers values
func NewContentValidationOK() *ContentValidationOK {
	return &ContentValidationOK{}
}

/* ContentValidationOK describes a response with status code 200, with default header values.

Content validation results
*/
type ContentValidationOK struct {
	Payload *models.ContentValidation
}

func (o *ContentValidationOK) Error() string {
	return fmt.Sprintf("[GET /content_validation][%d] contentValidationOK  %+v", 200, o.Payload)
}
func (o *ContentValidationOK) GetPayload() *models.ContentValidation {
	return o.Payload
}

func (o *ContentValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentValidationBadRequest creates a ContentValidationBadRequest with default headers values
func NewContentValidationBadRequest() *ContentValidationBadRequest {
	return &ContentValidationBadRequest{}
}

/* ContentValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ContentValidationBadRequest struct {
	Payload *models.Error
}

func (o *ContentValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_validation][%d] contentValidationBadRequest  %+v", 400, o.Payload)
}
func (o *ContentValidationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentValidationNotFound creates a ContentValidationNotFound with default headers values
func NewContentValidationNotFound() *ContentValidationNotFound {
	return &ContentValidationNotFound{}
}

/* ContentValidationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ContentValidationNotFound struct {
	Payload *models.Error
}

func (o *ContentValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /content_validation][%d] contentValidationNotFound  %+v", 404, o.Payload)
}
func (o *ContentValidationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentValidationUnprocessableEntity creates a ContentValidationUnprocessableEntity with default headers values
func NewContentValidationUnprocessableEntity() *ContentValidationUnprocessableEntity {
	return &ContentValidationUnprocessableEntity{}
}

/* ContentValidationUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ContentValidationUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ContentValidationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /content_validation][%d] contentValidationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ContentValidationUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ContentValidationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentValidationTooManyRequests creates a ContentValidationTooManyRequests with default headers values
func NewContentValidationTooManyRequests() *ContentValidationTooManyRequests {
	return &ContentValidationTooManyRequests{}
}

/* ContentValidationTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ContentValidationTooManyRequests struct {
	Payload *models.Error
}

func (o *ContentValidationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /content_validation][%d] contentValidationTooManyRequests  %+v", 429, o.Payload)
}
func (o *ContentValidationTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentValidationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}