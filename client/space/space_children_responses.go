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

// SpaceChildrenReader is a Reader for the SpaceChildren structure.
type SpaceChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpaceChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSpaceChildrenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSpaceChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSpaceChildrenOK creates a SpaceChildrenOK with default headers values
func NewSpaceChildrenOK() *SpaceChildrenOK {
	return &SpaceChildrenOK{}
}

/* SpaceChildrenOK describes a response with status code 200, with default header values.

Spaces
*/
type SpaceChildrenOK struct {
	Payload []*models.Space
}

func (o *SpaceChildrenOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenOK  %+v", 200, o.Payload)
}
func (o *SpaceChildrenOK) GetPayload() []*models.Space {
	return o.Payload
}

func (o *SpaceChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceChildrenBadRequest creates a SpaceChildrenBadRequest with default headers values
func NewSpaceChildrenBadRequest() *SpaceChildrenBadRequest {
	return &SpaceChildrenBadRequest{}
}

/* SpaceChildrenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SpaceChildrenBadRequest struct {
	Payload *models.Error
}

func (o *SpaceChildrenBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenBadRequest  %+v", 400, o.Payload)
}
func (o *SpaceChildrenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceChildrenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceChildrenNotFound creates a SpaceChildrenNotFound with default headers values
func NewSpaceChildrenNotFound() *SpaceChildrenNotFound {
	return &SpaceChildrenNotFound{}
}

/* SpaceChildrenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SpaceChildrenNotFound struct {
	Payload *models.Error
}

func (o *SpaceChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenNotFound  %+v", 404, o.Payload)
}
func (o *SpaceChildrenNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SpaceChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
