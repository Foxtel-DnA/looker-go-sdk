// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateLookReader is a Reader for the CreateLook structure.
type CreateLookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateLookConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateLookUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLookOK creates a CreateLookOK with default headers values
func NewCreateLookOK() *CreateLookOK {
	return &CreateLookOK{}
}

/*CreateLookOK handles this case with default header values.

Look
*/
type CreateLookOK struct {
	Payload *models.LookWithQuery
}

func (o *CreateLookOK) Error() string {
	return fmt.Sprintf("[POST /looks][%d] createLookOK  %+v", 200, o.Payload)
}

func (o *CreateLookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookWithQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookBadRequest creates a CreateLookBadRequest with default headers values
func NewCreateLookBadRequest() *CreateLookBadRequest {
	return &CreateLookBadRequest{}
}

/*CreateLookBadRequest handles this case with default header values.

Bad Request
*/
type CreateLookBadRequest struct {
	Payload *models.Error
}

func (o *CreateLookBadRequest) Error() string {
	return fmt.Sprintf("[POST /looks][%d] createLookBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookNotFound creates a CreateLookNotFound with default headers values
func NewCreateLookNotFound() *CreateLookNotFound {
	return &CreateLookNotFound{}
}

/*CreateLookNotFound handles this case with default header values.

Not Found
*/
type CreateLookNotFound struct {
	Payload *models.Error
}

func (o *CreateLookNotFound) Error() string {
	return fmt.Sprintf("[POST /looks][%d] createLookNotFound  %+v", 404, o.Payload)
}

func (o *CreateLookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookConflict creates a CreateLookConflict with default headers values
func NewCreateLookConflict() *CreateLookConflict {
	return &CreateLookConflict{}
}

/*CreateLookConflict handles this case with default header values.

Resource Already Exists
*/
type CreateLookConflict struct {
	Payload *models.Error
}

func (o *CreateLookConflict) Error() string {
	return fmt.Sprintf("[POST /looks][%d] createLookConflict  %+v", 409, o.Payload)
}

func (o *CreateLookConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookUnprocessableEntity creates a CreateLookUnprocessableEntity with default headers values
func NewCreateLookUnprocessableEntity() *CreateLookUnprocessableEntity {
	return &CreateLookUnprocessableEntity{}
}

/*CreateLookUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateLookUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateLookUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /looks][%d] createLookUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateLookUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
