/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package web

import (
	"encoding/json"

	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
)

func InitNav() error {
	var navKeys Keys
	Nav := []map[string]interface{}{
		{
			"id":    "app_views",
			"title": "应用视图",
			"icon":  "view",
			"children": []map[string]interface{}{
				{
					"id":    "views",
					"title": "页面管理",
					"icon":  "note_detail_duotone",
				},
				{
					"id":    "view_layout",
					"title": "母版管理",
					"icon":  "row_top_duotone",
				},
				{
					"id":    "app_nav",
					"title": "应用导航",
					"icon":  "tab_duotone",
				},
			},
		},
		{
			"id":    "modal_api",
			"title": "数据管理",
			"icon":  "gateway",
			"children": []map[string]interface{}{
				{
					"id":    "data_models",
					"title": "数据模型管理",
					"icon":  "database",
				},
				{
					"id":    "api_proxy",
					"title": "第三方 API 代理",
					"icon":  "api_outside",
				},
				{
					"id":    "orchestration_api",
					"title": "API 编排管理",
					"icon":  "api_arrange",
				},
				{
					"id":    "faas",
					"title": "FaaS 函数管理",
					"icon":  "faas_control",
				},
				{
					"id":    "key_api",
					"title": "API 密钥管理",
					"icon":  "api_key",
				},
				{
					"id":    "file_api",
					"title": "API 文档",
					"icon":  "api_inner",
				},
			},
		},
		{
			"id":    "setting_flow",
			"title": "工作流",
			"icon":  "data_model",
		},
		{
			"id":    "app_control",
			"title": "访问控制",
			"icon":  "role",
		},
		{
			"id":    "base_info",
			"title": "应用设置",
			"icon":  "app_setting",
		},
	}

	navJ, _ := json.Marshal(Nav)
	navKeys.Keys = append(navKeys.Keys, Key{
		Key:     "PORTAL_APPLICATION_SIDE_NAV",
		Version: "0.1.0",
		Value:   string(navJ),
	})

	return toolkits.DoPost("http://persona/api/v1/persona/batchSetValue", navKeys)
}
