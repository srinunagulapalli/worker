// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"time"

	"github.com/go-vela/server/queue"
	"github.com/go-vela/worker/executor"
	"github.com/go-vela/worker/runtime"

	"github.com/urfave/cli/v2"
)

// flags is a helper function to return the all
// supported command line interface (CLI) flags
// for the Worker.
func flags() []cli.Flag {
	f := []cli.Flag{

		&cli.StringFlag{
			EnvVars: []string{"WORKER_ADDR", "VELA_WORKER_ADDR", "VELA_WORKER"},
			Name:    "worker.addr",
			Usage:   "Worker server address as a fully qualified url (<scheme>://<host>)",
		},

		&cli.DurationFlag{
			EnvVars: []string{"WORKER_CHECK_IN", "VELA_CHECK_IN", "CHECK_IN"},
			Name:    "checkIn",
			Usage:   "time to wait in between checking in with the server",
			Value:   15 * time.Minute,
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"WORKER_BUILD_LIMIT", "VELA_BUILD_LIMIT", "BUILD_LIMIT"},
			Name:    "build.limit",
			Usage:   "maximum amount of builds that can run concurrently",
			Value:   1,
		},
		&cli.DurationFlag{
			EnvVars: []string{"WORKER_BUILD_TIMEOUT", "VELA_BUILD_TIMEOUT", "BUILD_TIMEOUT"},
			Name:    "build.timeout",
			Usage:   "maximum amount of time a build can run for",
			Value:   30 * time.Minute,
		},

		// Logger Flags

		&cli.StringFlag{
			EnvVars: []string{"WORKER_LOG_FORMAT", "VELA_LOG_FORMAT", "LOG_FORMAT"},
			Name:    "log.format",
			Usage:   "set log format for the worker",
			Value:   "json",
		},
		&cli.StringFlag{
			EnvVars: []string{"WORKER_LOG_LEVEL", "VELA_LOG_LEVEL", "LOG_LEVEL"},
			Name:    "log.level",
			Usage:   "set log level for the worker",
			Value:   "info",
		},

		// Server Flags

		&cli.StringFlag{
			EnvVars: []string{"WORKER_SERVER_ADDR", "VELA_SERVER_ADDR", "VELA_SERVER", "SERVER_ADDR"},
			Name:    "server.addr",
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},
		&cli.StringFlag{
			EnvVars: []string{"WORKER_SERVER_SECRET", "VELA_SERVER_SECRET", "SERVER_SECRET"},
			Name:    "server.secret",
			Usage:   "secret used for server <-> worker communication",
		},
		&cli.StringFlag{
			EnvVars: []string{"WORKER_SERVER_CERT", "VELA_SERVER_CERT", "SERVER_CERT"},
			Name:    "server.cert",
			Usage:   "optional TLS certificate for https",
		},
		&cli.StringFlag{
			EnvVars: []string{"WORKER_SERVER_CERT_KEY", "VELA_SERVER_CERT_KEY", "SERVER_CERT_KEY"},
			Name:    "server.cert-key",
			Usage:   "optional TLS certificate key",
		},
	}

	// Executor Flags

	f = append(f, executor.Flags...)

	// Queue Flags

	f = append(f, queue.Flags...)

	// Runtime Flags

	f = append(f, runtime.Flags...)

	return f
}
