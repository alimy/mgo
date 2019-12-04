// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import "testing"

func TestZip(t *testing.T) {
	hi := "Hi,I'm Bob"
	zips, err := Zip([]byte(hi))
	if err != nil {
		t.Errorf("zip failure: %s", err)
	}
	his, err := Unzip(zips)
	if err != nil {
		t.Errorf("unzip failure: %s", err)
	}
	if string(his) != hi {
		t.Errorf("unzip from zip's return not equal origin")
	}
	return
}
