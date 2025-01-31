// Code generated by go-swagger; DO NOT EDIT.

package datagroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new datagroup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for datagroup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllDatagroups(params *AllDatagroupsParams, opts ...ClientOption) (*AllDatagroupsOK, error)

	Datagroup(params *DatagroupParams, opts ...ClientOption) (*DatagroupOK, error)

	UpdateDatagroup(params *UpdateDatagroupParams, opts ...ClientOption) (*UpdateDatagroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AllDatagroups gets all datagroups

  ### Get information about all datagroups.

*/
func (a *Client) AllDatagroups(params *AllDatagroupsParams, opts ...ClientOption) (*AllDatagroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllDatagroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_datagroups",
		Method:             "GET",
		PathPattern:        "/datagroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllDatagroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllDatagroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_datagroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Datagroup gets datagroup

  ### Get information about a datagroup.

*/
func (a *Client) Datagroup(params *DatagroupParams, opts ...ClientOption) (*DatagroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDatagroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "datagroup",
		Method:             "GET",
		PathPattern:        "/datagroups/{datagroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DatagroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DatagroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for datagroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDatagroup updates datagroup

  ### Update a datagroup using the specified params.

*/
func (a *Client) UpdateDatagroup(params *UpdateDatagroupParams, opts ...ClientOption) (*UpdateDatagroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDatagroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_datagroup",
		Method:             "PATCH",
		PathPattern:        "/datagroups/{datagroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDatagroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDatagroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_datagroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
