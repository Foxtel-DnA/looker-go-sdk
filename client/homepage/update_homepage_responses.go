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

// UpdateHomepageReader is a Reader for the UpdateHomepage structure.
type UpdateHomepageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHomepageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateHomepageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateHomepageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateHomepageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateHomepageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateHomepageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateHomepageOK creates a UpdateHomepageOK with default headers values
func NewUpdateHomepageOK() *UpdateHomepageOK {
	return &UpdateHomepageOK{}
}

/* UpdateHomepageOK describes a response with status code 200, with default header values.

Homepage
*/
type UpdateHomepageOK struct {
	Payload *models.Homepage
}

func (o *UpdateHomepageOK) Error() string {
	return fmt.Sprintf("[PATCH /homepages/{homepage_id}][%d] updateHomepageOK  %+v", 200, o.Payload)
}
func (o *UpdateHomepageOK) GetPayload() *models.Homepage {
	return o.Payload
}

func (o *UpdateHomepageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Homepage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageBadRequest creates a UpdateHomepageBadRequest with default headers values
func NewUpdateHomepageBadRequest() *UpdateHomepageBadRequest {
	return &UpdateHomepageBadRequest{}
}

/* UpdateHomepageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateHomepageBadRequest struct {
	Payload *models.Error
}

func (o *UpdateHomepageBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /homepages/{homepage_id}][%d] updateHomepageBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateHomepageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageNotFound creates a UpdateHomepageNotFound with default headers values
func NewUpdateHomepageNotFound() *UpdateHomepageNotFound {
	return &UpdateHomepageNotFound{}
}

/* UpdateHomepageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateHomepageNotFound struct {
	Payload *models.Error
}

func (o *UpdateHomepageNotFound) Error() string {
	return fmt.Sprintf("[PATCH /homepages/{homepage_id}][%d] updateHomepageNotFound  %+v", 404, o.Payload)
}
func (o *UpdateHomepageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageUnprocessableEntity creates a UpdateHomepageUnprocessableEntity with default headers values
func NewUpdateHomepageUnprocessableEntity() *UpdateHomepageUnprocessableEntity {
	return &UpdateHomepageUnprocessableEntity{}
}

/* UpdateHomepageUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateHomepageUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateHomepageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /homepages/{homepage_id}][%d] updateHomepageUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateHomepageUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateHomepageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageTooManyRequests creates a UpdateHomepageTooManyRequests with default headers values
func NewUpdateHomepageTooManyRequests() *UpdateHomepageTooManyRequests {
	return &UpdateHomepageTooManyRequests{}
}

/* UpdateHomepageTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateHomepageTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateHomepageTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /homepages/{homepage_id}][%d] updateHomepageTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateHomepageTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHomepageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
