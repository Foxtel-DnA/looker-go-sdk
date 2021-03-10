// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// UpdateUserAttributeReader is a Reader for the UpdateUserAttribute structure.
type UpdateUserAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserAttributeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserAttributeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserAttributeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserAttributeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateUserAttributeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserAttributeOK creates a UpdateUserAttributeOK with default headers values
func NewUpdateUserAttributeOK() *UpdateUserAttributeOK {
	return &UpdateUserAttributeOK{}
}

/* UpdateUserAttributeOK describes a response with status code 200, with default header values.

User Attribute
*/
type UpdateUserAttributeOK struct {
	Payload *models.UserAttribute
}

func (o *UpdateUserAttributeOK) Error() string {
	return fmt.Sprintf("[PATCH /user_attributes/{user_attribute_id}][%d] updateUserAttributeOK  %+v", 200, o.Payload)
}
func (o *UpdateUserAttributeOK) GetPayload() *models.UserAttribute {
	return o.Payload
}

func (o *UpdateUserAttributeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAttributeBadRequest creates a UpdateUserAttributeBadRequest with default headers values
func NewUpdateUserAttributeBadRequest() *UpdateUserAttributeBadRequest {
	return &UpdateUserAttributeBadRequest{}
}

/* UpdateUserAttributeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateUserAttributeBadRequest struct {
	Payload *models.Error
}

func (o *UpdateUserAttributeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /user_attributes/{user_attribute_id}][%d] updateUserAttributeBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateUserAttributeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserAttributeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAttributeNotFound creates a UpdateUserAttributeNotFound with default headers values
func NewUpdateUserAttributeNotFound() *UpdateUserAttributeNotFound {
	return &UpdateUserAttributeNotFound{}
}

/* UpdateUserAttributeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUserAttributeNotFound struct {
	Payload *models.Error
}

func (o *UpdateUserAttributeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /user_attributes/{user_attribute_id}][%d] updateUserAttributeNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUserAttributeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserAttributeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAttributeUnprocessableEntity creates a UpdateUserAttributeUnprocessableEntity with default headers values
func NewUpdateUserAttributeUnprocessableEntity() *UpdateUserAttributeUnprocessableEntity {
	return &UpdateUserAttributeUnprocessableEntity{}
}

/* UpdateUserAttributeUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateUserAttributeUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateUserAttributeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /user_attributes/{user_attribute_id}][%d] updateUserAttributeUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateUserAttributeUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateUserAttributeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAttributeTooManyRequests creates a UpdateUserAttributeTooManyRequests with default headers values
func NewUpdateUserAttributeTooManyRequests() *UpdateUserAttributeTooManyRequests {
	return &UpdateUserAttributeTooManyRequests{}
}

/* UpdateUserAttributeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateUserAttributeTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateUserAttributeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /user_attributes/{user_attribute_id}][%d] updateUserAttributeTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateUserAttributeTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserAttributeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
