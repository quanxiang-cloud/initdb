package web

const headlessUIPropsSpecURL = "https://ofapkg.pek3b.qingstor.com/@one-for-all/headless-ui@latest/props-spec.json"
const headlessUIManifestURL = "https://ofapkg.pek3b.qingstor.com/@one-for-all/headless-ui@latest/manifest.json"
const iconManifestURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/one-for-all-icon-manifest.json"
const iconPropsPSpecURL = "https://ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/one-for-all-icon-props-spec.json"
const systemComponentManifestURL = "https: //ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/system-component-manifest.json"
const systemComponentPropsSpecURL = "https: //ofapkg.pek3b.qingstor.com/artery-engine-person-initialize-json/system-component-props-spec.json"
const ALLPKG = `
[
  {
    "name": "all",
    "label": "所有组件",
    "version": "1.0.0",
  },
  {
    "name": "system-components",
    "label": "系统组件",
    "version": "1.0.0",
  },
  {
    "name": "@one-for-all/icon",
    "label": "图标库",
    "version": "0.6.2",
    "categories": [
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
    ],
  },
  {
    "name": "@one-for-all/headless-ui",
    "label": "headless-ui",
    "version": "0.8.3",
    "categories": [
      "基础组件",
      "表单组件",
      "高级组件",
    ],
  },
]`
