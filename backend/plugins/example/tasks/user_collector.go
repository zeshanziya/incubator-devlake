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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/example/apimodels"
)

const RAW_USER_TABLE = "example_users"

var CollectUsersMeta = plugin.SubTaskMeta{
	Name:             "collectUsers",
	EntryPoint:       CollectUsers,
	EnabledByDefault: true,
	Description:      "Collect users from Random User API",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CROSS},
}

// CollectUsers collects user data from the Random User API
func CollectUsers(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*ExampleTaskData)
	logger := taskCtx.GetLogger()
	connectionId := data.Options.ConnectionId
	numOfDaysToCollect := int(data.Options.NumOfDaysToCollect)

	// If numOfDaysToCollect is not specified or is less than 1, default to 10 users
	numUsers := 10
	if numOfDaysToCollect > 0 {
		numUsers = numOfDaysToCollect
	}

	logger.Info("Collecting users from Random User API", "connectionId", connectionId, "numUsers", numUsers)

	// Create API client
	apiClient, err := CreateApiClient(taskCtx)
	if err != nil {
		return err
	}

	// Create collector
	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx:    taskCtx,
			Table:  RAW_USER_TABLE,
			Params: ExampleApiParams{ConnectionId: connectionId},
		},
		ApiClient: apiClient,
		// We don't need an iterator since we're making a single request
		UrlTemplate: fmt.Sprintf("?results=%d", numUsers),
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			body := &apimodels.ExampleUserApiResult{}
			err = api.UnmarshalResponse(res, body)
			if err != nil {
				return nil, errors.Convert(err)
			}
			return body.Results, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}
