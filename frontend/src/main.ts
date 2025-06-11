// ElasticView插件前端应用入口文件 - 初始化Vue应用并集成ElasticView插件SDK
// 负责配置插件的国际化、UI组件库和路由系统

import {setupEvPlugin} from '@elasticview/plugin-sdk' // 导入ElasticView插件SDK核心设置函数
import pluginJson from '../../plugin.json'           // 导入插件配置信息

import App from "./App.vue";           // 导入Vue根组件
import router from "./router/index.js" // 导入路由配置

// 导入国际化语言包
import enLocale from "./lang/en";      // 英文语言包
import zhCnLocale from "./lang/zh-cn"; // 中文简体语言包

// 引入Element Plus UI组件库依赖
import ElementPlus from 'element-plus'   // 导入Element Plus组件库
// 引入Element Plus全局CSS样式
import 'element-plus/dist/index.css'    // Element Plus样式文件

import * as ElementPlusIconsVue from '@element-plus/icons-vue' // 导入Element Plus图标组件

// init函数 - 插件启动时的第三方依赖初始化配置
// ElasticView插件SDK会在适当时机调用此函数来初始化Vue应用实例
// 参数 app: Vue应用实例
// 返回值: 配置完成的Vue应用实例
const init = (app: any)=>{

  // 注册Element Plus图标组件到Vue应用
  // 遍历所有图标组件并注册为全局组件，便于在模板中直接使用
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component) // 注册图标组件，key为组件名，component为组件实例
  }

  // 使用Element Plus UI组件库
  // 将Element Plus注册到Vue应用，提供丰富的UI组件
  app.use(ElementPlus)

  return app // 返回配置完成的Vue应用实例
}

// 调用ElasticView插件SDK初始化函数 - 启动插件前端应用
// 参数说明：
// 1. pluginJson: 插件配置信息，包含插件元数据和路由配置
// 2. App: Vue根组件，插件的主要UI入口
// 3. router: Vue路由实例，定义插件的页面路由规则
// 4. enLocale: 英文语言包，支持国际化
// 5. zhCnLocale: 中文语言包，支持国际化
// 6. init: 第三方依赖初始化函数，配置UI组件库等
setupEvPlugin(pluginJson,App,router,enLocale,zhCnLocale,init)
