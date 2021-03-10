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

// ManifestReader is a Reader for the Manifest structure.
type ManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewManifestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManifestOK creates a ManifestOK with default headers values
func NewManifestOK() *ManifestOK {
	return &ManifestOK{}
}

/* ManifestOK describes a response with status code 200, with default header values.

Manifest
*/
type ManifestOK struct {
	Payload *models.Manifest
}

func (o *ManifestOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/manifest][%d] manifestOK  %+v", 200, o.Payload)
}
func (o *ManifestOK) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *ManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestBadRequest creates a ManifestBadRequest with default headers values
func NewManifestBadRequest() *ManifestBadRequest {
	return &ManifestBadRequest{}
}

/* ManifestBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ManifestBadRequest struct {
	Payload *models.Error
}

func (o *ManifestBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/manifest][%d] manifestBadRequest  %+v", 400, o.Payload)
}
func (o *ManifestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestNotFound creates a ManifestNotFound with default headers values
func NewManifestNotFound() *ManifestNotFound {
	return &ManifestNotFound{}
}

/* ManifestNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ManifestNotFound struct {
	Payload *models.Error
}

func (o *ManifestNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/manifest][%d] manifestNotFound  %+v", 404, o.Payload)
}
func (o *ManifestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ManifestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
