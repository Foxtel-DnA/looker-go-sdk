// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// UserSessionReader is a Reader for the UserSession structure.
type UserSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserSessionOK creates a UserSessionOK with default headers values
func NewUserSessionOK() *UserSessionOK {
	return &UserSessionOK{}
}

/*UserSessionOK handles this case with default header values.

Web Login Session
*/
type UserSessionOK struct {
	Payload *models.Session
}

func (o *UserSessionOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions/{session_id}][%d] userSessionOK  %+v", 200, o.Payload)
}

func (o *UserSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSessionBadRequest creates a UserSessionBadRequest with default headers values
func NewUserSessionBadRequest() *UserSessionBadRequest {
	return &UserSessionBadRequest{}
}

/*UserSessionBadRequest handles this case with default header values.

Bad Request
*/
type UserSessionBadRequest struct {
	Payload *models.Error
}

func (o *UserSessionBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions/{session_id}][%d] userSessionBadRequest  %+v", 400, o.Payload)
}

func (o *UserSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSessionNotFound creates a UserSessionNotFound with default headers values
func NewUserSessionNotFound() *UserSessionNotFound {
	return &UserSessionNotFound{}
}

/*UserSessionNotFound handles this case with default header values.

Not Found
*/
type UserSessionNotFound struct {
	Payload *models.Error
}

func (o *UserSessionNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions/{session_id}][%d] userSessionNotFound  %+v", 404, o.Payload)
}

func (o *UserSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
