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

package configuration

import (
	"fmt"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	parser "github.com/haproxytech/config-parser/v4"
	"github.com/haproxytech/config-parser/v4/types"

	"github.com/theunknownport/client-native/v4/misc"
	"github.com/theunknownport/client-native/v4/models"
)

type Ring interface {
	GetRings(transactionID string) (int64, models.Rings, error)
	GetRing(name string, transactionID string) (int64, *models.Ring, error)
	DeleteRing(name string, transactionID string, version int64) error
	CreateRing(data *models.Ring, transactionID string, version int64) error
	EditRing(name string, data *models.Ring, transactionID string, version int64) error
}

// GetRings returns configuration version and an array of
// configured rings. Returns error on fail.
func (c *client) GetRings(transactionID string) (int64, models.Rings, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}

	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}

	rNames, err := p.SectionsGet(parser.Ring)
	if err != nil {
		return v, nil, err
	}

	rings := []*models.Ring{}
	for _, name := range rNames {
		_, a, err := c.GetRing(name, transactionID)
		if err == nil {
			rings = append(rings, a)
		}
	}

	return v, rings, nil
}

// GetRing returns configuration version and a requested ring.
// Returns error on fail or if ring does not exist.
func (c *client) GetRing(name string, transactionID string) (int64, *models.Ring, error) {
	p, err := c.GetParser(transactionID)
	if err != nil {
		return 0, nil, err
	}

	v, err := c.GetVersion(transactionID)
	if err != nil {
		return 0, nil, err
	}

	if !c.checkSectionExists(parser.Ring, name, p) {
		return v, nil, NewConfError(ErrObjectDoesNotExist, fmt.Sprintf("ring %s does not exist", name))
	}

	ring := &models.Ring{Name: name}

	description, err := p.Get(parser.Ring, name, "description", true)
	if err != nil {
		return v, nil, err
	}
	desc, ok := description.(*types.StringC)
	if !ok {
		return 0, nil, misc.CreateTypeAssertError("description")
	}

	if desc.Value != "" {
		ring.Description = desc.Value
	}

	format, err := p.Get(parser.Ring, name, "format", true)
	if err != nil {
		return v, nil, err
	}
	f, ok := format.(*types.StringC)
	if !ok {
		return v, nil, misc.CreateTypeAssertError("format")
	}
	if f.Value != "" {
		ring.Format = f.Value
	}

	maxlen, err := p.Get(parser.Ring, name, "maxlen", true)
	if err != nil {
		return v, nil, err
	}
	mx, ok := maxlen.(*types.Int64C)
	if !ok {
		return v, nil, misc.CreateTypeAssertError("maxlen")
	}
	if mx.Value > 0 {
		ring.Maxlen = &mx.Value
	}

	size, err := p.Get(parser.Ring, name, "size", true)
	if err != nil {
		return v, nil, err
	}
	sz, ok := size.(*types.StringC)
	if !ok {
		return v, nil, misc.CreateTypeAssertError("size")
	}
	if sz.Value != "" {
		ring.Size = misc.ParseSize(sz.Value)
	}

	// timeouts
	tConnect, err := p.Get(parser.Ring, name, "timeout connect", true)
	if err != nil {
		return v, nil, err
	}
	tc, ok := tConnect.(*types.SimpleTimeout)
	if !ok {
		return v, nil, misc.CreateTypeAssertError("timeout connect")
	}
	if tc.Value != "" {
		ring.TimeoutConnect = misc.ParseTimeout(tc.Value)
	}

	tServer, err := p.Get(parser.Ring, name, "timeout server", true)
	if err != nil {
		return v, nil, err
	}
	ts, ok := tServer.(*types.SimpleTimeout)
	if !ok {
		return v, nil, misc.CreateTypeAssertError("timeout server")
	}
	if ts.Value != "" {
		ring.TimeoutServer = misc.ParseTimeout(ts.Value)
	}

	return v, ring, nil
}

// DeleteRing deletes a ring in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) DeleteRing(name string, transactionID string, version int64) error {
	if err := c.deleteSection(parser.Ring, name, transactionID, version); err != nil {
		return err
	}
	return nil
}

