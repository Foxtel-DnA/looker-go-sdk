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

// ContentMetadataReader is a Reader for the ContentMetadata structure.
type ContentMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContentMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContentMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContentMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContentMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContentMetadataOK creates a ContentMetadataOK with default headers values
func NewContentMetadataOK() *ContentMetadataOK {
	return &ContentMetadataOK{}
}

/* ContentMetadataOK describes a response with status code 200, with default header values.

Content Metadata
*/
type ContentMetadataOK struct {
	Payload *models.ContentMeta
}

func (o *ContentMetadataOK) Error() string {
	return fmt.Sprintf("[GET /content_metadata/{content_metadata_id}][%d] contentMetadataOK  %+v", 200, o.Payload)
}
func (o *ContentMetadataOK) GetPayload() *models.ContentMeta {
	return o.Payload
}

func (o *ContentMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentMetadataBadRequest creates a ContentMetadataBadRequest with default headers values
func NewContentMetadataBadRequest() *ContentMetadataBadRequest {
	return &ContentMetadataBadRequest{}
}

/* ContentMetadataBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ContentMetadataBadRequest struct {
	Payload *models.Error
}

func (o *ContentMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_metadata/{content_metadata_id}][%d] contentMetadataBadRequest  %+v", 400, o.Payload)
}
func (o *ContentMetadataBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentMetadataNotFound creates a ContentMetadataNotFound with default headers values
func NewContentMetadataNotFound() *ContentMetadataNotFound {
	return &ContentMetadataNotFound{}
}

/* ContentMetadataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ContentMetadataNotFound struct {
	Payload *models.Error
}

func (o *ContentMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /content_metadata/{content_metadata_id}][%d] contentMetadataNotFound  %+v", 404, o.Payload)
}
func (o *ContentMetadataNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ContentMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
