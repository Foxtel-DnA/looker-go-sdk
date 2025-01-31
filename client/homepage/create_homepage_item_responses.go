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

// CreateHomepageItemReader is a Reader for the CreateHomepageItem structure.
type CreateHomepageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHomepageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHomepageItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHomepageItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateHomepageItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHomepageItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateHomepageItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateHomepageItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateHomepageItemOK creates a CreateHomepageItemOK with default headers values
func NewCreateHomepageItemOK() *CreateHomepageItemOK {
	return &CreateHomepageItemOK{}
}

/* CreateHomepageItemOK describes a response with status code 200, with default header values.

Homepage Item
*/
type CreateHomepageItemOK struct {
	Payload *models.HomepageItem
}

func (o *CreateHomepageItemOK) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemOK  %+v", 200, o.Payload)
}
func (o *CreateHomepageItemOK) GetPayload() *models.HomepageItem {
	return o.Payload
}

func (o *CreateHomepageItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomepageItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHomepageItemBadRequest creates a CreateHomepageItemBadRequest with default headers values
func NewCreateHomepageItemBadRequest() *CreateHomepageItemBadRequest {
	return &CreateHomepageItemBadRequest{}
}

/* CreateHomepageItemBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateHomepageItemBadRequest struct {
	Payload *models.Error
}

func (o *CreateHomepageItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHomepageItemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHomepageItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHomepageItemNotFound creates a CreateHomepageItemNotFound with default headers values
func NewCreateHomepageItemNotFound() *CreateHomepageItemNotFound {
	return &CreateHomepageItemNotFound{}
}

/* CreateHomepageItemNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateHomepageItemNotFound struct {
	Payload *models.Error
}

func (o *CreateHomepageItemNotFound) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemNotFound  %+v", 404, o.Payload)
}
func (o *CreateHomepageItemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHomepageItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHomepageItemConflict creates a CreateHomepageItemConflict with default headers values
func NewCreateHomepageItemConflict() *CreateHomepageItemConflict {
	return &CreateHomepageItemConflict{}
}

/* CreateHomepageItemConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateHomepageItemConflict struct {
	Payload *models.Error
}

func (o *CreateHomepageItemConflict) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemConflict  %+v", 409, o.Payload)
}
func (o *CreateHomepageItemConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHomepageItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHomepageItemUnprocessableEntity creates a CreateHomepageItemUnprocessableEntity with default headers values
func NewCreateHomepageItemUnprocessableEntity() *CreateHomepageItemUnprocessableEntity {
	return &CreateHomepageItemUnprocessableEntity{}
}

/* CreateHomepageItemUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateHomepageItemUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateHomepageItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateHomepageItemUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateHomepageItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHomepageItemTooManyRequests creates a CreateHomepageItemTooManyRequests with default headers values
func NewCreateHomepageItemTooManyRequests() *CreateHomepageItemTooManyRequests {
	return &CreateHomepageItemTooManyRequests{}
}

/* CreateHomepageItemTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateHomepageItemTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateHomepageItemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /homepage_items][%d] createHomepageItemTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateHomepageItemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHomepageItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
