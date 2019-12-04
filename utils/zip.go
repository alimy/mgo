// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Unzip unzips data by gzip.
func Unzip(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Zip zips data by gzip.
func Zip(data []byte) (res []byte, err error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err = w.Write(data)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	return buf.Bytes(), nil
}
