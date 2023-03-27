package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
)

func InitArtery() error {
	headlessUIPropsSpec, _ := getCtx(headlessUIPropsSpecURL)
	headlessUIManifest, _ := getCtx(headlessUIManifestURL)
	iconManifest, _ := getCtx(iconManifestURL)
	iconPropsPSpec, _ := getCtx(iconPropsPSpecURL)
	systemComponentManifest, _ := getCtx(systemComponentManifestURL)
	systemComponentPropsSpec, _ := getCtx(systemComponentPropsSpecURL)
	data := fmt.Sprintf(`[
    {
      "key": "PACKAGE_PROPS_SPEC:@one-for-all/headless-ui",
      "version": "0.8.3",
      "value": %s,
    },
    {
      "key": "PACKAGE_MANIFEST:@one-for-all/headless-ui",
      "version": "0.8.3",
      "value": %s,
    },
    {
      "key": "PACKAGE_MANIFEST:@one-for-all/icon",
      "version": "0.6.2",
      "value": %s,
    },
    {
      "key": "PACKAGE_PROPS_SPEC:@one-for-all/icon",
      "version": "0.6.2",
      "value": %s,
    },
    {
      "key": "PACKAGE_MANIFEST:system-components",
      "version": "1.0.0",
      "value": %s,
    },
    {
      "key": "PACKAGE_PROPS_SPEC:system-components",
      "version": "1.0.0",
      "value": %s,
    },
    {
      "key": "PACKAGES",
      "version": "1.0.0",
      "value": %s,
    },
  ]`, headlessUIPropsSpec, headlessUIManifest, iconManifest, iconPropsPSpec, systemComponentManifest, systemComponentPropsSpec, ALLPKG)
	return toolkits.DoPost("http://persona/batchSetValue", data)
}

func getCtx(URL string) (string, error) {
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return buf.String(), nil
}
