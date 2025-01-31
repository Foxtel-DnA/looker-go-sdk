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

// ContentThumbnailReader is a Reader for the ContentThumbnail structure.
type ContentThumbnailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContentThumbnailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContentThumbnailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContentThumbnailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContentThumbnailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContentThumbnailOK creates a ContentThumbnailOK with default headers values
func NewContentThumbnailOK() *ContentThumbnailOK {
	return &ContentThumbnailOK{}
}

/* ContentThumbnailOK describes a response with status code 200, with default header values.

Content thumbnail
*/
type ContentThumbnailOK struct {
	Payload string
}

func (o *ContentThumbnailOK) Error() string {
	return fmt.Sprintf("[GET /content_thumbnail/{type}/{resource_id}][%d] contentThumbnailOK  %+v", 200, o.Payload)
}
func (o *ContentThumbnailOK) GetPayload() string {
	return o.Payload
}

func (o *ContentThumbnailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentThumbnailBadRequest creates a ContentThumbnailBadRequest with default headers values
func NewContentThumbnailBadRequest() *ContentThumbnailBadRequest {
	return &ContentThumbnailBadRequest{}
}

/* ContentThumbnailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ContentThumbnailBadRequest struct {
	Payload *models.Error
}

func (o *ContentThumbnailBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_thumbnail/{type}/{resource_id}][%d] contentThumbnailBadRequest  %+v", 400, o.Payload)
}
func (o *ContentThumbnailBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentThumbnailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentThumbnailNotFound creates a ContentThumbnailNotFound with default headers values
func NewContentThumbnailNotFound() *ContentThumbnailNotFound {
	return &ContentThumbnailNotFound{}
}

/* ContentThumbnailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ContentThumbnailNotFound struct {
	Payload *models.Error
}

func (o *ContentThumbnailNotFound) Error() string {
	return fmt.Sprintf("[GET /content_thumbnail/{type}/{resource_id}][%d] contentThumbnailNotFound  %+v", 404, o.Payload)
}
func (o *ContentThumbnailNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentThumbnailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
