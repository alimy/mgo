// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"errors"
	"testing"
)

func TestGob(t *testing.T) {
	type user struct {
		Age  uint8
		Name string
	}
	alice := &user{
		Age:  18,
		Name: "alice",
	}
	var err error
	alx := &user{}
	aly := &user{}
	bytesAlice, err := GobEncode(alice)
	if err != nil {
		goto ErrEdge
	}
	if err = GobDecode(bytesAlice, alx); err != nil {
		goto ErrEdge
	}
	if err = GobDecodeString(string(bytesAlice), aly); err != nil {
		goto ErrEdge
	}
	if alx.Age != alice.Age || aly.Age != alice.Age || alx.Name != alice.Name || aly.Name != alice.Name {
		err = errors.New("decoded obj not equal wanted")
		goto ErrEdge
	}
	return // pass test case just return
ErrEdge:
	t.Errorf("gob encode/decode failure: %s", err)
}
