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

// UpdateContentMetadataAccessReader is a Reader for the UpdateContentMetadataAccess structure.
type UpdateContentMetadataAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateContentMetadataAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateContentMetadataAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateContentMetadataAccessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateContentMetadataAccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateContentMetadataAccessUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateContentMetadataAccessTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateContentMetadataAccessOK creates a UpdateContentMetadataAccessOK with default headers values
func NewUpdateContentMetadataAccessOK() *UpdateContentMetadataAccessOK {
	return &UpdateContentMetadataAccessOK{}
}

/* UpdateContentMetadataAccessOK describes a response with status code 200, with default header values.

Content Metadata Access
*/
type UpdateContentMetadataAccessOK struct {
	Payload *models.ContentMetaGroupUser
}

func (o *UpdateContentMetadataAccessOK) Error() string {
	return fmt.Sprintf("[PUT /content_metadata_access/{content_metadata_access_id}][%d] updateContentMetadataAccessOK  %+v", 200, o.Payload)
}
func (o *UpdateContentMetadataAccessOK) GetPayload() *models.ContentMetaGroupUser {
	return o.Payload
}

func (o *UpdateContentMetadataAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentMetaGroupUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentMetadataAccessBadRequest creates a UpdateContentMetadataAccessBadRequest with default headers values
func NewUpdateContentMetadataAccessBadRequest() *UpdateContentMetadataAccessBadRequest {
	return &UpdateContentMetadataAccessBadRequest{}
}

/* UpdateContentMetadataAccessBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateContentMetadataAccessBadRequest struct {
	Payload *models.Error
}

func (o *UpdateContentMetadataAccessBadRequest) Error() string {
	return fmt.Sprintf("[PUT /content_metadata_access/{content_metadata_access_id}][%d] updateContentMetadataAccessBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateContentMetadataAccessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateContentMetadataAccessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentMetadataAccessNotFound creates a UpdateContentMetadataAccessNotFound with default headers values
func NewUpdateContentMetadataAccessNotFound() *UpdateContentMetadataAccessNotFound {
	return &UpdateContentMetadataAccessNotFound{}
}

/* UpdateContentMetadataAccessNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateContentMetadataAccessNotFound struct {
	Payload *models.Error
}

func (o *UpdateContentMetadataAccessNotFound) Error() string {
	return fmt.Sprintf("[PUT /content_metadata_access/{content_metadata_access_id}][%d] updateContentMetadataAccessNotFound  %+v", 404, o.Payload)
}
func (o *UpdateContentMetadataAccessNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateContentMetadataAccessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentMetadataAccessUnprocessableEntity creates a UpdateContentMetadataAccessUnprocessableEntity with default headers values
func NewUpdateContentMetadataAccessUnprocessableEntity() *UpdateContentMetadataAccessUnprocessableEntity {
	return &UpdateContentMetadataAccessUnprocessableEntity{}
}

/* UpdateContentMetadataAccessUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateContentMetadataAccessUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateContentMetadataAccessUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /content_metadata_access/{content_metadata_access_id}][%d] updateContentMetadataAccessUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateContentMetadataAccessUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateContentMetadataAccessUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentMetadataAccessTooManyRequests creates a UpdateContentMetadataAccessTooManyRequests with default headers values
func NewUpdateContentMetadataAccessTooManyRequests() *UpdateContentMetadataAccessTooManyRequests {
	return &UpdateContentMetadataAccessTooManyRequests{}
}

/* UpdateContentMetadataAccessTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateContentMetadataAccessTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateContentMetadataAccessTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /content_metadata_access/{content_metadata_access_id}][%d] updateContentMetadataAccessTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateContentMetadataAccessTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateContentMetadataAccessTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
