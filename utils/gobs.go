// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"bytes"
	"encoding/gob"
)

// GobEncode encode obj by gob
func GobEncode(e interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	if err := enc.Encode(e); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// GobDecode decode byte slice to obj by gob
func GobDecode(bs []byte, e interface{}) error {
	b := bytes.NewBuffer(bs)
	dec := gob.NewDecoder(b)
	return dec.Decode(e)
}

// GobDecodeString decode string to obj by gob
func GobDecodeString(bs string, e interface{}) error {
	b := bytes.NewBufferString(bs)
	dec := gob.NewDecoder(b)
	return dec.Decode(e)
}
