import {setupEvPlugin} from './plugin_sdk/setup'

// 引入依赖
import ElementPlus from 'element-plus'
// 引入全局 CSS 样式
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// 插件启动所需安装第三方插件通过此方法
setupEvPlugin((app)=>{
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }

  app.use(ElementPlus)
  return app
})
