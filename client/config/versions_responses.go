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

// VersionsReader is a Reader for the Versions structure.
type VersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVersionsOK creates a VersionsOK with default headers values
func NewVersionsOK() *VersionsOK {
	return &VersionsOK{}
}

/* VersionsOK describes a response with status code 200, with default header values.

ApiVersion
*/
type VersionsOK struct {
	Payload *models.APIVersion
}

func (o *VersionsOK) Error() string {
	return fmt.Sprintf("[GET /versions][%d] versionsOK  %+v", 200, o.Payload)
}
func (o *VersionsOK) GetPayload() *models.APIVersion {
	return o.Payload
}

func (o *VersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsBadRequest creates a VersionsBadRequest with default headers values
func NewVersionsBadRequest() *VersionsBadRequest {
	return &VersionsBadRequest{}
}

/* VersionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VersionsBadRequest struct {
	Payload *models.Error
}

func (o *VersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /versions][%d] versionsBadRequest  %+v", 400, o.Payload)
}
func (o *VersionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *VersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsNotFound creates a VersionsNotFound with default headers values
func NewVersionsNotFound() *VersionsNotFound {
	return &VersionsNotFound{}
}

/* VersionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type VersionsNotFound struct {
	Payload *models.Error
}

func (o *VersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /versions][%d] versionsNotFound  %+v", 404, o.Payload)
}
func (o *VersionsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *VersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
