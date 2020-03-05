// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/mgo/mgo/internal/cmd"
)

func main() {
	// setup root cli command of application
	cmd.Setup(
		"mgo",              // command name
		"mgo help toolkit", // command short describe
		"mgo help toolkit", // command long describe
	)

	// execute start application
	cmd.Execute()
}
