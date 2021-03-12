// Code generated by go-swagger; DO NOT EDIT.

package color_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// DeleteColorCollectionReader is a Reader for the DeleteColorCollection structure.
type DeleteColorCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteColorCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteColorCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteColorCollectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteColorCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteColorCollectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteColorCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteColorCollectionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteColorCollectionOK creates a DeleteColorCollectionOK with default headers values
func NewDeleteColorCollectionOK() *DeleteColorCollectionOK {
	return &DeleteColorCollectionOK{}
}

/* DeleteColorCollectionOK describes a response with status code 200, with default header values.

Success
*/
type DeleteColorCollectionOK struct {
}

func (o *DeleteColorCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionOK ", 200)
}

func (o *DeleteColorCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteColorCollectionNoContent creates a DeleteColorCollectionNoContent with default headers values
func NewDeleteColorCollectionNoContent() *DeleteColorCollectionNoContent {
	return &DeleteColorCollectionNoContent{}
}

/* DeleteColorCollectionNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteColorCollectionNoContent struct {
	Payload string
}

func (o *DeleteColorCollectionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionNoContent  %+v", 204, o.Payload)
}
func (o *DeleteColorCollectionNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteColorCollectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteColorCollectionBadRequest creates a DeleteColorCollectionBadRequest with default headers values
func NewDeleteColorCollectionBadRequest() *DeleteColorCollectionBadRequest {
	return &DeleteColorCollectionBadRequest{}
}

/* DeleteColorCollectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteColorCollectionBadRequest struct {
	Payload *models.Error
}

func (o *DeleteColorCollectionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteColorCollectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColorCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteColorCollectionForbidden creates a DeleteColorCollectionForbidden with default headers values
func NewDeleteColorCollectionForbidden() *DeleteColorCollectionForbidden {
	return &DeleteColorCollectionForbidden{}
}

/* DeleteColorCollectionForbidden describes a response with status code 403, with default header values.

Permission Denied
*/
type DeleteColorCollectionForbidden struct {
	Payload *models.Error
}

func (o *DeleteColorCollectionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionForbidden  %+v", 403, o.Payload)
}
func (o *DeleteColorCollectionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColorCollectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteColorCollectionNotFound creates a DeleteColorCollectionNotFound with default headers values
func NewDeleteColorCollectionNotFound() *DeleteColorCollectionNotFound {
	return &DeleteColorCollectionNotFound{}
}

/* DeleteColorCollectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteColorCollectionNotFound struct {
	Payload *models.Error
}

func (o *DeleteColorCollectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteColorCollectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColorCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteColorCollectionTooManyRequests creates a DeleteColorCollectionTooManyRequests with default headers values
func NewDeleteColorCollectionTooManyRequests() *DeleteColorCollectionTooManyRequests {
	return &DeleteColorCollectionTooManyRequests{}
}

/* DeleteColorCollectionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteColorCollectionTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteColorCollectionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /color_collections/{collection_id}][%d] deleteColorCollectionTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteColorCollectionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColorCollectionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
