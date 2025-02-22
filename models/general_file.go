// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GeneralFile General use file
//
// General use file
//
// swagger:model general_file
type GeneralFile struct {

	// description
	Description string `json:"description,omitempty"`

	// file
	File string `json:"file,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// storage name
	StorageName string `json:"storage_name,omitempty"`
}

// Validate validates this general file
func (m *GeneralFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this general file based on context it is used
func (m *GeneralFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeneralFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralFile) UnmarshalBinary(b []byte) error {
	var res GeneralFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
