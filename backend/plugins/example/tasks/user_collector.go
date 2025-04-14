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
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.SubTaskEntryPoint = CollectUsers

// CollectChat collect all chats that bot is in
func CollectUsers(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*ExampleTaskData)
	// log the message
	taskCtx.GetLogger().Info("collect users", data.Options.ConnectionId, data.Options.NumOfDaysToCollect)
	// format above line into single string

	return nil
}

var CollectUsersMeta = plugin.SubTaskMeta{
	Name:             "collectUsers",
	EntryPoint:       CollectUsers,
	EnabledByDefault: true,
	Description:      "Collect users from Example plugin",
}
