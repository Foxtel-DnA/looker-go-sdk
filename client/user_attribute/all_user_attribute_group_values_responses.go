// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// AllUserAttributeGroupValuesReader is a Reader for the AllUserAttributeGroupValues structure.
type AllUserAttributeGroupValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllUserAttributeGroupValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllUserAttributeGroupValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllUserAttributeGroupValuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllUserAttributeGroupValuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllUserAttributeGroupValuesOK creates a AllUserAttributeGroupValuesOK with default headers values
func NewAllUserAttributeGroupValuesOK() *AllUserAttributeGroupValuesOK {
	return &AllUserAttributeGroupValuesOK{}
}

/* AllUserAttributeGroupValuesOK describes a response with status code 200, with default header values.

All group values for attribute.
*/
type AllUserAttributeGroupValuesOK struct {
	Payload []*models.UserAttributeGroupValue
}

func (o *AllUserAttributeGroupValuesOK) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}/group_values][%d] allUserAttributeGroupValuesOK  %+v", 200, o.Payload)
}
func (o *AllUserAttributeGroupValuesOK) GetPayload() []*models.UserAttributeGroupValue {
	return o.Payload
}

func (o *AllUserAttributeGroupValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserAttributeGroupValuesBadRequest creates a AllUserAttributeGroupValuesBadRequest with default headers values
func NewAllUserAttributeGroupValuesBadRequest() *AllUserAttributeGroupValuesBadRequest {
	return &AllUserAttributeGroupValuesBadRequest{}
}

/* AllUserAttributeGroupValuesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllUserAttributeGroupValuesBadRequest struct {
	Payload *models.Error
}

func (o *AllUserAttributeGroupValuesBadRequest) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}/group_values][%d] allUserAttributeGroupValuesBadRequest  %+v", 400, o.Payload)
}
func (o *AllUserAttributeGroupValuesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllUserAttributeGroupValuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserAttributeGroupValuesNotFound creates a AllUserAttributeGroupValuesNotFound with default headers values
func NewAllUserAttributeGroupValuesNotFound() *AllUserAttributeGroupValuesNotFound {
	return &AllUserAttributeGroupValuesNotFound{}
}

/* AllUserAttributeGroupValuesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllUserAttributeGroupValuesNotFound struct {
	Payload *models.Error
}

func (o *AllUserAttributeGroupValuesNotFound) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}/group_values][%d] allUserAttributeGroupValuesNotFound  %+v", 404, o.Payload)
}
func (o *AllUserAttributeGroupValuesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllUserAttributeGroupValuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
