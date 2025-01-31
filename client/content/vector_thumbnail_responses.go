// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// VectorThumbnailReader is a Reader for the VectorThumbnail structure.
type VectorThumbnailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VectorThumbnailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVectorThumbnailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVectorThumbnailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVectorThumbnailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVectorThumbnailOK creates a VectorThumbnailOK with default headers values
func NewVectorThumbnailOK() *VectorThumbnailOK {
	return &VectorThumbnailOK{}
}

/* VectorThumbnailOK describes a response with status code 200, with default header values.

Vector thumbnail
*/
type VectorThumbnailOK struct {
	Payload string
}

func (o *VectorThumbnailOK) Error() string {
	return fmt.Sprintf("[GET /vector_thumbnail/{type}/{resource_id}][%d] vectorThumbnailOK  %+v", 200, o.Payload)
}
func (o *VectorThumbnailOK) GetPayload() string {
	return o.Payload
}

func (o *VectorThumbnailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVectorThumbnailBadRequest creates a VectorThumbnailBadRequest with default headers values
func NewVectorThumbnailBadRequest() *VectorThumbnailBadRequest {
	return &VectorThumbnailBadRequest{}
}

/* VectorThumbnailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VectorThumbnailBadRequest struct {
	Payload *models.Error
}

func (o *VectorThumbnailBadRequest) Error() string {
	return fmt.Sprintf("[GET /vector_thumbnail/{type}/{resource_id}][%d] vectorThumbnailBadRequest  %+v", 400, o.Payload)
}
func (o *VectorThumbnailBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *VectorThumbnailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVectorThumbnailNotFound creates a VectorThumbnailNotFound with default headers values
func NewVectorThumbnailNotFound() *VectorThumbnailNotFound {
	return &VectorThumbnailNotFound{}
}

/* VectorThumbnailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type VectorThumbnailNotFound struct {
	Payload *models.Error
}

func (o *VectorThumbnailNotFound) Error() string {
	return fmt.Sprintf("[GET /vector_thumbnail/{type}/{resource_id}][%d] vectorThumbnailNotFound  %+v", 404, o.Payload)
}
func (o *VectorThumbnailNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *VectorThumbnailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
