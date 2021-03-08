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

// ColorCollectionsCustomReader is a Reader for the ColorCollectionsCustom structure.
type ColorCollectionsCustomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColorCollectionsCustomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColorCollectionsCustomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewColorCollectionsCustomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewColorCollectionsCustomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewColorCollectionsCustomOK creates a ColorCollectionsCustomOK with default headers values
func NewColorCollectionsCustomOK() *ColorCollectionsCustomOK {
	return &ColorCollectionsCustomOK{}
}

/* ColorCollectionsCustomOK describes a response with status code 200, with default header values.

ColorCollections
*/
type ColorCollectionsCustomOK struct {
	Payload []*models.ColorCollection
}

func (o *ColorCollectionsCustomOK) Error() string {
	return fmt.Sprintf("[GET /color_collections/custom][%d] colorCollectionsCustomOK  %+v", 200, o.Payload)
}
func (o *ColorCollectionsCustomOK) GetPayload() []*models.ColorCollection {
	return o.Payload
}

func (o *ColorCollectionsCustomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionsCustomBadRequest creates a ColorCollectionsCustomBadRequest with default headers values
func NewColorCollectionsCustomBadRequest() *ColorCollectionsCustomBadRequest {
	return &ColorCollectionsCustomBadRequest{}
}

/* ColorCollectionsCustomBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ColorCollectionsCustomBadRequest struct {
	Payload *models.Error
}

func (o *ColorCollectionsCustomBadRequest) Error() string {
	return fmt.Sprintf("[GET /color_collections/custom][%d] colorCollectionsCustomBadRequest  %+v", 400, o.Payload)
}
func (o *ColorCollectionsCustomBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionsCustomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionsCustomNotFound creates a ColorCollectionsCustomNotFound with default headers values
func NewColorCollectionsCustomNotFound() *ColorCollectionsCustomNotFound {
	return &ColorCollectionsCustomNotFound{}
}

/* ColorCollectionsCustomNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ColorCollectionsCustomNotFound struct {
	Payload *models.Error
}

func (o *ColorCollectionsCustomNotFound) Error() string {
	return fmt.Sprintf("[GET /color_collections/custom][%d] colorCollectionsCustomNotFound  %+v", 404, o.Payload)
}
func (o *ColorCollectionsCustomNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionsCustomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}