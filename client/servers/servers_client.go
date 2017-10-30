// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListServer lists a server
*/
func (a *Client) ListServer(params *ListServerParams) (*ListServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServer",
		Method:             "GET",
		PathPattern:        "/servers/{server_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListServerOK), nil

}

/*
ListServers lists all servers
*/
func (a *Client) ListServers(params *ListServersParams) (*ListServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServers",
		Method:             "GET",
		PathPattern:        "/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListServersOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
