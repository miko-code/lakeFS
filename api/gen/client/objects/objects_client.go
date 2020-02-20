// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new objects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for objects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteObject(params *DeleteObjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteObjectNoContent, error)

	GetObject(params *GetObjectParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetObjectOK, error)

	ListObjects(params *ListObjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListObjectsOK, error)

	StatObject(params *StatObjectParams, authInfo runtime.ClientAuthInfoWriter) (*StatObjectOK, error)

	UploadObject(params *UploadObjectParams, authInfo runtime.ClientAuthInfoWriter) (*UploadObjectCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteObject deletes object
*/
func (a *Client) DeleteObject(params *DeleteObjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteObjectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteObject",
		Method:             "DELETE",
		PathPattern:        "/repositories/{repositoryId}/branches/{branchId}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteObjectNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteObjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetObject gets object content
*/
func (a *Client) GetObject(params *GetObjectParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getObject",
		Method:             "GET",
		PathPattern:        "/repositories/{repositoryId}/branches/{branchId}/objects",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetObjectReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetObjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListObjects lists objects under a given tree
*/
func (a *Client) ListObjects(params *ListObjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListObjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listObjects",
		Method:             "GET",
		PathPattern:        "/repositories/{repositoryId}/branches/{branchId}/objects/ls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListObjectsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StatObject gets object metadata
*/
func (a *Client) StatObject(params *StatObjectParams, authInfo runtime.ClientAuthInfoWriter) (*StatObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statObject",
		Method:             "GET",
		PathPattern:        "/repositories/{repositoryId}/branches/{branchId}/objects/stat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StatObjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UploadObject uploads object content
*/
func (a *Client) UploadObject(params *UploadObjectParams, authInfo runtime.ClientAuthInfoWriter) (*UploadObjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadObject",
		Method:             "POST",
		PathPattern:        "/repositories/{repositoryId}/branches/{branchId}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadObjectCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadObjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