// CreateRing creates a ring in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) CreateRing(data *models.Ring, transactionID string, version int64) error {
	if c.UseModelsValidation {
		validationErr := data.Validate(strfmt.Default)
		if validationErr != nil {
			return NewConfError(ErrValidationError, validationErr.Error())
		}
	}
	p, t, err := c.loadDataForChange(transactionID, version)
	if err != nil {
		return err
	}

	if c.checkSectionExists(parser.Ring, data.Name, p) {
		e := NewConfError(ErrObjectAlreadyExists, fmt.Sprintf("%s %s already exists", parser.Ring, data.Name))
		return c.HandleError(data.Name, "", "", t, transactionID == "", e)
	}

	if err = p.SectionsCreate(parser.Ring, data.Name); err != nil {
		return c.HandleError(data.Name, "", "", t, transactionID == "", err)
	}

	if err = SerializeRingSection(p, data); err != nil {
		return err
	}

	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}

	return nil
}

// EditRing edits a ring in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *client) EditRing(name string, data *models.Ring, transactionID string, version int64) error {
	if c.UseModelsValidation {
		validationErr := data.Validate(strfmt.Default)
		if validationErr != nil {
			return NewConfError(ErrValidationError, validationErr.Error())
		}
	}
	p, t, err := c.loadDataForChange(transactionID, version)
	if err != nil {
		return err
	}

	if !c.checkSectionExists(parser.Ring, data.Name, p) {
		e := NewConfError(ErrObjectAlreadyExists, fmt.Sprintf("%s %s does not exists", parser.Ring, data.Name))
		return c.HandleError(data.Name, "", "", t, transactionID == "", e)
	}

	if err = SerializeRingSection(p, data); err != nil {
		return err
	}

	if err := c.SaveData(p, t, transactionID == ""); err != nil {
		return err
	}
	return nil
}

func SerializeRingSection(p parser.Parser, data *models.Ring) error { //nolint:gocognit
	if data == nil {
		return fmt.Errorf("empty ring")
	}

	var err error
	if data.Description == "" {
		if err = p.Set(parser.Ring, data.Name, "description", nil); err != nil {
			return err
		}
	} else {
		d := types.StringC{Value: data.Description}
		if err = p.Set(parser.Ring, data.Name, "description", d); err != nil {
			return err
		}
	}

	if data.Format == "" {
		if err = p.Set(parser.Ring, data.Name, "format", nil); err != nil {
			return err
		}
	} else {
		d := types.StringC{Value: data.Format}
		if err = p.Set(parser.Ring, data.Name, "format", d); err != nil {
			return err
		}
	}

	if data.Maxlen == nil {
		if err = p.Set(parser.Ring, data.Name, "maxlen", nil); err != nil {
			return err
		}
	} else {
		d := types.Int64C{Value: *data.Maxlen}
		if err = p.Set(parser.Ring, data.Name, "maxlen", d); err != nil {
			return err
		}
	}

	if data.Size == nil {
		if err = p.Set(parser.Ring, data.Name, "size", nil); err != nil {
			return err
		}
	} else {
		d := types.StringC{Value: fmt.Sprint(*data.Size)}
		if err = p.Set(parser.Ring, data.Name, "size", d); err != nil {
			return err
		}
	}

	if data.TimeoutConnect == nil {
		if err = p.Set(parser.Ring, data.Name, "timeout connect", nil); err != nil {
			return err
		}
	} else {
		tc := types.SimpleTimeout{Value: strconv.FormatInt(*data.TimeoutConnect, 10)}
		if err = p.Set(parser.Ring, data.Name, "timeout connect", tc); err != nil {
			return err
		}
	}

	if data.TimeoutServer == nil {
		if err = p.Set(parser.Ring, data.Name, "timeout server", nil); err != nil {
			return err
		}
	} else {
		ts := types.SimpleTimeout{Value: strconv.FormatInt(*data.TimeoutServer, 10)}
		if err = p.Set(parser.Ring, data.Name, "timeout server", ts); err != nil {
			return err
		}
	}

	return err
}
