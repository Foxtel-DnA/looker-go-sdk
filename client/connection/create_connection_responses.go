// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateConnectionReader is a Reader for the CreateConnection structure.
type CreateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateConnectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateConnectionOK creates a CreateConnectionOK with default headers values
func NewCreateConnectionOK() *CreateConnectionOK {
	return &CreateConnectionOK{}
}

/*CreateConnectionOK handles this case with default header values.

Connection
*/
type CreateConnectionOK struct {
	Payload *models.DBConnection
}

func (o *CreateConnectionOK) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionOK  %+v", 200, o.Payload)
}

func (o *CreateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DBConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionBadRequest creates a CreateConnectionBadRequest with default headers values
func NewCreateConnectionBadRequest() *CreateConnectionBadRequest {
	return &CreateConnectionBadRequest{}
}

/*CreateConnectionBadRequest handles this case with default header values.

Bad Request
*/
type CreateConnectionBadRequest struct {
	Payload *models.Error
}

func (o *CreateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionNotFound creates a CreateConnectionNotFound with default headers values
func NewCreateConnectionNotFound() *CreateConnectionNotFound {
	return &CreateConnectionNotFound{}
}

/*CreateConnectionNotFound handles this case with default header values.

Not Found
*/
type CreateConnectionNotFound struct {
	Payload *models.Error
}

func (o *CreateConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionNotFound  %+v", 404, o.Payload)
}

func (o *CreateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionConflict creates a CreateConnectionConflict with default headers values
func NewCreateConnectionConflict() *CreateConnectionConflict {
	return &CreateConnectionConflict{}
}

/*CreateConnectionConflict handles this case with default header values.

Resource Already Exists
*/
type CreateConnectionConflict struct {
	Payload *models.Error
}

func (o *CreateConnectionConflict) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionConflict  %+v", 409, o.Payload)
}

func (o *CreateConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionUnprocessableEntity creates a CreateConnectionUnprocessableEntity with default headers values
func NewCreateConnectionUnprocessableEntity() *CreateConnectionUnprocessableEntity {
	return &CreateConnectionUnprocessableEntity{}
}

/*CreateConnectionUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateConnectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateConnectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateConnectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}