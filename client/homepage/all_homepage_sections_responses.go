// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// AllHomepageSectionsReader is a Reader for the AllHomepageSections structure.
type AllHomepageSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllHomepageSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllHomepageSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllHomepageSectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllHomepageSectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllHomepageSectionsOK creates a AllHomepageSectionsOK with default headers values
func NewAllHomepageSectionsOK() *AllHomepageSectionsOK {
	return &AllHomepageSectionsOK{}
}

/* AllHomepageSectionsOK describes a response with status code 200, with default header values.

Homepage section
*/
type AllHomepageSectionsOK struct {
	Payload []*models.HomepageSection
}

func (o *AllHomepageSectionsOK) Error() string {
	return fmt.Sprintf("[GET /homepage_sections][%d] allHomepageSectionsOK  %+v", 200, o.Payload)
}
func (o *AllHomepageSectionsOK) GetPayload() []*models.HomepageSection {
	return o.Payload
}

func (o *AllHomepageSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllHomepageSectionsBadRequest creates a AllHomepageSectionsBadRequest with default headers values
func NewAllHomepageSectionsBadRequest() *AllHomepageSectionsBadRequest {
	return &AllHomepageSectionsBadRequest{}
}

/* AllHomepageSectionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllHomepageSectionsBadRequest struct {
	Payload *models.Error
}

func (o *AllHomepageSectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /homepage_sections][%d] allHomepageSectionsBadRequest  %+v", 400, o.Payload)
}
func (o *AllHomepageSectionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllHomepageSectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllHomepageSectionsNotFound creates a AllHomepageSectionsNotFound with default headers values
func NewAllHomepageSectionsNotFound() *AllHomepageSectionsNotFound {
	return &AllHomepageSectionsNotFound{}
}

/* AllHomepageSectionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllHomepageSectionsNotFound struct {
	Payload *models.Error
}

func (o *AllHomepageSectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /homepage_sections][%d] allHomepageSectionsNotFound  %+v", 404, o.Payload)
}
func (o *AllHomepageSectionsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllHomepageSectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
