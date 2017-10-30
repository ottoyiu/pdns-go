// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new zones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListZone zones managed by a server
*/
func (a *Client) ListZone(params *ListZoneParams) (*ListZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listZone",
		Method:             "GET",
		PathPattern:        "/servers/{server_id}/zones/{zone_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListZoneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListZoneOK), nil

}

/*
ListZones lists all zones in a server
*/
func (a *Client) ListZones(params *ListZonesParams) (*ListZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listZones",
		Method:             "GET",
		PathPattern:        "/servers/{server_id}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListZonesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListZonesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}