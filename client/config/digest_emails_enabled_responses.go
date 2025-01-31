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

// DigestEmailsEnabledReader is a Reader for the DigestEmailsEnabled structure.
type DigestEmailsEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DigestEmailsEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDigestEmailsEnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDigestEmailsEnabledBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDigestEmailsEnabledNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDigestEmailsEnabledOK creates a DigestEmailsEnabledOK with default headers values
func NewDigestEmailsEnabledOK() *DigestEmailsEnabledOK {
	return &DigestEmailsEnabledOK{}
}

/* DigestEmailsEnabledOK describes a response with status code 200, with default header values.

Digest_emails
*/
type DigestEmailsEnabledOK struct {
	Payload *models.DigestEmails
}

func (o *DigestEmailsEnabledOK) Error() string {
	return fmt.Sprintf("[GET /digest_emails_enabled][%d] digestEmailsEnabledOK  %+v", 200, o.Payload)
}
func (o *DigestEmailsEnabledOK) GetPayload() *models.DigestEmails {
	return o.Payload
}

func (o *DigestEmailsEnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigestEmails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDigestEmailsEnabledBadRequest creates a DigestEmailsEnabledBadRequest with default headers values
func NewDigestEmailsEnabledBadRequest() *DigestEmailsEnabledBadRequest {
	return &DigestEmailsEnabledBadRequest{}
}

/* DigestEmailsEnabledBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DigestEmailsEnabledBadRequest struct {
	Payload *models.Error
}

func (o *DigestEmailsEnabledBadRequest) Error() string {
	return fmt.Sprintf("[GET /digest_emails_enabled][%d] digestEmailsEnabledBadRequest  %+v", 400, o.Payload)
}
func (o *DigestEmailsEnabledBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DigestEmailsEnabledBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDigestEmailsEnabledNotFound creates a DigestEmailsEnabledNotFound with default headers values
func NewDigestEmailsEnabledNotFound() *DigestEmailsEnabledNotFound {
	return &DigestEmailsEnabledNotFound{}
}

/* DigestEmailsEnabledNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DigestEmailsEnabledNotFound struct {
	Payload *models.Error
}

func (o *DigestEmailsEnabledNotFound) Error() string {
	return fmt.Sprintf("[GET /digest_emails_enabled][%d] digestEmailsEnabledNotFound  %+v", 404, o.Payload)
}
func (o *DigestEmailsEnabledNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DigestEmailsEnabledNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
