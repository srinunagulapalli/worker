// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-vela/worker/version"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

// hostname stores the worker host name reported by the kernel.
var hostname string

// create an init function to set the hostname for the worker.
//
// https://golang.org/doc/effective_go.html#init
func init() {
	// attempt to capture the hostname for the worker
	hostname, _ = os.Hostname()
	// check if a hostname is set
	if len(hostname) == 0 {
		// default the hostname to localhost
		hostname = "localhost"
	}
}

func main() {
	// capture application version information
	v := version.New()

	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	app := cli.NewApp()

	// Worker Information

	app.Name = "vela-worker"
	app.HelpName = "vela-worker"
	app.Usage = "Vela build daemon designed for executing pipelines"
	app.Copyright = "Copyright (c) 2022 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Worker Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = v.Semantic()

	// Worker Flags

	app.Flags = flags()

	// Worker Start

	err = app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
