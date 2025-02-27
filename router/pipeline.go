// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package router

import (
	"github.com/gin-gonic/gin"

	"github.com/go-vela/worker/api"
)

// nolint: godot // ignore comment ending in period
//
// PipelineHandlers extends the provided base router group
// by adding a collection of endpoints for handling
// pipeline related requests.
//
// GET  /api/v1/executors/:executor/pipeline
func PipelineHandlers(base *gin.RouterGroup) {
	// add a collection of endpoints for handling pipeline related requests
	//
	// https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc#RouterGroup.Group
	pipeline := base.Group("/pipeline")
	{
		// add an endpoint for capturing the pipeline
		//
		// https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc#RouterGroup.GET
		pipeline.GET("", api.GetPipeline)
	}
}
