// ElasticView插件前端构建配置 - 基于Vite的现代化构建配置
// 支持Vue 3、TypeScript、微前端(qiankun)集成和Monaco编辑器

import { defineConfig } from 'vite'          // 导入Vite配置定义函数
import vue from '@vitejs/plugin-vue'         // 导入Vue 3官方Vite插件
import qiankun from 'vite-plugin-qiankun'    // 导入qiankun微前端插件，用于ElasticView插件系统集成
import pluginJson from './../plugin.json';  // 导入插件配置文件，获取插件基本信息
import { resolve } from "path";              // 导入路径解析工具
import { dynamicBase } from 'vite-plugin-dynamic-base' // 导入动态基础路径插件，支持运行时路径调整
import monacoEditorPlugin from "vite-plugin-monaco-editor"; // 导入Monaco编辑器插件，提供代码编辑功能

// 定义源代码路径，用于路径别名配置
const pathSrc = resolve(__dirname, "src");

// 导出Vite配置 - ElasticView插件前端构建的核心配置
export default defineConfig({
  // 依赖优化配置 - 确保第三方依赖正确预构建
  optimizeDeps: {
    force: true, // 强制重新预构建，解决插件开发中的依赖缓存问题
  },
  // 路径解析配置 - 设置模块解析规则
  resolve: {
    alias: {
      "@": pathSrc, // 设置@别名指向src目录，便于模块导入
    },
  },

  // 基础路径配置 - 根据环境动态设置
  // 生产环境使用动态基础路径，支持ElasticView插件系统的路径注入
  // 开发环境使用根路径，便于本地调试
  base: process.env.NODE_ENV === "production" ? "/__dynamic_base__/" : "/",

  // 生产模式标识（此处应该根据NODE_ENV动态设置）
  isProduction:false,
  
  // 插件配置 - ElasticView插件前端的核心插件集合
  plugins: [
    vue(), // Vue 3支持插件，提供.vue文件编译和热重载
    
    // Monaco编辑器插件配置 - 提供代码编辑功能（如JSON编辑器）
    monacoEditorPlugin.default({
      languageWorkers: ['json'], // 仅加载JSON语言支持，减少打包体积
    }),
    
    // qiankun微前端插件配置 - 集成到ElasticView插件系统
    qiankun(pluginJson.plugin_alias, {
      useDevMode: true // 开发模式，支持热重载和调试
    }),
    
    // 动态基础路径插件 - 支持运行时路径调整
    // ElasticView插件系统会在运行时注入正确的基础路径
    dynamicBase({
      // 从qiankun注入的公共路径中提取基础路径
      publicPath: 'window.proxy.__INJECTED_PUBLIC_PATH_BY_QIANKUN__.substring(0,window.proxy.__INJECTED_PUBLIC_PATH_BY_QIANKUN__.length-1)',
      transformIndexHtml: true // 转换index.html中的路径引用
    }),

  ],

  // 开发服务器配置 - 用于插件前端调试
  server: {
    port: pluginJson['frontend_dev_port'], // 使用plugin.json中配置的开发端口
    headers: {
      'Access-Control-Allow-Origin': '*'   // 允许跨域访问，支持ElasticView基座调用
    }
  }
})
