/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"net/http"
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

// CreateApiClient creates a new API client for the Random User API
func CreateApiClient(taskCtx plugin.SubTaskContext) (*api.ApiAsyncClient, errors.Error) {
	// Create a simple API client for the Random User API
	apiClient, err := api.NewApiClient(
		taskCtx.GetContext(),
		"https://randomuser.me/api/",
		nil,
		30*time.Second,
		"",
		taskCtx,
	)
	if err != nil {
		return nil, err
	}

	// Set common headers
	apiClient.SetHeaders(map[string]string{
		"Accept": "application/json",
	})

	// Add response handler to check status code
	apiClient.SetAfterFunction(func(res *http.Response) errors.Error {
		if res.StatusCode < 200 || res.StatusCode >= 300 {
			return errors.HttpStatus(res.StatusCode).New("HTTP status error")
		}
		return nil
	})

	// Create an async client with rate limiting
	asyncApiClient, err := api.CreateAsyncApiClient(
		taskCtx.TaskContext(),
		apiClient,
		&api.ApiRateLimitCalculator{
			UserRateLimitPerHour: 1000, // Set a reasonable rate limit
		},
	)
	if err != nil {
		return nil, err
	}

	return asyncApiClient, nil
}
