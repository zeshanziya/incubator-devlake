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

package impl

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/example/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/example/tasks"
)

var _ interface {
	plugin.PluginMeta
	plugin.PluginInit
	plugin.PluginTask
	plugin.PluginSource
	plugin.PluginMigration
} = (*Example)(nil)

type Example struct{}

func (p Example) Init(basicRes context.BasicRes) errors.Error {

	return nil
}


func (p Example) Description() string {
	return "To collect and enrich data from Example Source"
}

func (p Example) Name() string {
	return "example"
}

func (p Example) Connection() dal.Tabler {
	return nil
}

func (p Example) Scope() plugin.ToolLayerScope {
	return nil
}

func (p Example) ScopeConfig() dal.Tabler {
	return nil
}

func (p Example) SubTaskMetas() []plugin.SubTaskMeta {
	return []plugin.SubTaskMeta{
		tasks.CollectUsersMeta,
		tasks.ExtractUserItemMeta,
	}
}

func (p Example) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	var op tasks.ExampleOptions
	if err := helper.Decode(options, &op, nil); err != nil {
		return nil, err
	}
	taskCtx.GetLogger().Info("collect data", "data", op)
	return &tasks.ExampleTaskData{
		Options:   &op,
	}, nil
}

func (p Example) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/example"
}

func (p Example) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}
