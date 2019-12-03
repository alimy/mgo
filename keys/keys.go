// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package keys

// IntsPool int-string 型池化kv-pairs接口
type IntsPool interface {
	Get(key int) string
}

// IntsPool string-string 型池化kv-pairs接口
type StrsPool interface {
	Get(key string) string
}

type keyGenBase struct {
	prefix   string
	suffix   string
	capacity int
}

// NewIntsPool 新建一个通过int型key获取string的IntsPool
func NewIntsPool(prefix, suffix string, cap int) IntsPool {
	return newKeyGenInts(prefix, suffix, cap)
}

// NewStrsPool 新建一个通过int型key获取string的IntsPool
func NewStrsPool(prefix, suffix string, cap int) StrsPool {
	return newKeyGenStrs(prefix, suffix, cap)
}
