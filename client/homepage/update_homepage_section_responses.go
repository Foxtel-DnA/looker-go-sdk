// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UpdateHomepageSectionReader is a Reader for the UpdateHomepageSection structure.
type UpdateHomepageSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHomepageSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateHomepageSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateHomepageSectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateHomepageSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateHomepageSectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateHomepageSectionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateHomepageSectionOK creates a UpdateHomepageSectionOK with default headers values
func NewUpdateHomepageSectionOK() *UpdateHomepageSectionOK {
	return &UpdateHomepageSectionOK{}
}

/* UpdateHomepageSectionOK describes a response with status code 200, with default header values.

Homepage section
*/
type UpdateHomepageSectionOK struct {
	Payload *models.HomepageSection
}

func (o *UpdateHomepageSectionOK) Error() string {
	return fmt.Sprintf("[PATCH /homepage_sections/{homepage_section_id}][%d] updateHomepageSectionOK  %+v", 200, o.Payload)
}
func (o *UpdateHomepageSectionOK) GetPayload() *models.HomepageSection {
	return o.Payload
}

func (o *UpdateHomepageSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomepageSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageSectionBadRequest creates a UpdateHomepageSectionBadRequest with default headers values
func NewUpdateHomepageSectionBadRequest() *UpdateHomepageSectionBadRequest {
	return &UpdateHomepageSectionBadRequest{}
}

/* UpdateHomepageSectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateHomepageSectionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateHomepageSectionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /homepage_sections/{homepage_section_id}][%d] updateHomepageSectionBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateHomepageSectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageSectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageSectionNotFound creates a UpdateHomepageSectionNotFound with default headers values
func NewUpdateHomepageSectionNotFound() *UpdateHomepageSectionNotFound {
	return &UpdateHomepageSectionNotFound{}
}

/* UpdateHomepageSectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateHomepageSectionNotFound struct {
	Payload *models.Error
}

func (o *UpdateHomepageSectionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /homepage_sections/{homepage_section_id}][%d] updateHomepageSectionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateHomepageSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageSectionUnprocessableEntity creates a UpdateHomepageSectionUnprocessableEntity with default headers values
func NewUpdateHomepageSectionUnprocessableEntity() *UpdateHomepageSectionUnprocessableEntity {
	return &UpdateHomepageSectionUnprocessableEntity{}
}

/* UpdateHomepageSectionUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateHomepageSectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateHomepageSectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /homepage_sections/{homepage_section_id}][%d] updateHomepageSectionUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateHomepageSectionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateHomepageSectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageSectionTooManyRequests creates a UpdateHomepageSectionTooManyRequests with default headers values
func NewUpdateHomepageSectionTooManyRequests() *UpdateHomepageSectionTooManyRequests {
	return &UpdateHomepageSectionTooManyRequests{}
}

/* UpdateHomepageSectionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateHomepageSectionTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateHomepageSectionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /homepage_sections/{homepage_section_id}][%d] updateHomepageSectionTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateHomepageSectionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageSectionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
