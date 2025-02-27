// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

import "github.com/go-vela/types/constants"

// Driver outputs the configured runtime driver.
func (c *client) Driver() string {
	return constants.DriverDocker
}
