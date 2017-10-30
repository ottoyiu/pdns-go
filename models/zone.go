// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Zone zone
// swagger:model Zone

type Zone struct {

	// MAY be set. Its value is defined by local policy
	Account string `json:"account,omitempty"`

	//  Whether or not the zone will be rectified on data changes via the API
	APIRectify bool `json:"api_rectify,omitempty"`

	// Whether or not this zone is DNSSEC signed (inferred from presigned being true XOR presence of at least one cryptokey with active being true)
	Dnssec bool `json:"dnssec,omitempty"`

	// Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs.
	ID string `json:"id,omitempty"`

	// Zone kind, one of “Native”, “Master”, “Slave”
	Kind string `json:"kind,omitempty"`

	//  List of IP addresses configured as a master for this zone (“Slave” type zones only)
	Masters []string `json:"masters"`

	// Name of the zone (e.g. “example.com.”) MUST have a trailing dot
	Name string `json:"name,omitempty"`

	// MAY be sent in client bodies during creation, and MUST NOT be sent by the server. Simple list of strings of nameserver names, including the trailing dot. Not required for slave zones.
	Nameservers []string `json:"nameservers"`

	// The SOA serial notifications have been sent out for
	NotifiedSerial int64 `json:"notified_serial,omitempty"`

	// Whether or not the zone uses NSEC3 narrow
	Nsec3narrow bool `json:"nsec3narrow,omitempty"`

	// The NSEC3PARAM record
	Nsec3param string `json:"nsec3param,omitempty"`

	// Whether or not the zone is pre-signed
	Presigned bool `json:"presigned,omitempty"`

	// rrsets
	Rrsets ZoneRrsets `json:"rrsets"`

	// The SOA serial number
	Serial int64 `json:"serial,omitempty"`

	// The SOA-EDIT metadata item
	SoaEdit string `json:"soa_edit,omitempty"`

	// The SOA-EDIT-API metadata item
	SoaEditAPI string `json:"soa_edit_api,omitempty"`

	// Set to “Zone”
	Type string `json:"type,omitempty"`

	// API endpoint for this zone
	URL string `json:"url,omitempty"`

	// MAY contain a BIND-style zone file when creating a zone
	Zone string `json:"zone,omitempty"`
}

/* polymorph Zone account false */

/* polymorph Zone api_rectify false */

/* polymorph Zone dnssec false */

/* polymorph Zone id false */

/* polymorph Zone kind false */

/* polymorph Zone masters false */

/* polymorph Zone name false */

/* polymorph Zone nameservers false */

/* polymorph Zone notified_serial false */

/* polymorph Zone nsec3narrow false */

/* polymorph Zone nsec3param false */

/* polymorph Zone presigned false */

/* polymorph Zone rrsets false */

/* polymorph Zone serial false */

/* polymorph Zone soa_edit false */

/* polymorph Zone soa_edit_api false */

/* polymorph Zone type false */

/* polymorph Zone url false */

/* polymorph Zone zone false */

// Validate validates this zone
func (m *Zone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMasters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNameservers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var zoneTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Native","Master","Slave"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zoneTypeKindPropEnum = append(zoneTypeKindPropEnum, v)
	}
}

const (
	// ZoneKindNative captures enum value "Native"
	ZoneKindNative string = "Native"
	// ZoneKindMaster captures enum value "Master"
	ZoneKindMaster string = "Master"
	// ZoneKindSLAVE captures enum value "Slave"
	ZoneKindSLAVE string = "Slave"
)

// prop value enum
func (m *Zone) validateKindEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, zoneTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Zone) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *Zone) validateMasters(formats strfmt.Registry) error {

	if swag.IsZero(m.Masters) { // not required
		return nil
	}

	return nil
}

func (m *Zone) validateNameservers(formats strfmt.Registry) error {

	if swag.IsZero(m.Nameservers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zone) UnmarshalBinary(b []byte) error {
	var res Zone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
