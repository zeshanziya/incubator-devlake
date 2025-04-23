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
	"github.com/apache/incubator-devlake/plugins/axlzee/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/axlzee/tasks"
)

var _ interface {
	plugin.PluginMeta
	plugin.PluginInit
	plugin.PluginTask
	plugin.PluginSource
	plugin.PluginMigration
} = (*Axlzee)(nil)

type Axlzee struct{}

func (p Axlzee) Init(basicRes context.BasicRes) errors.Error {

	return nil
}


func (p Axlzee) Description() string {
	return "To collect and enrich data from axlzee Source"
}

func (p Axlzee) Name() string {
	return "axlzee"
}

func (p Axlzee) Connection() dal.Tabler {
	return nil
}

func (p Axlzee) Scope() plugin.ToolLayerScope {
	return nil
}

func (p Axlzee) ScopeConfig() dal.Tabler {
	return nil
}

func (p Axlzee) SubTaskMetas() []plugin.SubTaskMeta {
	return []plugin.SubTaskMeta{
		tasks.CollectUsersMeta,
		tasks.ExtractUserItemMeta,
	}
}

func (p Axlzee) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	var op tasks.AxlzeeOptions
	if err := helper.Decode(options, &op, nil); err != nil {
		return nil, err
	}
	taskCtx.GetLogger().Info("collect data", "data", op)
	return &tasks.AxlzeeTaskData{
		Options:   &op,
	}, nil
}

func (p Axlzee) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/axlzee"
}

func (p Axlzee) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}
