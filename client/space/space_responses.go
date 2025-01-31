// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// SpaceReader is a Reader for the Space structure.
type SpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSpaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSpaceOK creates a SpaceOK with default headers values
func NewSpaceOK() *SpaceOK {
	return &SpaceOK{}
}

/* SpaceOK describes a response with status code 200, with default header values.

Space
*/
type SpaceOK struct {
	Payload *models.Space
}

func (o *SpaceOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}][%d] spaceOK  %+v", 200, o.Payload)
}
func (o *SpaceOK) GetPayload() *models.Space {
	return o.Payload
}

func (o *SpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Space)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceBadRequest creates a SpaceBadRequest with default headers values
func NewSpaceBadRequest() *SpaceBadRequest {
	return &SpaceBadRequest{}
}

/* SpaceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SpaceBadRequest struct {
	Payload *models.Error
}

func (o *SpaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}][%d] spaceBadRequest  %+v", 400, o.Payload)
}
func (o *SpaceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceNotFound creates a SpaceNotFound with default headers values
func NewSpaceNotFound() *SpaceNotFound {
	return &SpaceNotFound{}
}

/* SpaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SpaceNotFound struct {
	Payload *models.Error
}

func (o *SpaceNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}][%d] spaceNotFound  %+v", 404, o.Payload)
}
func (o *SpaceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
