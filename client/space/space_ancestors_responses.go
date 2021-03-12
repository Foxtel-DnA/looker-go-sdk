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

// SpaceAncestorsReader is a Reader for the SpaceAncestors structure.
type SpaceAncestorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceAncestorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpaceAncestorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSpaceAncestorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSpaceAncestorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSpaceAncestorsOK creates a SpaceAncestorsOK with default headers values
func NewSpaceAncestorsOK() *SpaceAncestorsOK {
	return &SpaceAncestorsOK{}
}

/* SpaceAncestorsOK describes a response with status code 200, with default header values.

Spaces
*/
type SpaceAncestorsOK struct {
	Payload []*models.Space
}

func (o *SpaceAncestorsOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/ancestors][%d] spaceAncestorsOK  %+v", 200, o.Payload)
}
func (o *SpaceAncestorsOK) GetPayload() []*models.Space {
	return o.Payload
}

func (o *SpaceAncestorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceAncestorsBadRequest creates a SpaceAncestorsBadRequest with default headers values
func NewSpaceAncestorsBadRequest() *SpaceAncestorsBadRequest {
	return &SpaceAncestorsBadRequest{}
}

/* SpaceAncestorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SpaceAncestorsBadRequest struct {
	Payload *models.Error
}

func (o *SpaceAncestorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/ancestors][%d] spaceAncestorsBadRequest  %+v", 400, o.Payload)
}
func (o *SpaceAncestorsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceAncestorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceAncestorsNotFound creates a SpaceAncestorsNotFound with default headers values
func NewSpaceAncestorsNotFound() *SpaceAncestorsNotFound {
	return &SpaceAncestorsNotFound{}
}

/* SpaceAncestorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SpaceAncestorsNotFound struct {
	Payload *models.Error
}

func (o *SpaceAncestorsNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/ancestors][%d] spaceAncestorsNotFound  %+v", 404, o.Payload)
}
func (o *SpaceAncestorsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceAncestorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
