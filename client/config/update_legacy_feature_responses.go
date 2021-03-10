// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// UpdateLegacyFeatureReader is a Reader for the UpdateLegacyFeature structure.
type UpdateLegacyFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLegacyFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLegacyFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLegacyFeatureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLegacyFeatureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateLegacyFeatureUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateLegacyFeatureTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLegacyFeatureOK creates a UpdateLegacyFeatureOK with default headers values
func NewUpdateLegacyFeatureOK() *UpdateLegacyFeatureOK {
	return &UpdateLegacyFeatureOK{}
}

/* UpdateLegacyFeatureOK describes a response with status code 200, with default header values.

Legacy Feature
*/
type UpdateLegacyFeatureOK struct {
	Payload *models.LegacyFeature
}

func (o *UpdateLegacyFeatureOK) Error() string {
	return fmt.Sprintf("[PATCH /legacy_features/{legacy_feature_id}][%d] updateLegacyFeatureOK  %+v", 200, o.Payload)
}
func (o *UpdateLegacyFeatureOK) GetPayload() *models.LegacyFeature {
	return o.Payload
}

func (o *UpdateLegacyFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LegacyFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLegacyFeatureBadRequest creates a UpdateLegacyFeatureBadRequest with default headers values
func NewUpdateLegacyFeatureBadRequest() *UpdateLegacyFeatureBadRequest {
	return &UpdateLegacyFeatureBadRequest{}
}

/* UpdateLegacyFeatureBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateLegacyFeatureBadRequest struct {
	Payload *models.Error
}

func (o *UpdateLegacyFeatureBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /legacy_features/{legacy_feature_id}][%d] updateLegacyFeatureBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateLegacyFeatureBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLegacyFeatureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLegacyFeatureNotFound creates a UpdateLegacyFeatureNotFound with default headers values
func NewUpdateLegacyFeatureNotFound() *UpdateLegacyFeatureNotFound {
	return &UpdateLegacyFeatureNotFound{}
}

/* UpdateLegacyFeatureNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateLegacyFeatureNotFound struct {
	Payload *models.Error
}

func (o *UpdateLegacyFeatureNotFound) Error() string {
	return fmt.Sprintf("[PATCH /legacy_features/{legacy_feature_id}][%d] updateLegacyFeatureNotFound  %+v", 404, o.Payload)
}
func (o *UpdateLegacyFeatureNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLegacyFeatureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLegacyFeatureUnprocessableEntity creates a UpdateLegacyFeatureUnprocessableEntity with default headers values
func NewUpdateLegacyFeatureUnprocessableEntity() *UpdateLegacyFeatureUnprocessableEntity {
	return &UpdateLegacyFeatureUnprocessableEntity{}
}

/* UpdateLegacyFeatureUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateLegacyFeatureUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateLegacyFeatureUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /legacy_features/{legacy_feature_id}][%d] updateLegacyFeatureUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateLegacyFeatureUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateLegacyFeatureUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLegacyFeatureTooManyRequests creates a UpdateLegacyFeatureTooManyRequests with default headers values
func NewUpdateLegacyFeatureTooManyRequests() *UpdateLegacyFeatureTooManyRequests {
	return &UpdateLegacyFeatureTooManyRequests{}
}

/* UpdateLegacyFeatureTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateLegacyFeatureTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateLegacyFeatureTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /legacy_features/{legacy_feature_id}][%d] updateLegacyFeatureTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateLegacyFeatureTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLegacyFeatureTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
