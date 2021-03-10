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

// FindGitBranchReader is a Reader for the FindGitBranch structure.
type FindGitBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindGitBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindGitBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFindGitBranchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindGitBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindGitBranchOK creates a FindGitBranchOK with default headers values
func NewFindGitBranchOK() *FindGitBranchOK {
	return &FindGitBranchOK{}
}

/* FindGitBranchOK describes a response with status code 200, with default header values.

Git Branch
*/
type FindGitBranchOK struct {
	Payload *models.GitBranch
}

func (o *FindGitBranchOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branch/{branch_name}][%d] findGitBranchOK  %+v", 200, o.Payload)
}
func (o *FindGitBranchOK) GetPayload() *models.GitBranch {
	return o.Payload
}

func (o *FindGitBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitBranch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindGitBranchBadRequest creates a FindGitBranchBadRequest with default headers values
func NewFindGitBranchBadRequest() *FindGitBranchBadRequest {
	return &FindGitBranchBadRequest{}
}

/* FindGitBranchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FindGitBranchBadRequest struct {
	Payload *models.Error
}

func (o *FindGitBranchBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branch/{branch_name}][%d] findGitBranchBadRequest  %+v", 400, o.Payload)
}
func (o *FindGitBranchBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindGitBranchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindGitBranchNotFound creates a FindGitBranchNotFound with default headers values
func NewFindGitBranchNotFound() *FindGitBranchNotFound {
	return &FindGitBranchNotFound{}
}

/* FindGitBranchNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FindGitBranchNotFound struct {
	Payload *models.Error
}

func (o *FindGitBranchNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branch/{branch_name}][%d] findGitBranchNotFound  %+v", 404, o.Payload)
}
func (o *FindGitBranchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindGitBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
