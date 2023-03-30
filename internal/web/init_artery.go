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
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
)

func InitArtery() error {
	var ArteryKeys Keys
	//var headlessUIPropsSpec, headlessUIManifest, iconManifest, iconPropsPSpec, systemComponentManifest, systemComponentPropsSpec map[interface{}]interface{}
	var ALLPKG = []interface{}{
		map[string]interface{}{"name": "all",
			"label":   "所有组件",
			"version": "1.0.0"},
		map[string]interface{}{
			"name":    "system-components",
			"label":   "系统组件",
			"version": "1.0.0",
		},
		map[string]interface{}{
			"name":    "@one-for-all/icon",
			"label":   "图标库",
			"version": "0.6.2",
			"categories": []string{
				"action",
				"alert",
				"av",
				"background",
				"communication",
				"content",
				"device",
				"editor",
				"file",
				"hardware",
				"home",
				"image",
				"maps",
				"mobile",
				"navigation",
				"notification",
				"places",
				"qxp_ui",
				"social",
				"toggle",
			},
		},
		map[string]interface{}{
			"name":    "@one-for-all/headless-ui",
			"label":   "headless-ui",
			"version": "0.8.3",
			"categories": []string{
				"基础组件",
				"表单组件",
				"高级组件",
			},
		},
	}
	allpkg, _ := json.Marshal(&ALLPKG)
	headlessUIPropsSpec := getCtx(headlessUIPropsSpecURL)
	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_PROPS_SPEC:@one-for-all/headless-ui",
		Version: "0.8.3",
		Value:   headlessUIPropsSpec,
	})
	headlessUIManifest := getCtx(headlessUIManifestURL)

	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_MANIFEST:@one-for-all/headless-ui",
		Version: "0.8.3",
		Value:   headlessUIManifest,
	})
	iconManifest := getCtx(iconManifestURL)

	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_MANIFEST:@one-for-all/icon",
		Version: "0.6.2",
		Value:   iconManifest,
	})
	iconPropsPSpec := getCtx(iconPropsPSpecURL)
	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_PROPS_SPEC:@one-for-all/icon",
		Version: "0.6.2",
		Value:   iconPropsPSpec,
	})
	systemComponentManifest := getCtx(systemComponentManifestURL)

	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_MANIFEST:system-components",
		Version: "1.0.0",
		Value:   systemComponentManifest,
	})
	systemComponentPropsSpec := getCtx(systemComponentPropsSpecURL)

	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGE_PROPS_SPEC:system-components",
		Version: "1.0.0",
		Value:   systemComponentPropsSpec,
	})
	ArteryKeys.Keys = append(ArteryKeys.Keys, Key{
		Key:     "PACKAGES",
		Version: "1.0.0",
		Value:   string(allpkg),
	})
	return toolkits.DoPost("http://persona/api/v1/persona/batchSetValue", ArteryKeys)
}

func getCtx(URL string) string {
	response, err := http.Get(URL)
	if err != nil {
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return buf.String()
}
