// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package keys

import "fmt"

type keyGenStrs struct {
	keyGenBase
	keyPool map[string]string
}

func newKeyGenStrs(prefix, suffix string, cap int) StrsPool {
	return &keyGenStrs{
		keyGenBase: keyGenBase{
			prefix:   prefix,
			suffix:   suffix,
			capacity: cap,
		},
		keyPool: make(map[string]string, cap),
	}
}

// Get 通过key获取string型值
func (k *keyGenStrs) Get(key string) string {
	if res, exist := k.keyPool[key]; exist {
		return res
	}
	// 暴力失效法，超过capacity就直接失效，后面的热点请求按需重建keyPool
	if len(k.keyPool) == k.capacity {
		k.keyPool = make(map[string]string, k.capacity)
	}
	res := fmt.Sprintf("%s%s%s", k.prefix, key, k.suffix)
	k.keyPool[key] = res
	return res
}
