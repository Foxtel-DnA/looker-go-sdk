// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UpdateLookReader is a Reader for the UpdateLook structure.
type UpdateLookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateLookUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateLookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLookOK creates a UpdateLookOK with default headers values
func NewUpdateLookOK() *UpdateLookOK {
	return &UpdateLookOK{}
}

/* UpdateLookOK describes a response with status code 200, with default header values.

Look
*/
type UpdateLookOK struct {
	Payload *models.LookWithQuery
}

func (o *UpdateLookOK) Error() string {
	return fmt.Sprintf("[PATCH /looks/{look_id}][%d] updateLookOK  %+v", 200, o.Payload)
}
func (o *UpdateLookOK) GetPayload() *models.LookWithQuery {
	return o.Payload
}

func (o *UpdateLookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookWithQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookBadRequest creates a UpdateLookBadRequest with default headers values
func NewUpdateLookBadRequest() *UpdateLookBadRequest {
	return &UpdateLookBadRequest{}
}

/* UpdateLookBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateLookBadRequest struct {
	Payload *models.Error
}

func (o *UpdateLookBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /looks/{look_id}][%d] updateLookBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateLookBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookNotFound creates a UpdateLookNotFound with default headers values
func NewUpdateLookNotFound() *UpdateLookNotFound {
	return &UpdateLookNotFound{}
}

/* UpdateLookNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateLookNotFound struct {
	Payload *models.Error
}

func (o *UpdateLookNotFound) Error() string {
	return fmt.Sprintf("[PATCH /looks/{look_id}][%d] updateLookNotFound  %+v", 404, o.Payload)
}
func (o *UpdateLookNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookUnprocessableEntity creates a UpdateLookUnprocessableEntity with default headers values
func NewUpdateLookUnprocessableEntity() *UpdateLookUnprocessableEntity {
	return &UpdateLookUnprocessableEntity{}
}

/* UpdateLookUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateLookUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateLookUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /looks/{look_id}][%d] updateLookUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateLookUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateLookUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookTooManyRequests creates a UpdateLookTooManyRequests with default headers values
func NewUpdateLookTooManyRequests() *UpdateLookTooManyRequests {
	return &UpdateLookTooManyRequests{}
}

/* UpdateLookTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateLookTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateLookTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /looks/{look_id}][%d] updateLookTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateLookTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
