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

const headlessUIPropsSpecURL = "https://ofapkg.pek3b.qingstor.com/@one-for-all/headless-ui@latest/props-spec.json"
const headlessUIManifestURL = "https://ofapkg.pek3b.qingstor.com/@one-for-all/headless-ui@latest/manifest.json"
const iconManifestURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/one-for-all-icon-manifest.json"
const iconPropsPSpecURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/one-for-all-icon-props-spec.json"
const systemComponentManifestURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/system-component-manifest.json"
const systemComponentPropsSpecURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/system-component-props-spec.json"

type Key struct {
	Key     string      `json:"key"`
	Version string      `json:"version"`
	Value   interface{} `json:"value"`
}

type Keys struct {
	Keys []Key `json:"keys"`
}
