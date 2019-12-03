// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

// +build jsoniter

package json

import (
	"github.com/json-iterator/go"

	stdJson "encoding/json"
)

var (
	_json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal json marshal
	Marshal = _json.Marshal
	// Unmarshal json unmarshal
	Unmarshal = _json.Unmarshal
	// MarshalIndent json marshal indent
	MarshalIndent = _json.MarshalIndent
	// NewDecoder json new decoder function
	NewDecoder = _json.NewDecoder
	// NewEncoder json new encoder function
	NewEncoder = _json.NewEncoder
)

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage = stdJson.RawMessage
