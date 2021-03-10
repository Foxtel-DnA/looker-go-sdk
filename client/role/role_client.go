// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new role API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for role API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllModelSets(params *AllModelSetsParams, opts ...ClientOption) (*AllModelSetsOK, error)

	AllPermissionSets(params *AllPermissionSetsParams, opts ...ClientOption) (*AllPermissionSetsOK, error)

	AllPermissions(params *AllPermissionsParams, opts ...ClientOption) (*AllPermissionsOK, error)

	AllRoles(params *AllRolesParams, opts ...ClientOption) (*AllRolesOK, error)

	CreateModelSet(params *CreateModelSetParams, opts ...ClientOption) (*CreateModelSetOK, error)

	CreatePermissionSet(params *CreatePermissionSetParams, opts ...ClientOption) (*CreatePermissionSetOK, error)

	CreateRole(params *CreateRoleParams, opts ...ClientOption) (*CreateRoleOK, error)

	DeleteModelSet(params *DeleteModelSetParams, opts ...ClientOption) (*DeleteModelSetNoContent, error)

	DeletePermissionSet(params *DeletePermissionSetParams, opts ...ClientOption) (*DeletePermissionSetNoContent, error)

	DeleteRole(params *DeleteRoleParams, opts ...ClientOption) (*DeleteRoleNoContent, error)

	ModelSet(params *ModelSetParams, opts ...ClientOption) (*ModelSetOK, error)

	PermissionSet(params *PermissionSetParams, opts ...ClientOption) (*PermissionSetOK, error)

	Role(params *RoleParams, opts ...ClientOption) (*RoleOK, error)

	RoleGroups(params *RoleGroupsParams, opts ...ClientOption) (*RoleGroupsOK, error)

	RoleUsers(params *RoleUsersParams, opts ...ClientOption) (*RoleUsersOK, error)

	SearchModelSets(params *SearchModelSetsParams, opts ...ClientOption) (*SearchModelSetsOK, error)

	SearchPermissionSets(params *SearchPermissionSetsParams, opts ...ClientOption) (*SearchPermissionSetsOK, error)

	SearchRoles(params *SearchRolesParams, opts ...ClientOption) (*SearchRolesOK, error)

	SetRoleGroups(params *SetRoleGroupsParams, opts ...ClientOption) (*SetRoleGroupsOK, error)

	SetRoleUsers(params *SetRoleUsersParams, opts ...ClientOption) (*SetRoleUsersOK, error)

	UpdateModelSet(params *UpdateModelSetParams, opts ...ClientOption) (*UpdateModelSetOK, error)

	UpdatePermissionSet(params *UpdatePermissionSetParams, opts ...ClientOption) (*UpdatePermissionSetOK, error)

	UpdateRole(params *UpdateRoleParams, opts ...ClientOption) (*UpdateRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AllModelSets gets all model sets

  ### Get information about all model sets.

*/
func (a *Client) AllModelSets(params *AllModelSetsParams, opts ...ClientOption) (*AllModelSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllModelSetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_model_sets",
		Method:             "GET",
		PathPattern:        "/model_sets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllModelSetsReader{formats: a.formats},
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
	success, ok := result.(*AllModelSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_model_sets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AllPermissionSets gets all permission sets

  ### Get information about all permission sets.

*/
func (a *Client) AllPermissionSets(params *AllPermissionSetsParams, opts ...ClientOption) (*AllPermissionSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllPermissionSetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_permission_sets",
		Method:             "GET",
		PathPattern:        "/permission_sets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllPermissionSetsReader{formats: a.formats},
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
	success, ok := result.(*AllPermissionSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_permission_sets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AllPermissions gets all permissions

  ### Get all supported permissions.

*/
func (a *Client) AllPermissions(params *AllPermissionsParams, opts ...ClientOption) (*AllPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_permissions",
		Method:             "GET",
		PathPattern:        "/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllPermissionsReader{formats: a.formats},
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
	success, ok := result.(*AllPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_permissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AllRoles gets all roles

  ### Get information about all roles.

*/
func (a *Client) AllRoles(params *AllRolesParams, opts ...ClientOption) (*AllRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_roles",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllRolesReader{formats: a.formats},
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
	success, ok := result.(*AllRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateModelSet creates model set

  ### Create a model set with the specified information. Model sets are used by Roles.

*/
func (a *Client) CreateModelSet(params *CreateModelSetParams, opts ...ClientOption) (*CreateModelSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateModelSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_model_set",
		Method:             "POST",
		PathPattern:        "/model_sets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateModelSetReader{formats: a.formats},
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
	success, ok := result.(*CreateModelSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_model_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePermissionSet creates permission set

  ### Create a permission set with the specified information. Permission sets are used by Roles.

*/
func (a *Client) CreatePermissionSet(params *CreatePermissionSetParams, opts ...ClientOption) (*CreatePermissionSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePermissionSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_permission_set",
		Method:             "POST",
		PathPattern:        "/permission_sets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePermissionSetReader{formats: a.formats},
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
	success, ok := result.(*CreatePermissionSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_permission_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateRole creates role

  ### Create a role with the specified information.

*/
func (a *Client) CreateRole(params *CreateRoleParams, opts ...ClientOption) (*CreateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_role",
		Method:             "POST",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRoleReader{formats: a.formats},
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
	success, ok := result.(*CreateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_role: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteModelSet deletes model set

  ### Delete the model set with a specific id.

*/
func (a *Client) DeleteModelSet(params *DeleteModelSetParams, opts ...ClientOption) (*DeleteModelSetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModelSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_model_set",
		Method:             "DELETE",
		PathPattern:        "/model_sets/{model_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteModelSetReader{formats: a.formats},
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
	success, ok := result.(*DeleteModelSetNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_model_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePermissionSet deletes permission set

  ### Delete the permission set with a specific id.

*/
func (a *Client) DeletePermissionSet(params *DeletePermissionSetParams, opts ...ClientOption) (*DeletePermissionSetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePermissionSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_permission_set",
		Method:             "DELETE",
		PathPattern:        "/permission_sets/{permission_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePermissionSetReader{formats: a.formats},
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
	success, ok := result.(*DeletePermissionSetNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_permission_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRole deletes role

  ### Delete the role with a specific id.

*/
func (a *Client) DeleteRole(params *DeleteRoleParams, opts ...ClientOption) (*DeleteRoleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_role",
		Method:             "DELETE",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoleReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_role: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModelSet gets model set

  ### Get information about the model set with a specific id.

*/
func (a *Client) ModelSet(params *ModelSetParams, opts ...ClientOption) (*ModelSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModelSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "model_set",
		Method:             "GET",
		PathPattern:        "/model_sets/{model_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModelSetReader{formats: a.formats},
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
	success, ok := result.(*ModelSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for model_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PermissionSet gets permission set

  ### Get information about the permission set with a specific id.

*/
func (a *Client) PermissionSet(params *PermissionSetParams, opts ...ClientOption) (*PermissionSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermissionSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "permission_set",
		Method:             "GET",
		PathPattern:        "/permission_sets/{permission_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PermissionSetReader{formats: a.formats},
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
	success, ok := result.(*PermissionSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for permission_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Role gets role

  ### Get information about the role with a specific id.

*/
func (a *Client) Role(params *RoleParams, opts ...ClientOption) (*RoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoleReader{formats: a.formats},
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
	success, ok := result.(*RoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RoleGroups gets role groups

  ### Get information about all the groups with the role that has a specific id.

*/
func (a *Client) RoleGroups(params *RoleGroupsParams, opts ...ClientOption) (*RoleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role_groups",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoleGroupsReader{formats: a.formats},
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
	success, ok := result.(*RoleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RoleUsers gets role users

  ### Get information about all the users with the role that has a specific id.

*/
func (a *Client) RoleUsers(params *RoleUsersParams, opts ...ClientOption) (*RoleUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoleUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "role_users",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoleUsersReader{formats: a.formats},
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
	success, ok := result.(*RoleUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for role_users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchModelSets searches model sets

  ### Search model sets
Returns all model set records that match the given search criteria.
If multiple search params are given and `filter_or` is FALSE or not specified,
search params are combined in a logical AND operation.
Only rows that match *all* search param criteria will be returned.

If `filter_or` is TRUE, multiple search params are combined in a logical OR operation.
Results will include rows that match **any** of the search criteria.

String search params use case-insensitive matching.
String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions.
example="dan%" will match "danger" and "Danzig" but not "David"
example="D_m%" will match "Damage" and "dump"

Integer search params can accept a single value or a comma separated list of values. The multiple
values will be combined under a logical OR operation - results will match at least one of
the given values.

Most search params can accept "IS NULL" and "NOT NULL" as special expressions to match
or exclude (respectively) rows where the column is null.

Boolean search params accept only "true" and "false" as values.


*/
func (a *Client) SearchModelSets(params *SearchModelSetsParams, opts ...ClientOption) (*SearchModelSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchModelSetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_model_sets",
		Method:             "GET",
		PathPattern:        "/model_sets/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchModelSetsReader{formats: a.formats},
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
	success, ok := result.(*SearchModelSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_model_sets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchPermissionSets searches permission sets

  ### Search permission sets
Returns all permission set records that match the given search criteria.
If multiple search params are given and `filter_or` is FALSE or not specified,
search params are combined in a logical AND operation.
Only rows that match *all* search param criteria will be returned.

If `filter_or` is TRUE, multiple search params are combined in a logical OR operation.
Results will include rows that match **any** of the search criteria.

String search params use case-insensitive matching.
String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions.
example="dan%" will match "danger" and "Danzig" but not "David"
example="D_m%" will match "Damage" and "dump"

Integer search params can accept a single value or a comma separated list of values. The multiple
values will be combined under a logical OR operation - results will match at least one of
the given values.

Most search params can accept "IS NULL" and "NOT NULL" as special expressions to match
or exclude (respectively) rows where the column is null.

Boolean search params accept only "true" and "false" as values.


*/
func (a *Client) SearchPermissionSets(params *SearchPermissionSetsParams, opts ...ClientOption) (*SearchPermissionSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchPermissionSetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_permission_sets",
		Method:             "GET",
		PathPattern:        "/permission_sets/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchPermissionSetsReader{formats: a.formats},
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
	success, ok := result.(*SearchPermissionSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_permission_sets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchRoles searches roles

  ### Search roles

Returns all role records that match the given search criteria.

If multiple search params are given and `filter_or` is FALSE or not specified,
search params are combined in a logical AND operation.
Only rows that match *all* search param criteria will be returned.

If `filter_or` is TRUE, multiple search params are combined in a logical OR operation.
Results will include rows that match **any** of the search criteria.

String search params use case-insensitive matching.
String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions.
example="dan%" will match "danger" and "Danzig" but not "David"
example="D_m%" will match "Damage" and "dump"

Integer search params can accept a single value or a comma separated list of values. The multiple
values will be combined under a logical OR operation - results will match at least one of
the given values.

Most search params can accept "IS NULL" and "NOT NULL" as special expressions to match
or exclude (respectively) rows where the column is null.

Boolean search params accept only "true" and "false" as values.


*/
func (a *Client) SearchRoles(params *SearchRolesParams, opts ...ClientOption) (*SearchRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_roles",
		Method:             "GET",
		PathPattern:        "/roles/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchRolesReader{formats: a.formats},
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
	success, ok := result.(*SearchRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetRoleGroups updates role groups

  ### Set all groups for a role, removing all existing group associations from that role.

*/
func (a *Client) SetRoleGroups(params *SetRoleGroupsParams, opts ...ClientOption) (*SetRoleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRoleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set_role_groups",
		Method:             "PUT",
		PathPattern:        "/roles/{role_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetRoleGroupsReader{formats: a.formats},
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
	success, ok := result.(*SetRoleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set_role_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetRoleUsers updates role users

  ### Set all the users of the role with a specific id.

*/
func (a *Client) SetRoleUsers(params *SetRoleUsersParams, opts ...ClientOption) (*SetRoleUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRoleUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set_role_users",
		Method:             "PUT",
		PathPattern:        "/roles/{role_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetRoleUsersReader{formats: a.formats},
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
	success, ok := result.(*SetRoleUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set_role_users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateModelSet updates model set

  ### Update information about the model set with a specific id.

*/
func (a *Client) UpdateModelSet(params *UpdateModelSetParams, opts ...ClientOption) (*UpdateModelSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateModelSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_model_set",
		Method:             "PATCH",
		PathPattern:        "/model_sets/{model_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateModelSetReader{formats: a.formats},
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
	success, ok := result.(*UpdateModelSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_model_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePermissionSet updates permission set

  ### Update information about the permission set with a specific id.

*/
func (a *Client) UpdatePermissionSet(params *UpdatePermissionSetParams, opts ...ClientOption) (*UpdatePermissionSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePermissionSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_permission_set",
		Method:             "PATCH",
		PathPattern:        "/permission_sets/{permission_set_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePermissionSetReader{formats: a.formats},
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
	success, ok := result.(*UpdatePermissionSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_permission_set: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRole updates role

  ### Update information about the role with a specific id.

*/
func (a *Client) UpdateRole(params *UpdateRoleParams, opts ...ClientOption) (*UpdateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_role",
		Method:             "PATCH",
		PathPattern:        "/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoleReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_role: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
