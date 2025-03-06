import {setupEvPlugin} from '@elasticview/plugin-sdk'
import pluginJson from '../../plugin.json'

import App from "./App.vue";
import router from "./router/index.js"

import enLocale from "./lang/en";
import zhCnLocale from "./lang/zh-cn";

// 引入依赖
import ElementPlus from 'element-plus'
// 引入全局 CSS 样式
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// 插件启动所需安装第三方插件通过此方法
const init = (app)=>{

  // 注册Element Plus图标
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }

  // 使用Element Plus
  app.use(ElementPlus)

  return app
}

setupEvPlugin(pluginJson,App,router,enLocale,zhCnLocale,init)
