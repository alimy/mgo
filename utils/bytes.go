// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import "unsafe"

// BytesToString convert byte slice to string
func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// StringToBytes convert string to byte slice
func StringToBytes(s string) []byte {
	xs := (*[2]uintptr)(unsafe.Pointer(&s))
	hs := [3]uintptr{xs[0], xs[1], xs[1]}
	return *(*[]byte)(unsafe.Pointer(&hs))
}
