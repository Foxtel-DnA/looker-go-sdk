// Code generated by go-swagger; DO NOT EDIT.

package api_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// LoginUserReader is a Reader for the LoginUser structure.
type LoginUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoginUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoginUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoginUserOK creates a LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {
	return &LoginUserOK{}
}

/* LoginUserOK describes a response with status code 200, with default header values.

Access token with metadata.
*/
type LoginUserOK struct {
	Payload *models.AccessToken
}

func (o *LoginUserOK) Error() string {
	return fmt.Sprintf("[POST /login/{user_id}][%d] loginUserOK  %+v", 200, o.Payload)
}
func (o *LoginUserOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *LoginUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUserBadRequest creates a LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {
	return &LoginUserBadRequest{}
}

/* LoginUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type LoginUserBadRequest struct {
	Payload *models.Error
}

func (o *LoginUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /login/{user_id}][%d] loginUserBadRequest  %+v", 400, o.Payload)
}
func (o *LoginUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *LoginUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUserNotFound creates a LoginUserNotFound with default headers values
func NewLoginUserNotFound() *LoginUserNotFound {
	return &LoginUserNotFound{}
}

/* LoginUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LoginUserNotFound struct {
	Payload *models.Error
}

func (o *LoginUserNotFound) Error() string {
	return fmt.Sprintf("[POST /login/{user_id}][%d] loginUserNotFound  %+v", 404, o.Payload)
}
func (o *LoginUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *LoginUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
