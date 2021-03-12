// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UserCredentialsAPIReader is a Reader for the UserCredentialsAPI structure.
type UserCredentialsAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCredentialsAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserCredentialsAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserCredentialsAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCredentialsAPIOK creates a UserCredentialsAPIOK with default headers values
func NewUserCredentialsAPIOK() *UserCredentialsAPIOK {
	return &UserCredentialsAPIOK{}
}

/*UserCredentialsAPIOK handles this case with default header values.

API Credential
*/
type UserCredentialsAPIOK struct {
	Payload *models.CredentialsAPI
}

func (o *UserCredentialsAPIOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api][%d] userCredentialsApiOK  %+v", 200, o.Payload)
}

func (o *UserCredentialsAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsAPI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsAPIBadRequest creates a UserCredentialsAPIBadRequest with default headers values
func NewUserCredentialsAPIBadRequest() *UserCredentialsAPIBadRequest {
	return &UserCredentialsAPIBadRequest{}
}

/*UserCredentialsAPIBadRequest handles this case with default header values.

Bad Request
*/
type UserCredentialsAPIBadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsAPIBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api][%d] userCredentialsApiBadRequest  %+v", 400, o.Payload)
}

func (o *UserCredentialsAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsAPINotFound creates a UserCredentialsAPINotFound with default headers values
func NewUserCredentialsAPINotFound() *UserCredentialsAPINotFound {
	return &UserCredentialsAPINotFound{}
}

/*UserCredentialsAPINotFound handles this case with default header values.

Not Found
*/
type UserCredentialsAPINotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsAPINotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api][%d] userCredentialsApiNotFound  %+v", 404, o.Payload)
}

func (o *UserCredentialsAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
