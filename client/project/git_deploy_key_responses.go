// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// GitDeployKeyReader is a Reader for the GitDeployKey structure.
type GitDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GitDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGitDeployKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGitDeployKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGitDeployKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGitDeployKeyOK creates a GitDeployKeyOK with default headers values
func NewGitDeployKeyOK() *GitDeployKeyOK {
	return &GitDeployKeyOK{}
}

/* GitDeployKeyOK describes a response with status code 200, with default header values.

The text of the public key portion of the deploy_key
*/
type GitDeployKeyOK struct {
	Payload string
}

func (o *GitDeployKeyOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git/deploy_key][%d] gitDeployKeyOK  %+v", 200, o.Payload)
}
func (o *GitDeployKeyOK) GetPayload() string {
	return o.Payload
}

func (o *GitDeployKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGitDeployKeyBadRequest creates a GitDeployKeyBadRequest with default headers values
func NewGitDeployKeyBadRequest() *GitDeployKeyBadRequest {
	return &GitDeployKeyBadRequest{}
}

/* GitDeployKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GitDeployKeyBadRequest struct {
	Payload *models.Error
}

func (o *GitDeployKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git/deploy_key][%d] gitDeployKeyBadRequest  %+v", 400, o.Payload)
}
func (o *GitDeployKeyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GitDeployKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGitDeployKeyNotFound creates a GitDeployKeyNotFound with default headers values
func NewGitDeployKeyNotFound() *GitDeployKeyNotFound {
	return &GitDeployKeyNotFound{}
}

/* GitDeployKeyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GitDeployKeyNotFound struct {
	Payload *models.Error
}

func (o *GitDeployKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git/deploy_key][%d] gitDeployKeyNotFound  %+v", 404, o.Payload)
}
func (o *GitDeployKeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GitDeployKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
