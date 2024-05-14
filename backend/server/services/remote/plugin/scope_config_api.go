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

package plugin

import (
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

func (pa *pluginAPI) PostScopeConfigs(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ScopeConfigApi.Post(input)
}

func (pa *pluginAPI) PatchScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ScopeConfigApi.Patch(input)
}

func (pa *pluginAPI) GetScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ConnApi.GetDetail(input)
}

func (pa *pluginAPI) GetProjectsByScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ScopeConfigApi.GetProjectsByScopeConfig(input)
}

func (pa *pluginAPI) ListScopeConfigs(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ScopeConfigApi.GetAll(input)
}

func (pa *pluginAPI) DeleteScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return pa.dsHelper.ScopeConfigApi.Delete(input)
}
