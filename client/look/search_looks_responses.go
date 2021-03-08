// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// SearchLooksReader is a Reader for the SearchLooks structure.
type SearchLooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchLooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLooksOK creates a SearchLooksOK with default headers values
func NewSearchLooksOK() *SearchLooksOK {
	return &SearchLooksOK{}
}

/* SearchLooksOK describes a response with status code 200, with default header values.

looks
*/
type SearchLooksOK struct {
	Payload []*models.Look
}

func (o *SearchLooksOK) Error() string {
	return fmt.Sprintf("[GET /looks/search][%d] searchLooksOK  %+v", 200, o.Payload)
}
func (o *SearchLooksOK) GetPayload() []*models.Look {
	return o.Payload
}

func (o *SearchLooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLooksBadRequest creates a SearchLooksBadRequest with default headers values
func NewSearchLooksBadRequest() *SearchLooksBadRequest {
	return &SearchLooksBadRequest{}
}

/* SearchLooksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchLooksBadRequest struct {
	Payload *models.Error
}

func (o *SearchLooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /looks/search][%d] searchLooksBadRequest  %+v", 400, o.Payload)
}
func (o *SearchLooksBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchLooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLooksNotFound creates a SearchLooksNotFound with default headers values
func NewSearchLooksNotFound() *SearchLooksNotFound {
	return &SearchLooksNotFound{}
}

/* SearchLooksNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchLooksNotFound struct {
	Payload *models.Error
}

func (o *SearchLooksNotFound) Error() string {
	return fmt.Sprintf("[GET /looks/search][%d] searchLooksNotFound  %+v", 404, o.Payload)
}
func (o *SearchLooksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchLooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
