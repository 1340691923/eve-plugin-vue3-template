// ElasticView插件路由系统入口文件 - 初始化Vue Router实例
// 根据ElasticView插件开发规范配置路由系统

import {
  createRouter,         // 导入Vue Router创建函数
  createWebHashHistory  // 导入Hash路由历史模式，适用于插件环境
} from 'vue-router'
import routes from './routes' // 导入路由配置定义

// 创建Vue Router实例 - ElasticView插件的路由管理器
// 使用Hash历史模式确保插件在微前端环境下的路由兼容性
const router = createRouter({
  history: createWebHashHistory(), // Hash模式：插件运行在ElasticView基座中，使用Hash避免路径冲突
  routes                           // 路由配置：包含插件的所有页面路由定义
})

// 导出路由实例 - 供main.ts中的setupEvPlugin函数使用
export default router

