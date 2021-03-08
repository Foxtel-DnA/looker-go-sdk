// Code generated by go-swagger; DO NOT EDIT.

package color_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateColorCollectionReader is a Reader for the UpdateColorCollection structure.
type UpdateColorCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateColorCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateColorCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateColorCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateColorCollectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateColorCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateColorCollectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateColorCollectionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateColorCollectionOK creates a UpdateColorCollectionOK with default headers values
func NewUpdateColorCollectionOK() *UpdateColorCollectionOK {
	return &UpdateColorCollectionOK{}
}

/* UpdateColorCollectionOK describes a response with status code 200, with default header values.

ColorCollection
*/
type UpdateColorCollectionOK struct {
	Payload *models.ColorCollection
}

func (o *UpdateColorCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionOK  %+v", 200, o.Payload)
}
func (o *UpdateColorCollectionOK) GetPayload() *models.ColorCollection {
	return o.Payload
}

func (o *UpdateColorCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ColorCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColorCollectionBadRequest creates a UpdateColorCollectionBadRequest with default headers values
func NewUpdateColorCollectionBadRequest() *UpdateColorCollectionBadRequest {
	return &UpdateColorCollectionBadRequest{}
}

/* UpdateColorCollectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateColorCollectionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateColorCollectionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateColorCollectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColorCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColorCollectionForbidden creates a UpdateColorCollectionForbidden with default headers values
func NewUpdateColorCollectionForbidden() *UpdateColorCollectionForbidden {
	return &UpdateColorCollectionForbidden{}
}

/* UpdateColorCollectionForbidden describes a response with status code 403, with default header values.

Permission Denied
*/
type UpdateColorCollectionForbidden struct {
	Payload *models.Error
}

func (o *UpdateColorCollectionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionForbidden  %+v", 403, o.Payload)
}
func (o *UpdateColorCollectionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColorCollectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColorCollectionNotFound creates a UpdateColorCollectionNotFound with default headers values
func NewUpdateColorCollectionNotFound() *UpdateColorCollectionNotFound {
	return &UpdateColorCollectionNotFound{}
}

/* UpdateColorCollectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateColorCollectionNotFound struct {
	Payload *models.Error
}

func (o *UpdateColorCollectionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateColorCollectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColorCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColorCollectionUnprocessableEntity creates a UpdateColorCollectionUnprocessableEntity with default headers values
func NewUpdateColorCollectionUnprocessableEntity() *UpdateColorCollectionUnprocessableEntity {
	return &UpdateColorCollectionUnprocessableEntity{}
}

/* UpdateColorCollectionUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateColorCollectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateColorCollectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateColorCollectionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateColorCollectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColorCollectionTooManyRequests creates a UpdateColorCollectionTooManyRequests with default headers values
func NewUpdateColorCollectionTooManyRequests() *UpdateColorCollectionTooManyRequests {
	return &UpdateColorCollectionTooManyRequests{}
}

/* UpdateColorCollectionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateColorCollectionTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateColorCollectionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /color_collections/{collection_id}][%d] updateColorCollectionTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateColorCollectionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColorCollectionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}