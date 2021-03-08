// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UpdateDigestEmailsEnabledReader is a Reader for the UpdateDigestEmailsEnabled structure.
type UpdateDigestEmailsEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDigestEmailsEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDigestEmailsEnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDigestEmailsEnabledBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDigestEmailsEnabledNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDigestEmailsEnabledUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDigestEmailsEnabledTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDigestEmailsEnabledOK creates a UpdateDigestEmailsEnabledOK with default headers values
func NewUpdateDigestEmailsEnabledOK() *UpdateDigestEmailsEnabledOK {
	return &UpdateDigestEmailsEnabledOK{}
}

/* UpdateDigestEmailsEnabledOK describes a response with status code 200, with default header values.

Digest_emails
*/
type UpdateDigestEmailsEnabledOK struct {
	Payload *models.DigestEmails
}

func (o *UpdateDigestEmailsEnabledOK) Error() string {
	return fmt.Sprintf("[PATCH /digest_emails_enabled][%d] updateDigestEmailsEnabledOK  %+v", 200, o.Payload)
}
func (o *UpdateDigestEmailsEnabledOK) GetPayload() *models.DigestEmails {
	return o.Payload
}

func (o *UpdateDigestEmailsEnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigestEmails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDigestEmailsEnabledBadRequest creates a UpdateDigestEmailsEnabledBadRequest with default headers values
func NewUpdateDigestEmailsEnabledBadRequest() *UpdateDigestEmailsEnabledBadRequest {
	return &UpdateDigestEmailsEnabledBadRequest{}
}

/* UpdateDigestEmailsEnabledBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateDigestEmailsEnabledBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDigestEmailsEnabledBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /digest_emails_enabled][%d] updateDigestEmailsEnabledBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDigestEmailsEnabledBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDigestEmailsEnabledBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDigestEmailsEnabledNotFound creates a UpdateDigestEmailsEnabledNotFound with default headers values
func NewUpdateDigestEmailsEnabledNotFound() *UpdateDigestEmailsEnabledNotFound {
	return &UpdateDigestEmailsEnabledNotFound{}
}

/* UpdateDigestEmailsEnabledNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateDigestEmailsEnabledNotFound struct {
	Payload *models.Error
}

func (o *UpdateDigestEmailsEnabledNotFound) Error() string {
	return fmt.Sprintf("[PATCH /digest_emails_enabled][%d] updateDigestEmailsEnabledNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDigestEmailsEnabledNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDigestEmailsEnabledNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDigestEmailsEnabledUnprocessableEntity creates a UpdateDigestEmailsEnabledUnprocessableEntity with default headers values
func NewUpdateDigestEmailsEnabledUnprocessableEntity() *UpdateDigestEmailsEnabledUnprocessableEntity {
	return &UpdateDigestEmailsEnabledUnprocessableEntity{}
}

/* UpdateDigestEmailsEnabledUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateDigestEmailsEnabledUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateDigestEmailsEnabledUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /digest_emails_enabled][%d] updateDigestEmailsEnabledUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateDigestEmailsEnabledUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateDigestEmailsEnabledUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDigestEmailsEnabledTooManyRequests creates a UpdateDigestEmailsEnabledTooManyRequests with default headers values
func NewUpdateDigestEmailsEnabledTooManyRequests() *UpdateDigestEmailsEnabledTooManyRequests {
	return &UpdateDigestEmailsEnabledTooManyRequests{}
}

/* UpdateDigestEmailsEnabledTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDigestEmailsEnabledTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateDigestEmailsEnabledTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /digest_emails_enabled][%d] updateDigestEmailsEnabledTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateDigestEmailsEnabledTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDigestEmailsEnabledTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
