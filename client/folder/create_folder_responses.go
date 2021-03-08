// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// CreateFolderReader is a Reader for the CreateFolder structure.
type CreateFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFolderConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateFolderUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFolderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFolderOK creates a CreateFolderOK with default headers values
func NewCreateFolderOK() *CreateFolderOK {
	return &CreateFolderOK{}
}

/* CreateFolderOK describes a response with status code 200, with default header values.

Folder
*/
type CreateFolderOK struct {
	Payload *models.Folder
}

func (o *CreateFolderOK) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderOK  %+v", 200, o.Payload)
}
func (o *CreateFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *CreateFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFolderBadRequest creates a CreateFolderBadRequest with default headers values
func NewCreateFolderBadRequest() *CreateFolderBadRequest {
	return &CreateFolderBadRequest{}
}

/* CreateFolderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateFolderBadRequest struct {
	Payload *models.Error
}

func (o *CreateFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFolderBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFolderNotFound creates a CreateFolderNotFound with default headers values
func NewCreateFolderNotFound() *CreateFolderNotFound {
	return &CreateFolderNotFound{}
}

/* CreateFolderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateFolderNotFound struct {
	Payload *models.Error
}

func (o *CreateFolderNotFound) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderNotFound  %+v", 404, o.Payload)
}
func (o *CreateFolderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFolderConflict creates a CreateFolderConflict with default headers values
func NewCreateFolderConflict() *CreateFolderConflict {
	return &CreateFolderConflict{}
}

/* CreateFolderConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateFolderConflict struct {
	Payload *models.Error
}

func (o *CreateFolderConflict) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderConflict  %+v", 409, o.Payload)
}
func (o *CreateFolderConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFolderConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFolderUnprocessableEntity creates a CreateFolderUnprocessableEntity with default headers values
func NewCreateFolderUnprocessableEntity() *CreateFolderUnprocessableEntity {
	return &CreateFolderUnprocessableEntity{}
}

/* CreateFolderUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateFolderUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateFolderUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateFolderUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateFolderUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFolderTooManyRequests creates a CreateFolderTooManyRequests with default headers values
func NewCreateFolderTooManyRequests() *CreateFolderTooManyRequests {
	return &CreateFolderTooManyRequests{}
}

/* CreateFolderTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateFolderTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateFolderTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /folders][%d] createFolderTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateFolderTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFolderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
