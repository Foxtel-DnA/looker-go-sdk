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

// DeleteUserCredentialsAPIReader is a Reader for the DeleteUserCredentialsAPI structure.
type DeleteUserCredentialsAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialsAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsAPINoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsAPINoContent creates a DeleteUserCredentialsAPINoContent with default headers values
func NewDeleteUserCredentialsAPINoContent() *DeleteUserCredentialsAPINoContent {
	return &DeleteUserCredentialsAPINoContent{}
}

/*DeleteUserCredentialsAPINoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsAPINoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsAPINoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api][%d] deleteUserCredentialsApiNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsAPINoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsAPIBadRequest creates a DeleteUserCredentialsAPIBadRequest with default headers values
func NewDeleteUserCredentialsAPIBadRequest() *DeleteUserCredentialsAPIBadRequest {
	return &DeleteUserCredentialsAPIBadRequest{}
}

/*DeleteUserCredentialsAPIBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsAPIBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsAPIBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api][%d] deleteUserCredentialsApiBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsAPINotFound creates a DeleteUserCredentialsAPINotFound with default headers values
func NewDeleteUserCredentialsAPINotFound() *DeleteUserCredentialsAPINotFound {
	return &DeleteUserCredentialsAPINotFound{}
}

/*DeleteUserCredentialsAPINotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsAPINotFound struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsAPINotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api][%d] deleteUserCredentialsApiNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
