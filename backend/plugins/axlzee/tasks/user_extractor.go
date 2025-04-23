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

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/axlzee/apimodels"
	"github.com/apache/incubator-devlake/plugins/axlzee/models"
)

var _ plugin.SubTaskEntryPoint = ExtractUserItem

func ExtractUserItem(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*AxlzeeTaskData)
	extractor, err := api.NewApiExtractor(api.ApiExtractorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: AxlzeeApiParams{
				ConnectionId: data.Options.ConnectionId,
			},
			Table: RAW_USER_TABLE,
		},
		Extract: func(row *api.RawData) ([]interface{}, errors.Error) {
			body := &apimodels.AxlzeeUserItem{}
			err := errors.Convert(json.Unmarshal(row.Data, body))
			if err != nil {
				return nil, err
			}
			user := &models.AxlzeeUser{}
			user.ConnectionId = data.Options.ConnectionId
			user.Age = body.Dob.Age
			user.Email = body.Email
			user.FirstName = body.Name.First
			user.LastName = body.Name.Last
			user.Gender = body.Gender
			user.City = body.Location.City
			user.State = body.Location.State
			user.Phone = body.Phone
			user.UserId = body.Login.Uuid

			return []interface{}{user}, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractUserItemMeta = plugin.SubTaskMeta{
	Name:             "extractUserItem",
	EntryPoint:       ExtractUserItem,
	EnabledByDefault: true,
	Description:      "Extract raw users data into tool layer table axlzee_users",
}
